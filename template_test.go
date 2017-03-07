package template_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-qbit/template"
)

func TestTemplate_Parse(t *testing.T) {
	tpl := template.New()

	if !assert.NoError(t, tpl.ParseFile("test/templates/index.gtt")) {
		return
	}

	buf := &bytes.Buffer{}
	tpl.WriteGo(buf, "test")

	assert.Contains(t, buf.String(), `"   spaces  "`)
	assert.Contains(t, buf.String(), `"rspaces  "`)
	assert.Contains(t, buf.String(), `"   lspaces"`)

	//fmt.Println(buf.String())
}
