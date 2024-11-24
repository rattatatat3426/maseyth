//go:build gomock || generate

package quic

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/rattatatat3426/maseyth -destination mock_send_conn_test.go github.com/rattatatat3426/maseyth SendConn"
type SendConn = sendConn

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/rattatatat3426/maseyth -destination mock_raw_conn_test.go github.com/rattatatat3426/maseyth RawConn"
type RawConn = rawConn

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/rattatatat3426/maseyth -destination mock_sender_test.go github.com/rattatatat3426/maseyth Sender"
type Sender = sender

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/rattatatat3426/maseyth -destination mock_stream_internal_test.go github.com/rattatatat3426/maseyth StreamI"
type StreamI = streamI

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/rattatatat3426/maseyth -destination mock_crypto_stream_test.go github.com/rattatatat3426/maseyth CryptoStream"
type CryptoStream = cryptoStream

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/rattatatat3426/maseyth -destination mock_receive_stream_internal_test.go github.com/rattatatat3426/maseyth ReceiveStreamI"
type ReceiveStreamI = receiveStreamI

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/rattatatat3426/maseyth -destination mock_send_stream_internal_test.go github.com/rattatatat3426/maseyth SendStreamI"
type SendStreamI = sendStreamI

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/rattatatat3426/maseyth -destination mock_stream_getter_test.go github.com/rattatatat3426/maseyth StreamGetter"
type StreamGetter = streamGetter

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/rattatatat3426/maseyth -destination mock_stream_sender_test.go github.com/rattatatat3426/maseyth StreamSender"
type StreamSender = streamSender

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/rattatatat3426/maseyth -destination mock_crypto_data_handler_test.go github.com/rattatatat3426/maseyth CryptoDataHandler"
type CryptoDataHandler = cryptoDataHandler

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/rattatatat3426/maseyth -destination mock_frame_source_test.go github.com/rattatatat3426/maseyth FrameSource"
type FrameSource = frameSource

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/rattatatat3426/maseyth -destination mock_ack_frame_source_test.go github.com/rattatatat3426/maseyth AckFrameSource"
type AckFrameSource = ackFrameSource

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/rattatatat3426/maseyth -destination mock_stream_manager_test.go github.com/rattatatat3426/maseyth StreamManager"
type StreamManager = streamManager

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/rattatatat3426/maseyth -destination mock_sealing_manager_test.go github.com/rattatatat3426/maseyth SealingManager"
type SealingManager = sealingManager

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/rattatatat3426/maseyth -destination mock_unpacker_test.go github.com/rattatatat3426/maseyth Unpacker"
type Unpacker = unpacker

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/rattatatat3426/maseyth -destination mock_packer_test.go github.com/rattatatat3426/maseyth Packer"
type Packer = packer

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/rattatatat3426/maseyth -destination mock_mtu_discoverer_test.go github.com/rattatatat3426/maseyth MTUDiscoverer"
type MTUDiscoverer = mtuDiscoverer

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/rattatatat3426/maseyth -destination mock_conn_runner_test.go github.com/rattatatat3426/maseyth ConnRunner"
type ConnRunner = connRunner

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/rattatatat3426/maseyth -destination mock_quic_conn_test.go github.com/rattatatat3426/maseyth QUICConn"
type QUICConn = quicConn

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/rattatatat3426/maseyth -destination mock_packet_handler_test.go github.com/rattatatat3426/maseyth PacketHandler"
type PacketHandler = packetHandler

//go:generate sh -c "go run go.uber.org/mock/mockgen -build_flags=\"-tags=gomock\" -package quic -self_package github.com/rattatatat3426/maseyth -destination mock_packet_handler_manager_test.go github.com/rattatatat3426/maseyth PacketHandlerManager"

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/rattatatat3426/maseyth -destination mock_packet_handler_manager_test.go github.com/rattatatat3426/maseyth PacketHandlerManager"
type PacketHandlerManager = packetHandlerManager

// Need to use source mode for the batchConn, since reflect mode follows type aliases.
// See https://github.com/golang/mock/issues/244 for details.
//
//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -package quic -self_package github.com/rattatatat3426/maseyth -source sys_conn_oob.go -destination mock_batch_conn_test.go -mock_names batchConn=MockBatchConn"

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -package quic -self_package github.com/rattatatat3426/maseyth -self_package github.com/rattatatat3426/maseyth -destination mock_token_store_test.go github.com/rattatatat3426/maseyth TokenStore"
//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -package quic -self_package github.com/rattatatat3426/maseyth -self_package github.com/rattatatat3426/maseyth -destination mock_packetconn_test.go net PacketConn"
