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

	if !assert.NoError(t, tpl.ParseFile("test/templates/test_syntax.gtt")) {
		return
	}

	buf := &bytes.Buffer{}
	tpl.WriteGo(buf, "test", "")

	assert.Contains(t, buf.String(), `[]byte{0x20, 0x20, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x20, 0x20}`)
	assert.Contains(t, buf.String(), `[]byte{0x72, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x20, 0x20}`)
	assert.Contains(t, buf.String(), `[]byte{0x20, 0x20, 0x20, 0x6c, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73}`)

	//fmt.Println(buf.String())
}
