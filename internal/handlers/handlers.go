package handlers

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	contextPackage = protogen.GoImportPath("context")
	grpcPackage    = protogen.GoImportPath("google.golang.org/grpc")
	fmtPackage     = protogen.GoImportPath("fmt")
)

// generateFile generates a _mock.pb.go file containing mock gRPC service definitions.
func GenerateFile(gen *protogen.Plugin, file *protogen.File, version string) *protogen.GeneratedFile {
	if len(file.Services) == 0 {
		return nil
	}
	filename := file.GeneratedFilenamePrefix + "_mock.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-mock. DO NOT EDIT.")
	g.P("// versions:")
	g.P("// - protoc-gen-go-mock v", version)
	g.P("// - protoc             ", protocVersion(gen))
	g.P("// source: ", file.Desc.Path())
	g.P()
	g.P("package ", file.GoPackageName)

	generateFileContent(gen, file, g)
	return g
}

// generateFileContent generates the gRPC service definitions, excluding the package statement.
func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Services) == 0 {
		return
	}

	for _, service := range file.Services {
		genService(gen, file, g, service)
	}
}

func genService(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) {
	clientName := service.GoName + "MockClient"

	// Check if implements interface
	g.P("var _ ", service.GoName, "Client", "=&", clientName, "{}")

	g.P("type ", clientName, " struct {")
	for _, method := range service.Methods {
		g.P(service.GoName, method.GoName, "Handler")
	}
	g.P("}")
	g.P()

	for _, method := range service.Methods {
		g.P("type ", service.GoName, method.GoName, "Handler func ", methodSignature(g, method))
	}
	g.P()

	g.P("func New", clientName, " () *", service.GoName, "MockClient {")
	g.P("return &", clientName, "{}")
	g.P("}")
	g.P()

	// Client method implementations.
	for _, method := range service.Methods {
		genClientMethod(gen, file, g, method)
	}
}

func genClientMethod(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, method *protogen.Method) {
	service := method.Parent
	name := fmt.Sprint(service.GoName, method.GoName, "Handler")
	g.P("func (c *", service.GoName, "MockClient) ", clientSignature(g, method), "{")
	// Unary
	if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
		g.P("handler := c.", name)
		g.P("if handler == nil {")
		g.P("return nil, ", g.QualifiedGoIdent(fmtPackage.Ident("Errorf")), "(\"no handler registered for ", method.GoName, "\")")
		g.P("}")
		g.P()
		g.P("return handler(ctx, in, opts...)")
		g.P("}")
		g.P()
		return
	}

	g.P("handler := c.", service.GoName, method.GoName, "Handler")
	g.P("if handler == nil {")
	g.P("return nil, ", g.QualifiedGoIdent(fmtPackage.Ident("Errorf")), "(\"no handler registered for ", method.GoName, "\")")
	g.P("}")
	g.P()
	s := "return handler(ctx"
	if !method.Desc.IsStreamingClient() {
		s += ", in"
	}
	s += ", opts...)"
	g.P(s)
	g.P("}")
	g.P()
}

func clientSignature(g *protogen.GeneratedFile, method *protogen.Method) string {
	s := method.GoName + methodSignature(g, method)
	return s
}

func methodSignature(g *protogen.GeneratedFile, method *protogen.Method) string {
	s := "(ctx " + g.QualifiedGoIdent(contextPackage.Ident("Context"))
	if !method.Desc.IsStreamingClient() {
		s += ", in *" + g.QualifiedGoIdent(method.Input.GoIdent)
	}
	s += ", opts ..." + g.QualifiedGoIdent(grpcPackage.Ident("CallOption")) + ") ("
	if !method.Desc.IsStreamingClient() && !method.Desc.IsStreamingServer() {
		s += "*" + g.QualifiedGoIdent(method.Output.GoIdent)
	} else {
		s += method.Parent.GoName + "_" + method.GoName + "Client"
	}
	s += ", error)"
	return s
}

func protocVersion(gen *protogen.Plugin) string {
	v := gen.Request.GetCompilerVersion()
	if v == nil {
		return "(unknown)"
	}
	var suffix string
	if s := v.GetSuffix(); s != "" {
		suffix = "-" + s
	}
	return fmt.Sprintf("v%d.%d.%d%s", v.GetMajor(), v.GetMinor(), v.GetPatch(), suffix)
}
