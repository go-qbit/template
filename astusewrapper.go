package template

import "io"

type astUseWrapper struct {
	pkgName string
	name    string
	params  *astList
}

func (*astUseWrapper) GetImports() []string {
	return []string{}
}

func (*astUseWrapper) GetStrings() []string {
	return []string{}
}

func (n *astUseWrapper) WriteGo(w io.Writer, opts *GenGoOpts) {}
