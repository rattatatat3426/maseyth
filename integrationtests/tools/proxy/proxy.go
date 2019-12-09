package quicproxy

import (
	"net"
	"sync"
	"time"

	"github.com/Psiphon-Labs/quic-go/internal/protocol"
	"github.com/Psiphon-Labs/quic-go/internal/utils"
)

// Connection is a UDP connection
type connection struct {
	ClientAddr *net.UDPAddr // Address of the client
	ServerConn *net.UDPConn // UDP connection to server
}

// Direction is the direction a packet is sent.
type Direction int

const (
	// DirectionIncoming is the direction from the client to the server.
	DirectionIncoming Direction = iota
	// DirectionOutgoing is the direction from the server to the client.
	DirectionOutgoing
	// DirectionBoth is both incoming and outgoing
	DirectionBoth
)

func (d Direction) String() string {
	switch d {
	case DirectionIncoming:
		return "incoming"
	case DirectionOutgoing:
		return "outgoing"
	case DirectionBoth:
		return "both"
	default:
		panic("unknown direction")
	}
}

// Is says if one direction matches another direction.
// For example, incoming matches both incoming and both, but not outgoing.
func (d Direction) Is(dir Direction) bool {
	if d == DirectionBoth || dir == DirectionBoth {
		return true
	}
	return d == dir
}

// DropCallback is a callback that determines which packet gets dropped.
type DropCallback func(dir Direction, packet []byte) bool

// NoDropper doesn't drop packets.
var NoDropper DropCallback = func(Direction, []byte) bool {
	return false
}

// DelayCallback is a callback that determines how much delay to apply to a packet.
type DelayCallback func(dir Direction, packet []byte) time.Duration

// NoDelay doesn't apply a delay.
var NoDelay DelayCallback = func(Direction, []byte) time.Duration {
	return 0
}

// Opts are proxy options.
type Opts struct {
	// The address this proxy proxies packets to.
	RemoteAddr string
	// DropPacket determines whether a packet gets dropped.
	DropPacket DropCallback
	// DelayPacket determines how long a packet gets delayed. This allows
	// simulating a connection with non-zero RTTs.
	// Note that the RTT is the sum of the delay for the incoming and the outgoing packet.
	DelayPacket DelayCallback
}

// QuicProxy is a QUIC proxy that can drop and delay packets.
type QuicProxy struct {
	mutex sync.Mutex

	conn       *net.UDPConn
	serverAddr *net.UDPAddr

	dropPacket  DropCallback
	delayPacket DelayCallback

	timerID uint64
	timers  map[uint64]*time.Timer

	// Mapping from client addresses (as host:port) to connection
	clientDict map[string]*connection

	logger utils.Logger
}

// NewQuicProxy creates a new UDP proxy
func NewQuicProxy(local string, opts *Opts) (*QuicProxy, error) {
	if opts == nil {
		opts = &Opts{}
	}
	laddr, err := net.ResolveUDPAddr("udp", local)
	if err != nil {
		return nil, err
	}
	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		return nil, err
	}
	raddr, err := net.ResolveUDPAddr("udp", opts.RemoteAddr)
	if err != nil {
		return nil, err
	}

	packetDropper := NoDropper
	if opts.DropPacket != nil {
		packetDropper = opts.DropPacket
	}

	packetDelayer := NoDelay
	if opts.DelayPacket != nil {
		packetDelayer = opts.DelayPacket
	}

	p := QuicProxy{
		clientDict:  make(map[string]*connection),
		conn:        conn,
		serverAddr:  raddr,
		dropPacket:  packetDropper,
		delayPacket: packetDelayer,
		timers:      make(map[uint64]*time.Timer),
		logger:      utils.DefaultLogger.WithPrefix("proxy"),
	}

	p.logger.Debugf("Starting UDP Proxy %s <-> %s", conn.LocalAddr(), raddr)
	go p.runProxy()
	return &p, nil
}

// Close stops the UDP Proxy
func (p *QuicProxy) Close() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for _, c := range p.clientDict {
		if err := c.ServerConn.Close(); err != nil {
			return err
		}
	}
	for _, t := range p.timers {
		t.Stop()
	}
	return p.conn.Close()
}

// LocalAddr is the address the proxy is listening on.
func (p *QuicProxy) LocalAddr() net.Addr {
	return p.conn.LocalAddr()
}

// LocalPort is the UDP port number the proxy is listening on.
func (p *QuicProxy) LocalPort() int {
	return p.conn.LocalAddr().(*net.UDPAddr).Port
}

func (p *QuicProxy) newConnection(cliAddr *net.UDPAddr) (*connection, error) {
	srvudp, err := net.DialUDP("udp", nil, p.serverAddr)
	if err != nil {
		return nil, err
	}
	return &connection{
		ClientAddr: cliAddr,
		ServerConn: srvudp,
	}, nil
}

// runProxy listens on the proxy address and handles incoming packets.
func (p *QuicProxy) runProxy() error {
	for {
		buffer := make([]byte, protocol.MaxReceivePacketSize)
		n, cliaddr, err := p.conn.ReadFromUDP(buffer)
		if err != nil {
			return err
		}
		raw := buffer[0:n]

		saddr := cliaddr.String()
		p.mutex.Lock()
		conn, ok := p.clientDict[saddr]

		if !ok {
			conn, err = p.newConnection(cliaddr)
			if err != nil {
				p.mutex.Unlock()
				return err
			}
			p.clientDict[saddr] = conn
			go p.runConnection(conn)
		}
		p.mutex.Unlock()

		if p.dropPacket(DirectionIncoming, raw) {
			if p.logger.Debug() {
				p.logger.Debugf("dropping incoming packet(%d bytes)", n)
			}
			continue
		}

		// Send the packet to the server
		delay := p.delayPacket(DirectionIncoming, raw)
		if delay != 0 {
			if p.logger.Debug() {
				p.logger.Debugf("delaying incoming packet (%d bytes) to %s by %s", n, conn.ServerConn.RemoteAddr(), delay)
			}

			p.mutex.Lock()
			p.timerID++
			id := p.timerID
			timer := time.AfterFunc(delay, func() {
				_, _ = conn.ServerConn.Write(raw) // TODO: handle error

				p.mutex.Lock()
				delete(p.timers, id)
				p.mutex.Unlock()
			})
			p.timers[id] = timer
			p.mutex.Unlock()
		} else {
			if p.logger.Debug() {
				p.logger.Debugf("forwarding incoming packet (%d bytes) to %s", n, conn.ServerConn.RemoteAddr())
			}
			if _, err := conn.ServerConn.Write(raw); err != nil {
				return err
			}
		}
	}
}

// runConnection handles packets from server to a single client
func (p *QuicProxy) runConnection(conn *connection) error {
	for {
		buffer := make([]byte, protocol.MaxReceivePacketSize)
		n, err := conn.ServerConn.Read(buffer)
		if err != nil {
			return err
		}
		raw := buffer[0:n]

		if p.dropPacket(DirectionOutgoing, raw) {
			if p.logger.Debug() {
				p.logger.Debugf("dropping outgoing packet(%d bytes)", n)
			}
			continue
		}

		delay := p.delayPacket(DirectionOutgoing, raw)
		if delay != 0 {
			if p.logger.Debug() {
				p.logger.Debugf("delaying outgoing packet (%d bytes) to %s by %s", n, conn.ClientAddr, delay)
			}

			p.mutex.Lock()
			p.timerID++
			id := p.timerID
			timer := time.AfterFunc(delay, func() {
				_, _ = p.conn.WriteToUDP(raw, conn.ClientAddr) // TODO: handle error
				p.mutex.Lock()
				delete(p.timers, id)
				p.mutex.Unlock()
			})
			p.timers[id] = timer
			p.mutex.Unlock()
		} else {
			if p.logger.Debug() {
				p.logger.Debugf("forwarding outgoing packet (%d bytes) to %s", n, conn.ClientAddr)
			}
			if _, err := p.conn.WriteToUDP(raw, conn.ClientAddr); err != nil {
				return err
			}
		}
	}
}
