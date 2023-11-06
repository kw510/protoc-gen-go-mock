package main

import (
	"flag"
	"fmt"

	"github.com/kw510/protoc-gen-go-mock/internal/handlers"
	"github.com/kw510/protoc-gen-go-mock/internal/testify"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const version = "0.1.0"

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-mock %v\n", version)
		return
	}

	var flags flag.FlagSet
	framework := flags.String("framework", "", "set the framework")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		var fn func(*protogen.Plugin, *protogen.File, string) *protogen.GeneratedFile
		switch *framework {
		case "testify":
			fn = testify.GenerateFile
		case "handlers":
			fn = handlers.GenerateFile
		default:
			fn = testify.GenerateFile
		}
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			fn(gen, f, version)
		}
		return nil
	})
}
