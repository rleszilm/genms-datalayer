package golang

import "google.golang.org/protobuf/compiler/protogen"

type Plugin struct {
	*protogen.Plugin
}

func NewPlugin(p *protogen.Plugin) *Plugin {
	return &Plugin{Plugin: p}
}
