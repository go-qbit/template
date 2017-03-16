//go:generate go run ../ttgen/main.go -package templates templates/*.gtt
//go:generate go run ../ttgen/main.go -package extwrapper templates/extwrapper/*.gtt

package test

import (
	"bytes"
	"context"
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-qbit/template/test/templates"
)

const coreTemplate = `
{{define "header"}}
	<!DOCTYPE html>
	<html>
	<head>
        <title>{{ . }}</title>
    </head>
	<body>
{{ end }}

{{define "footer"}}{{ .Name }} {{ .Lastname }}{{ end }}

{{define "UserName"}}{{ .Name }} {{ .Lastname }}{{ end }}

	{{template "header" .Header}}
    <h1>{{ .Header }}</h1>
    <hr>
    {{range .Users }}
        <p>
        	{{template "UserName" .}}:
        	{{ .Age }} ({{if .IsMan }}Man{{else}}Woman{{end}})
        </p>
    {{ end }}
    {{template "footer"}}
`

var coreTpl *template.Template
var BenchmarkUsers []templates.User
var buf *bytes.Buffer

func init() {
	buf = bytes.NewBuffer(make([]byte, 100*1024))

	var err error
	coreTpl, err = template.New("test").Parse(coreTemplate)
	if err != nil {
		panic(err)
	}

	BenchmarkUsers = make([]templates.User, 1000)
	for i := 0; i < 1000; i++ {
		BenchmarkUsers[i] = templates.User{"Name", "Lastname", 32, true}
	}

	buf := &bytes.Buffer{}
	coreTpl.Execute(buf, struct {
		Header string
		Users  []templates.User
	}{"<Header>", BenchmarkUsers})
	//buf.WriteTo(os.Stdout)

}

func TestProcessTest(t *testing.T) {
	buf := &bytes.Buffer{}
	templates.ProcessTest(context.Background(), buf, "<Header>", []templates.User{
		{"Ivan", "Sidorov", 20, true},
		{"Petr", "Ivanov", 30, true},
		{"James", "Bond", 40, true},
		{"John", "Connor", 50, true},
		{"Sara", "Connor", 60, false},
	})

	assert.Contains(t, buf.String(), "<h1>&lt;Header&gt;</h1>")
	assert.Regexp(t, `<p>\s*James Bond:\s*40\s*\(Man\)\s*</p>`, buf.String())

	//buf.WriteTo(os.Stdout)
}

func BenchmarkGoCoreTemplate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf.Reset()
		coreTpl.Execute(buf, struct {
			Header string
			Users  []templates.User
		}{"<Header>", BenchmarkUsers})
	}
}

func BenchmarkQBitTemplate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf.Reset()
		templates.ProcessTest(context.Background(), buf, "<Header>", BenchmarkUsers)
	}
}
