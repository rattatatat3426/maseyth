//go:build generate

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"strings"
	"text/template"

	"golang.org/x/tools/imports"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatalf("Usage: %s <struct_name> <input_file> <template_file> <output_file>", os.Args[0])
	}

	structName := os.Args[1]
	inputFile := os.Args[2]
	templateFile := os.Args[3]
	outputFile := os.Args[4]

	fset := token.NewFileSet()

	// Parse the input file containing the struct type
	file, err := parser.ParseFile(fset, inputFile, nil, parser.AllErrors)
	if err != nil {
		log.Fatalf("Failed to parse file: %v", err)
	}

	var fields []*ast.Field

	// Find the specified struct type in the AST
	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok || typeSpec.Name.Name != structName {
				continue
			}
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				log.Fatalf("%s is not a struct", structName)
			}
			fields = structType.Fields.List
			break
		}
	}

	if fields == nil {
		log.Fatalf("Could not find %s type", structName)
	}

	// Prepare data for the template
	type FieldData struct {
		Name        string
		Params      string
		Args        string
		HasParams   bool
		ReturnTypes string
		HasReturn   bool
	}

	var fieldDataList []FieldData

	for _, field := range fields {
		funcType, ok := field.Type.(*ast.FuncType)
		if !ok {
			continue
		}
		for _, name := range field.Names {
			fieldData := FieldData{Name: name.Name}

			// extract parameters
			var params []string
			var args []string
			if funcType.Params != nil {
				for i, param := range funcType.Params.List {
					// We intentionally reject unnamed (and, further down, "_") function parameters.
					// We could auto-generate parameter names,
					// but having meaningful variable names will be more helpful for the user.
					if len(param.Names) == 0 {
						log.Fatalf("encountered unnamed parameter at position %d in function %s", i, fieldData.Name)
					}
					var buf bytes.Buffer
					printer.Fprint(&buf, fset, param.Type)
					paramType := buf.String()
					for _, paramName := range param.Names {
						if paramName.Name == "_" {
							log.Fatalf("encountered underscore parameter at position %d in function %s", i, fieldData.Name)
						}
						params = append(params, fmt.Sprintf("%s %s", paramName.Name, paramType))
						args = append(args, paramName.Name)
					}
				}
			}
			fieldData.Params = strings.Join(params, ", ")
			fieldData.Args = strings.Join(args, ", ")
			fieldData.HasParams = len(params) > 0

			// extract return types
			if funcType.Results != nil && len(funcType.Results.List) > 0 {
				fieldData.HasReturn = true
				var returns []string
				for _, result := range funcType.Results.List {
					var buf bytes.Buffer
					printer.Fprint(&buf, fset, result.Type)
					returns = append(returns, buf.String())
				}
				if len(returns) == 1 {
					fieldData.ReturnTypes = fmt.Sprintf(" %s", returns[0])
				} else {
					fieldData.ReturnTypes = fmt.Sprintf(" (%s)", strings.Join(returns, ", "))
				}
			}

			fieldDataList = append(fieldDataList, fieldData)
		}
	}

	// Read the template from file
	templateContent, err := os.ReadFile(templateFile)
	if err != nil {
		log.Fatalf("Failed to read template file: %v", err)
	}

	// Generate the code using the template
	tmpl, err := template.New("multiplexer").Funcs(template.FuncMap{"join": strings.Join}).Parse(string(templateContent))
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	var generatedCode bytes.Buffer
	if err = tmpl.Execute(&generatedCode, map[string]interface{}{
		"Fields":     fieldDataList,
		"StructName": structName,
	}); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	// Format the generated code and add imports
	formattedCode, err := imports.Process(outputFile, generatedCode.Bytes(), nil)
	if err != nil {
		log.Fatalf("Failed to process imports: %v", err)
	}

	if err := os.WriteFile(outputFile, formattedCode, 0o644); err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}
}
