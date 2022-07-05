package mongo

import "github.com/rleszilm/genms-datalayer/protoc-gen-genms-datalayer/internal/golang"

type Message struct {
	*golang.Message
	Fields []*Field
}

func NewMessage(m *golang.Message) *Message {
	fields := []*Field{}
	for _, f := range m.Fields {
		fields = append(fields, NewField(f))
	}

	return &Message{
		Message: m,
		Fields:  fields,
	}
}
