package protoc

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// Message adds functionality to the underlying message.
type Message struct {
	File   *File
	Fields *Fields
	*protogen.Message
}

// NewMessage returns a new Message.
func NewMessage(f *File, message *protogen.Message) *Message {
	m := &Message{
		File:    f,
		Message: message,
	}
	m.Fields = NewFields(f, m)

	return m
}

// GoName returns the name of the message.
func (m *Message) GoName() string {
	return m.GoIdent.GoName
}

// QualifiedGoIdent returns the fully qualified type.
func (m *Message) QualifiedGoIdent() string {
	return m.File.QualifiedGoIdent(m.GoIdent)
}
