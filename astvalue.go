package template

import "io"

type astValue struct {
	name string
}

func (*astValue) GetImports() []string {
	return []string{}
}

func (*astValue) GetStrings() []string {
	return []string{}
}

func (n *astValue) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, n.name)
}
