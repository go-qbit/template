//go:generate go tool yacc template.y

package template

import (
	"bytes"
	"github.com/go-qbit/qerror"
	"strings"
	"github.com/davecgh/go-spew/spew"
)

type Template struct {
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
			newPos := strings.Index(text[pos+2:], "%]")

			if newPos < 0 {
				return qerror.New("Unclosed [%")
			}

			inBlock = false

			buf.WriteString(text[pos+2 : pos+newPos+2])
			buf.WriteByte(';')

			pos += newPos + 2
		} else {
			newPos := strings.Index(text[pos+2:], "[%")

			if newPos < 0 {
				buf.WriteString(strQuote(text[pos+2:]))
				buf.WriteByte(';')
				break
			}

			inBlock = true
			if pos == 0 {
				buf.WriteString(strQuote(text[:pos+2+newPos]))
			} else {
				buf.WriteString(strQuote(text[pos+2 : pos+2+newPos]))
			}
			buf.WriteByte(';')

			pos += newPos + 2
		}
	}

	//println(buf.String())

	lexer := &exprLex{text: buf.String()}
	yyErrorVerbose = true
	//yyDebug = 5
	yyParse(lexer)

	if lexer.err != nil {
		return lexer.err
	}

	spew.Dump(lexer.result)

	return nil
}

func strQuote(s string) string {
	s = strings.Replace(s, `\`, `\\`, -1)

	return `"` + strings.Replace(s, `"`, `\"`, -1) + `"`
}
