//go:generate goyacc template.y

package template

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/go-qbit/qerror"
)

var trimRe = regexp.MustCompile(`(?s:^(\s*)(.*?)(\s*)$)`)

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
				return qerror.Errorf("Unclosed [%")
			}

			inBlock = false

			from := pos
			to := pos + newPos
			if text[pos] == '+' {
				from++
			}
			if pos+newPos-1 > 0 && text[pos+newPos-1] == '+' {
				to--
			}
			buf.WriteString(text[from:to])
			buf.WriteByte(';')

			pos += newPos + 2
		} else {
			newPos := strings.Index(text[pos:], "[%")

			if newPos < 0 {
				parts := trimRe.FindStringSubmatch(text[pos:])
				buf.WriteString(parts[1])
				if len(parts[2]) > 0 {
					buf.WriteString(strQuote(parts[2]))
				}
				buf.WriteString(parts[3])
				if len(parts[2]) > 0 {
					buf.WriteByte(';')
				}
				break
			}

			inBlock = true
			if newPos != 0 {
				parts := trimRe.FindStringSubmatch(text[pos : pos+newPos])
				if pos-3 > 0 && text[pos-3] == '+' {
					parts[2] = parts[1] + parts[2]
					parts[1] = ""
				}

				if pos+newPos+2 < len(text) && text[pos+newPos+2] == '+' {
					parts[2] = parts[2] + parts[3]
					parts[3] = ""
				}

				buf.WriteString(parts[1])
				if len(parts[2]) > 0 {
					buf.WriteString(strQuote(parts[2]))
				}
				buf.WriteString(parts[3])
				if len(parts[2]) > 0 {
					buf.WriteByte(';')
				}
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

	err = t.Parse(string(data))
	if err != nil {
		return fmt.Errorf("%s: %s", filename, err.Error())
	}

	return nil
}

func (t *Template) WriteGo(w io.Writer, packageName, fileNme string) {
	buf := &bytes.Buffer{}

	t.astTree.WriteGo(buf, &GenGoOpts{
		PackageName: packageName,
		FileName:    fileNme,
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
