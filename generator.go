package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

type generatorParam struct {
	Services   []Service
	GrpcAddr   string
	AdminPort  string
	PbPath     string
	AddImports []string
}

type Options struct {
	writer    io.Writer
	grpcAddr  string
	adminPort string
	pbPath    string
}

type ProtobufGolangReference struct {
	GolangReference string
	GolangImport    string
}

// This should contain all protobuf references that we want to support, with at least the well known types
var referenceTypeRegistry = map[string]ProtobufGolangReference{
	// General use Well Known Types
	"google.protobuf.Any": {
		GolangImport:    "github.com/golang/protobuf/ptypes/any",
		GolangReference: "any.Any",
	},
	"google.protobuf.Duration": {
		GolangImport:    "github.com/golang/protobuf/ptypes/duration",
		GolangReference: "duration.Duration",
	},
	"google.protobuf.Empty": {
		GolangImport:    "github.com/golang/protobuf/ptypes/empty",
		GolangReference: "empty.Empty",
	},
	"google.protobuf.Timestamp": {
		GolangImport:    "github.com/golang/protobuf/ptypes/timestamp",
		GolangReference: "timestamp.Timestamp",
	},

	// Basic Value Wrapper Well Known Types
	"google.protobuf.BoolValue": {
		GolangImport:    "github.com/golang/protobuf/ptypes/wrappers",
		GolangReference: "wrappers.BoolValue",
	},
	"google.protobuf.BytesValue": {
		GolangImport:    "github.com/golang/protobuf/ptypes/wrappers",
		GolangReference: "wrappers.BytesValue",
	},
	"google.protobuf.DoubleValue": {
		GolangImport:    "github.com/golang/protobuf/ptypes/wrappers",
		GolangReference: "wrappers.DoubleValue",
	},
	"google.protobuf.FloatValue": {
		GolangImport:    "github.com/golang/protobuf/ptypes/wrappers",
		GolangReference: "wrappers.FloatValue",
	},
	"google.protobuf.Int32Value": {
		GolangImport:    "github.com/golang/protobuf/ptypes/wrappers",
		GolangReference: "wrappers.Int32Value",
	},
	"google.protobuf.Int64Value": {
		GolangImport:    "github.com/golang/protobuf/ptypes/wrappers",
		GolangReference: "wrappers.Int64Value",
	},
	"google.protobuf.StringValue": {
		GolangImport:    "github.com/golang/protobuf/ptypes/wrappers",
		GolangReference: "wrappers.StringValue",
	},
	"google.protobuf.UInt32Value": {
		GolangImport:    "github.com/golang/protobuf/ptypes/wrappers",
		GolangReference: "wrappers.UInt32Value",
	},
	"google.protobuf.UInt64Value": {
		GolangImport:    "github.com/golang/protobuf/ptypes/wrappers",
		GolangReference: "wrappers.UInt64Value",
	},

	// Special Value Wrapper Well Known Types
	"google.protobuf.ListValue": {
		GolangImport:    "github.com/golang/protobuf/ptypes/struct",
		GolangReference: "struct.ListValue",
	},
	"google.protobuf.NullValue": {
		GolangImport:    "github.com/golang/protobuf/ptypes/struct",
		GolangReference: "struct.NullValue",
	},
	"google.protobuf.Struct": {
		GolangImport:    "github.com/golang/protobuf/ptypes/struct",
		GolangReference: "struct.Struct",
	},
	"google.protobuf.Value": {
		GolangImport:    "github.com/golang/protobuf/ptypes/struct",
		GolangReference: "struct.Value",
	},
}

func GenerateServer(services []Service, opt *Options) error {
	addImportsSet := make(map[string]bool)
	// Iterate over services and replace the type, add the correct import to Imports
	for _, service := range services {
		for _, method := range service.Methods {
			if golangReference, ok := referenceTypeRegistry[method.Input]; ok {
				method.Input = golangReference.GolangReference
				addImportsSet[fmt.Sprintf("\"%s\"", golangReference.GolangImport)] = true
			}
			if golangReference, ok := referenceTypeRegistry[method.Output]; ok {
				method.Output = golangReference.GolangReference
				addImportsSet[fmt.Sprintf("\"%s\"", golangReference.GolangImport)] = true
			}
		}
	}
	addImports := make([]string, 0, len(addImportsSet))
	for key := range addImportsSet {
		addImports = append(addImports, key)
	}

	param := generatorParam{
		Services:   services,
		GrpcAddr:   opt.grpcAddr,
		AdminPort:  opt.adminPort,
		PbPath:     opt.pbPath,
		AddImports: addImports,
	}

	if opt == nil {
		opt = &Options{}
	}

	if opt.writer == nil {
		opt.writer = os.Stdout
	}

	tmpl := template.New("server.tmpl").Funcs(template.FuncMap{
		"Title": strings.Title,
	})
	tmpl, err := tmpl.Parse(SERVER_TEMPLATE)
	if err != nil {
		return err
	}

	return tmpl.Execute(opt.writer, param)
}

const SERVER_TEMPLATE = `// DO NOT EDIT. This file is autogenerated by GripMock
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	{{ range .AddImports }}
		{{ . }}
	{{ end }}

	"github.com/mitchellh/mapstructure"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	TCP_ADDRESS  = "{{.GrpcAddr}}"
	HTTP_PORT = ":{{.AdminPort}}"
)

{{ range .Services }}
{{ template "services" . }}
{{ end }}

func main() {
	lis, err := net.Listen("tcp", TCP_ADDRESS)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	{{ range .Services }}
	{{ template "register_services" . }}
	{{ end }}

	reflection.Register(s)
	fmt.Println("Serving gRPC on tcp://" + TCP_ADDRESS)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

{{ template "find_stub" }}

{{ define "services" }}
type {{.Name}} struct{}

{{ template "methods" .}}
{{ end }}

{{ define "methods" }}
{{ $serviceName := .Name }}
{{ range .Methods}}
{{ $methodName := .Name | Title }}
func (s *{{$serviceName}}) {{$methodName}}(ctx context.Context, in *{{.Input}}) (*{{.Output}},error){
	out := &{{.Output}}{}
	err := findStub("{{$serviceName}}", "{{$methodName}}", in, out)
	return out, err
}
{{end}}
{{end}}

{{ define "register_services" }}
	Register{{.Name}}Server(s, &{{.Name}}{})
{{ end }}

{{ define "find_stub" }}
type payload struct {
	Service string      ` + "`json:\"service\"`" + `
	Method  string      ` + "`json:\"method\"`" + `
	Data    interface{} ` + "`json:\"data\"`" + `
}

type grpcError struct {
	Message string     ` + "`json:\"message\"`" + `
	Code    codes.Code ` + "`json:\"code\"`" + `
}

type response struct {
	Data  interface{}       ` + "`json:\"data\"`" + `
	ErrorObject grpcError   ` + "`json:\"errorObject\"`" + `
	Error string            ` + "`json:\"error\"`" + `
}

func findStub(service, method string, in, out interface{}) error {
	url := fmt.Sprintf("http://localhost%s/find", HTTP_PORT)
	pyl := payload{
		Service: service,
		Method:  method,
		Data:    in,
	}
	byt, err := json.Marshal(pyl)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(byt)
	resp, err := http.DefaultClient.Post(url, "application/json", reader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf(string(body))
	}

	respRPC := new(response)
	err = json.NewDecoder(resp.Body).Decode(respRPC)
	if err != nil {
		return err
	}

	if (respRPC.ErrorObject != grpcError{}) {
		return status.Error(respRPC.ErrorObject.Code, respRPC.ErrorObject.Message)
	} else if respRPC.Error != "" {
		return fmt.Errorf(respRPC.Error)
	}

	return mapstructure.Decode(respRPC.Data, out)
}
{{ end }}`