package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/rleszilm/genms-datalayer/pkg/annotations"
	"github.com/rleszilm/genms-datalayer/protoc-gen-genms-datalayer/internal/golang"
	"github.com/rleszilm/genms-datalayer/protoc-gen-genms-datalayer/internal/mongo"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
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

	goPlugin := golang.NewPlugin(plugin)

	msgByPackage := map[string][]*golang.Message{}
	for _, file := range goPlugin.Files {
		goFile := golang.NewFile(file)

		if _, ok := msgByPackage[string(goFile.GoImportPath)]; !ok {
			msgByPackage[string(goFile.GoImportPath)] = []*golang.Message{}
		}

		for _, msg := range file.Messages {
			options := msg.Desc.Options().(*descriptorpb.MessageOptions)
			opts := proto.GetExtension(options, annotations.E_MessageOptions).(*annotations.Collection)

			goMessage := golang.NewMessage(msg)
			if opts != nil && opts.Generate != annotations.Generate_Skip {
				msgByPackage[string(file.GoImportPath)] = append(msgByPackage[string(file.GoImportPath)], goMessage)

				for _, d := range opts.Datastores {
					switch d {
					case annotations.Datastore_Mongo:
						if err = mongo.GenerateCollection(goPlugin, goFile, goMessage); err != nil {
							return
						}
					}
				}
			}
		}
	}
}
