package pgl

import (
	"github.com/go-test/deep"
	"golang.org/x/tools/imports"
	"google.golang.org/protobuf/compiler/protogen"
)

// Format formats the generated file.
func Format(plugin *protogen.Plugin, name string, out *protogen.GeneratedFile) error {
	original, err := out.Content()
	if err != nil {
		return err
	}

	formatted, err := imports.Process(name, original, nil)
	if err != nil {
		return err
	}

	if diff := deep.Equal(original, formatted); diff != nil {
		out.Skip()

		new := plugin.NewGeneratedFile(name, ".")
		if _, err := new.Write(formatted); err != nil {
			return nil
		}
	}

	return nil
}
