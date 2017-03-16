package template

import (
	"fmt"
	"io"
)

type astFile struct {
	header, macroses iAstNode
}

func (n *astFile) GetImports() []string {
	res := []string{}
	if n.macroses != nil {
		res = append(res, n.macroses.GetImports()...)
	}

	return res
}

func (n *astFile) GetStrings() []string {
	res := []string{}
	if n.macroses != nil {
		res = append(res, n.macroses.GetStrings()...)
	}

	return res
}

func (n *astFile) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "package "+opts.PackageName+"\n")

	if n.header == nil {
		n.header = &astHeader{}
	}

	opts.Imports = append(opts.Imports, n.GetImports()...)

	n.header.WriteGo(w, opts)

	stringsSet := make(map[string]struct{})
	for _, s := range n.GetStrings() {
		stringsSet[s] = struct{}{}
	}

	if len(stringsSet) > 0 {
		io.WriteString(w, "var (\n")
		for s := range stringsSet {
			io.WriteString(w, strVarName(s, opts.FileName)+" = "+fmt.Sprintf("%#v", []byte(s))+"\n")
		}
		io.WriteString(w, ")\n\n")
	}

	if n.macroses != nil {
		n.macroses.WriteGo(w, opts)
	}
}
