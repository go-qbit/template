//go:generate go tool yacc template.y

package template

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"strings"

	"github.com/go-qbit/qerror"
)

type Template struct {
	astTree iAstNode
}

func New() *Template {
	return &Template{}
}

func (t *Template) Parse(text string) error {
	buf := bytes.NewBuffer(make([]byte, 0, len(text)*2))

	inBlock := false
	pos := 0

	for pos >= 0 && pos < len(text) {
		if inBlock {
			newPos := strings.Index(text[pos:], "%]")

			if newPos < 0 {
				return qerror.New("Unclosed [%")
			}

			inBlock = false

			buf.WriteString(text[pos : pos+newPos])
			buf.WriteByte(';')

			pos += newPos + 2
		} else {
			newPos := strings.Index(text[pos:], "[%")

			if newPos < 0 {
				buf.WriteString(strQuote(text[pos:]))
				buf.WriteByte(';')
				break
			}

			inBlock = true
			if newPos != 0 {
				buf.WriteString(strQuote(text[pos : pos+newPos]))
				buf.WriteByte(';')
			}

			pos += newPos + 2
		}
	}

	//fmt.Println(buf.String())

	lexer := &exprLex{text: buf.String()}
	yyErrorVerbose = true
	//yyDebug = 5
	yyParse(lexer)

	if lexer.err != nil {
		return lexer.err
	}

	t.astTree = lexer.result
	//spew.Dump(lexer.result)

	return nil
}

func (t *Template) ParseFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("%s: %s", filename, err.Error())
	}

	return t.Parse(string(data))
}

func (t *Template) WriteGo(w io.Writer, packageName string) {
	buf := &bytes.Buffer{}

	t.astTree.WriteGo(buf, &GenGoOpts{
		PackageName: packageName,
	})

	source, err := format.Source(buf.Bytes())
	if err != nil {
		println(buf.String())
		panic(err)
	}

	w.Write(source)
}

func strQuote(s string) string {
	s = strings.Replace(s, `\`, `\\`, -1)

	return `"` + strings.Replace(s, `"`, `\"`, -1) + `"`
}
