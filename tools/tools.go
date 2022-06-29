package tools

import (
	// imported so tools can be vendored
	_ "github.com/kisielk/errcheck"
	_ "github.com/maxbrunsfeld/counterfeiter/v6"
	_ "github.com/rleszilm/genms-version"
	_ "golang.org/x/lint/golint"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
