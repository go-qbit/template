package template

import "io"

type astVariableDef struct {
	vName, vType string
}

func (*astVariableDef) GetImports() []string {
	return []string{}
}

func (*astVariableDef) GetStrings() []string {
	return []string{}
}

func (n *astVariableDef) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, n.vName+" "+n.vType)
}
