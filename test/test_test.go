//go:generate go run ../ttgen/main.go -filename templates/index.gtt -out template.go -package test

package test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessTest(t *testing.T) {
	buf := &bytes.Buffer{}
	ProcessTest(buf, "<Header>", []User{
		{"Ivan", "Sidorov"},
		{"Petr", "Ivanov"},
		{"James", "Bond"},
		{"John", "Connor"},
		{"Sara", "Connor"},
	})

	assert.Contains(t, buf.String(), "<h1>&lt;Header&gt;</h1>")
	assert.Contains(t, buf.String(), "<p>James Bond</p>")

	//buf.WriteTo(os.Stdout)
}
