package golang

import (
	"github.com/rleszilm/genms-datalayer/pkg/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type Message struct {
	*protogen.Message
	Fields        []*Field
	GeneratedFile *GeneratedFile
	Opts          *annotations.Collection
}

func (m *Message) WithGeneratedFile(gf *GeneratedFile) {
	m.GeneratedFile = gf

	for _, f := range m.Fields {
		f.WithGeneratedFile(gf)
	}
}

func (m *Message) GoKind() string {
	return m.GoIdent.GoName
}

func (m *Message) QualifiedGoKind() string {
	return m.GeneratedFile.QualifiedGoIdent(m.GoIdent)
}

func NewMessage(msg *protogen.Message) *Message {
	fields := []*Field{}
	for _, f := range msg.Fields {
		options := f.Desc.Options().(*descriptorpb.FieldOptions)
		opts := proto.GetExtension(options, annotations.E_FieldOptions).(*annotations.Field)
		fields = append(fields, &Field{Field: f, Opts: opts})
	}

	options := msg.Desc.Options().(*descriptorpb.MessageOptions)
	opts := proto.GetExtension(options, annotations.E_MessageOptions).(*annotations.Collection)

	return &Message{
		Message: msg,
		Fields:  fields,
		Opts:    opts,
	}
}
