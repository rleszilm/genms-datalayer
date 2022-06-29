package protoc

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

// File adds functionality to the underlying File.
type File struct {
	*protogen.File
	*protogen.GeneratedFile
}

// NewFile returns a new File.
func NewFile(file *protogen.File, genfile *protogen.GeneratedFile) *File {
	return &File{
		File:          file,
		GeneratedFile: genfile,
	}
}

// Write writes to file.
func (f *File) Write(p []byte) (int, error) {
	return f.GeneratedFile.Write(p)
}

// QualifiedPackageName adds the import path to the outfile and returns the auto-generated
// alias used for the package.
func (f *File) QualifiedPackageName(path string) string {
	ident := protogen.GoIdent{GoImportPath: protogen.GoImportPath(path)}
	return strings.Split(f.GeneratedFile.QualifiedGoIdent(ident), ".")[0]
}
