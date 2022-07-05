package golang

import "google.golang.org/protobuf/compiler/protogen"

type GeneratedFile struct {
	*protogen.GeneratedFile
	Filename string
}

func NewGeneratedFile(gf *protogen.GeneratedFile, fn string) *GeneratedFile {
	return &GeneratedFile{
		GeneratedFile: gf,
		Filename:      fn,
	}
}
