package template

import "io"

type astWriteContent struct{}

func (*astWriteContent) GetImports() []string {
	return []string{}
}

func (*astWriteContent) GetStrings() []string {
	return []string{}
}

func (n *astWriteContent) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "tplClbF()\n")
}
