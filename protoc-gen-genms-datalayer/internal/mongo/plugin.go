package mongo

import "github.com/rleszilm/genms-datalayer/protoc-gen-genms-datalayer/internal/golang"

type Plugin struct {
	*golang.Plugin
}

func NewPlugin(p *golang.Plugin) *Plugin {
	return &Plugin{
		Plugin: p,
	}
}
