package golang

import "google.golang.org/protobuf/compiler/protogen"

type File struct {
	*protogen.File
}

func NewFile(f *protogen.File) *File {
	return &File{File: f}
}
