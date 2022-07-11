package mongo

import (
	"github.com/rleszilm/genms-datalayer/pkg/annotations/bson"
	"github.com/rleszilm/genms-datalayer/protoc-gen-genms-datalayer/internal/golang"
	"google.golang.org/protobuf/compiler/protogen"
)

type Fields []*Field

type Field struct {
	*golang.Field
}

func (f *Field) MongoKind() string {
	switch f.Opts.GetMongo().GetPrimitive() {
	case bson.Primitive_ObjectID:
		i := protogen.GoIdent{GoName: "ObjectID", GoImportPath: "go.mongodb.org/mongo-driver/bson"}
		return f.GeneratedFile.QualifiedGoIdent(i)
	}

	return f.GoKind()
}

func (f *Field) QueryName() string {
	if f.Opts.GetMongo().GetName() != "" {
		return f.Opts.GetMongo().GetName()
	}
	if f.Opts.GetName() != "" {
		return f.Opts.GetName()
	}
	return f.Desc.TextName()
}

func NewField(f *golang.Field) *Field {
	return &Field{
		Field: f,
	}
}
