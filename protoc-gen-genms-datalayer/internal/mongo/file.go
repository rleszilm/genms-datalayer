package mongo

import "github.com/rleszilm/genms-datalayer/protoc-gen-genms-datalayer/internal/golang"

type File struct {
	*golang.File
}

func NewFile(f *golang.File) *File {
	return &File{
		File: f,
	}
}
