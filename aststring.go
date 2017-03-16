package template

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strings"
)

type astString struct {
	value string
}

func (*astString) GetImports() []string {
	return []string{}
}

func (n *astString) GetStrings() []string {
	return []string{strings.Trim(n.value, `"`)}
}

func (n *astString) WriteGo(w io.Writer, opts *GenGoOpts) {
	//io.WriteString(w, strings.Replace(n.value, "\n", `\n`, -1))
	io.WriteString(w, strVarName(strings.Trim(n.value, `"`), opts.FileName))
}

func strVarName(s, filename string) string {
	dgst := md5.Sum([]byte(s))

	return "s" + string(hex.EncodeToString(dgst[:]))
}
