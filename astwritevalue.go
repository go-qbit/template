package template

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

type astWriteValue struct {
	value iAstNode
}

func (n *astWriteValue) GetImports() []string {
	res := append([]string{"io"}, n.value.GetImports()...)

	if _, ok := n.value.(*astString); !ok {
		res = append(res, "github.com/go-qbit/template/utils")
	}

	return res
}

func (n *astWriteValue) GetStrings() []string {
	switch value := n.value.(type) {
	case *astString:
		return []string{value.value[1 : len(value.value)-1]}
	default:
		return n.value.GetStrings()
	}
}

func (n *astWriteValue) WriteGo(w io.Writer, opts *GenGoOpts) {
	switch value := n.value.(type) {
	case *astString:
		io.WriteString(w, "w.Write("+strVarName(value.value[1:len(value.value)-1], opts.FileName)+")\n")
	default:
		io.WriteString(w, "io.WriteString(w, utils.ToString(")
		n.value.WriteGo(w, opts)
		io.WriteString(w, "))\n")
	}
}

func strVarName(s, filename string) string {
	dgst := md5.Sum([]byte(filename + s))

	return "s" + string(hex.EncodeToString(dgst[:]))
}
