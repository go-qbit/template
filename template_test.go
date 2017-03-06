package template_test

import (
	"bytes"
	"testing"

	"github.com/go-qbit/template"

	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestTemplate_Parse(t *testing.T) {
	tpl := template.New()

	if !assert.NoError(t, tpl.ParseFile("test/templates/index.gtt")) {
		return
	}

	buf := &bytes.Buffer{}
	tpl.WriteGo(buf, "test")

	fmt.Println(buf.String())
}
