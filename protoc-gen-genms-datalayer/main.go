package main

import (
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	var plugin *protogen.Plugin
	var err error

	defer func() {
		if plugin != nil {
			plugin.Error(err)

			out, err := proto.Marshal(plugin.Response())
			if err == nil {
				if _, err := os.Stdout.Write(out); err != nil {
					log.Fatalln(err)
				}
			}
		}

		if err != nil {
			log.Fatalln(err)
		}
	}()

	var buf []byte
	buf, err = ioutil.ReadAll(os.Stdin)
	if err != nil {
		return
	}

	req := &pluginpb.CodeGeneratorRequest{}
	if err = proto.Unmarshal(buf, req); err != nil {
		return
	}

	opts := protogen.Options{}
	plugin, err = opts.New(req)
	if err != nil {
		return
	}
	plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	for _, file := range plugin.Files {
		for _, msg := range file.Messages {
			if err = generate(plugin, file, msg); err != nil {
				return
			}
		}
	}
}

func generate(plugin *protogen.Plugin, file *protogen.File, svc *protogen.Message) error {
	return protocGenGenms.GenerateMicroService(plugin, file, svc, opts)
}
