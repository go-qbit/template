package template

import "io"

type astImport struct {
	pkgName string
}

func (*astImport) GetImports() []string { return []string{} }

func (*astImport) GetStrings() []string { return []string{} }

func (n *astImport) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, n.pkgName+"\n")
}
