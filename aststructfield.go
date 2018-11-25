package template

import "io"

type astStructField struct {
	vName iAstNode
	fName string
}

func (*astStructField) GetImports() []string {
	return []string{}
}

func (n *astStructField) GetStrings() []string {
	return n.vName.GetStrings()
}

func (n *astStructField) WriteGo(w io.Writer, opts *GenGoOpts) {
	n.vName.WriteGo(w, opts)
	io.WriteString(w, "."+n.fName)
}
