package protoc

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// Field adds functionality to the underlying Field.
type Field struct {
	File    *File
	Message *Message
	*protogen.Field
}

// NewField returns a new Field.
func NewField(f *File, m *Message, field *protogen.Field) *Field {
	return &Field{
		File:    f,
		Message: m,
		Field:   field,
	}
}

// Name returns the name of the field.
func (f *Field) Name() string {
	return string(f.Desc.Name())
}

// GoIdent returns the fields go type.
func (f *Field) GoIdent() string {
	if f.Message != nil {
		return f.Message.GoIdent.GoName
	}
	if f.Enum != nil {
		return f.Enum.GoIdent.GoName
	}
	return ""
}

// QualifiedGoIdent returns the fully qualified kind.
func (f *Field) QualifiedGoIdent() string {
	if f.Message != nil {
		return f.File.QualifiedGoIdent(f.Message.GoIdent)
	}
	if f.Enum != nil {
		return f.File.QualifiedGoIdent(f.Enum.GoIdent)
	}
	return ""
}
