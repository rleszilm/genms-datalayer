package mongo

import (
	"bytes"
	"fmt"
	"path"
	"strings"
	"text/template"

	"github.com/go-test/deep"
	"github.com/rleszilm/genms-datalayer/protoc-gen-genms-datalayer/internal/golang"
	"github.com/rleszilm/genms-datalayer/protoc-gen-genms-datalayer/internal/pgd"
	"golang.org/x/tools/imports"
	"google.golang.org/protobuf/compiler/protogen"
)

// GenerateCollection generates the code for a collection.
func GenerateCollection(p *golang.Plugin, f *golang.File, m *golang.Message) error {
	mgoPlugin := NewPlugin(p)
	mgoFile := NewFile(f)
	mgoMessage := NewMessage(m)

	dir := path.Dir(f.GeneratedFilenamePrefix)
	filename := path.Join(dir, fmt.Sprintf("dal/mongo/%s.dal.go", strings.ToLower(mgoMessage.GoKind())))

	goGeneratedFile := golang.NewGeneratedFile(mgoPlugin.NewGeneratedFile(filename, mgoFile.GoImportPath), filename)
	mgoMessage.WithGeneratedFile(goGeneratedFile)

	generator := &generator{
		Plugin:        mgoPlugin,
		File:          mgoFile,
		GeneratedFile: goGeneratedFile,
		Message:       mgoMessage,
	}

	return generator.render()
}

type generator struct {
	Plugin        *Plugin
	File          *File
	GeneratedFile *golang.GeneratedFile
	Message       *Message
}

func (g *generator) QualifiedPackageName(path string) string {
	i := protogen.GoIdent{GoImportPath: protogen.GoImportPath(path)}
	return strings.Split(g.GeneratedFile.QualifiedGoIdent(i), ".")[0]
}

func (g *generator) render() error {
	steps := []func() error{
		g.definePackage,
		g.defineConfig,
		g.defineMongoStructs,
		/*
			g.defineCollection,
			g.defineService,
			g.defineDefaultQueries,
			g.defineQueries,
			g.defineNewCollection,
		*/
	}

	for _, s := range steps {
		if err := s(); err != nil {
			return err
		}
	}

	original, err := g.GeneratedFile.Content()
	if err != nil {
		return err
	}

	formatted, err := imports.Process(g.GeneratedFile.Filename, original, nil)
	if err != nil {
		return err
	}

	if diff := deep.Equal(original, formatted); diff != nil {
		formattedOutfile := g.Plugin.NewGeneratedFile(g.GeneratedFile.Filename, ".")
		if _, err := formattedOutfile.Write(formatted); err != nil {
			return err
		}
		g.GeneratedFile.Skip()
	}

	return nil
}

func (g *generator) definePackage() error {
	tmplSrc := `// Package dalmongo is generated by protoc-gen-go-dal. *DO NOT EDIT*
package dalmongo
`

	tmpl, err := template.New("defineMongoPackage").
		Funcs(template.FuncMap{}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"G": g,
	}); err != nil {
		return err
	}

	if _, err := g.GeneratedFile.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (g *generator) defineConfig() error {
	tmplSrc := `// $.G.Message.GoKind }}Config is a struct that can be used to configure a {{ $.G.Message.GoKind }}Collection
type {{ $.G.Message.GoKind }}Config struct {
	Name string ` + "`" + `envconfig:"name" default:"dal-{{ tolower $.G.Message.GoKind }}"` + "`" + `
	Database string ` + "`" + `envconfig:"database" default:"vvv-repl"` + "`" + `
	Collection string ` + "`" + `envconfig:"collection" default:"{{ tolower $.G.Message.GoKind }}"` + "`" + `
	Timeout {{ $.P.Time }}.Duration ` + "`" + `envconfig:"timeout" default:"5s"` + "`" + `
}
`

	tmpl, err := template.New("defineMongoConfig").
		Funcs(template.FuncMap{
			"tolower": strings.ToLower,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Time": g.QualifiedPackageName("time"),
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"G": g,
		"P": p,
	}); err != nil {
		return err
	}

	if _, err := g.GeneratedFile.Write(buf.Bytes()); err != nil {
		return err
	}

	return nil
}

func (g *generator) defineMongoStructs() error {
	tmplSrc := `/*

{{ $.G.Message.GoKind }} {{ $.G.Message.QualifiedGoKind }}
ignore || name || query name || go kind || qualified go kind || qualified mongo kind
{{ range $n, $f := $.G.Message.Fields -}}
	{{ $f.Ignore }} || {{ $f.GoName }} || {{ $f.QueryName }} || {{ $f.GoKind }} || {{ $f.QualifiedGoKind }} || {{ $f.QualifiedMongoKind }} 
{{ end -}}
*/`

	tmpl, err := template.New("defineMongoStructs").
		Funcs(template.FuncMap{
			"tolower":     strings.ToLower,
			"totitlecase": pgd.ToTitleCase,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"bson": g.QualifiedPackageName("go.mongodb.org/mongo-driver/bson"),
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"G": g,
		"P": p,
	}); err != nil {
		return err
	}

	if _, err := g.GeneratedFile.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}