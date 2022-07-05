package golang

import (
	"fmt"

	"github.com/rleszilm/genms-datalayer/pkg/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Field struct {
	*protogen.Field
	GeneratedFile *GeneratedFile
	Opts          *annotations.Field
}

func (f *Field) WithGeneratedFile(gf *GeneratedFile) {
	f.GeneratedFile = gf
}

func (f *Field) GoKind() (goType string) {
	if f.Desc.IsWeak() {
		return "struct{}"
	}

	pointer := f.Desc.HasPresence()
	switch f.Desc.Kind() {
	case protoreflect.BoolKind:
		goType = "bool"
	case protoreflect.EnumKind:
		goType = f.GeneratedFile.QualifiedGoIdent(f.Enum.GoIdent)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		goType = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		goType = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		goType = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		goType = "uint64"
	case protoreflect.FloatKind:
		goType = "float32"
	case protoreflect.DoubleKind:
		goType = "float64"
	case protoreflect.StringKind:
		goType = "string"
	case protoreflect.BytesKind:
		goType = "[]byte"
		pointer = false
	case protoreflect.MessageKind, protoreflect.GroupKind:
		goType = "*" + f.GeneratedFile.QualifiedGoIdent(f.Message.GoIdent)
		pointer = false
	}

	switch {
	case f.Desc.IsList():
		goType = "[]" + goType
		pointer = false
	case f.Desc.IsMap():
		keyField := &Field{Field: f.Message.Fields[0], GeneratedFile: f.GeneratedFile, Opts: f.Opts}
		valField := &Field{Field: f.Message.Fields[1], GeneratedFile: f.GeneratedFile, Opts: f.Opts}
		goType = fmt.Sprintf("map[%v]%v", keyField.GoKind(), valField.GoKind())
		pointer = false
	}

	if pointer {
		return "*" + goType
	}
	return goType
}

func (f *Field) QualifiedGoKind() string {
	return f.GeneratedFile.QualifiedGoIdent(f.GoIdent)
}

func NewField(f *protogen.Field, opts *annotations.Field) *Field {
	return &Field{
		Field: f,
		Opts:  opts,
	}
}
