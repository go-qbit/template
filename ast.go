package template

import (
	"io"
	"strings"
)

type iAstNode interface {
	WriteGo(io.Writer, *GenGoOpts)
}

type GenGoOpts struct {
	PackageName string
	Imports     []string
	TemplateId  string
}

type astString struct {
	value string
}

func (n *astString) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, strings.Replace(n.value, "\n", `\n`, -1))
}

type astValue struct {
	name string
}

func (n *astValue) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, n.name)
}

type astWriteValue struct {
	value iAstNode
}

func (n *astWriteValue) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "io.WriteString(w, fmt.Sprint(")
	n.value.WriteGo(w, opts)
	io.WriteString(w, "))\n")
}

type astWriteString struct {
	value string
}

func (n *astWriteString) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "io.WriteString(w, ")
	s := &astString{n.value}
	s.WriteGo(w, opts)
	io.WriteString(w, ")\n")
}

type astList struct {
	children []iAstNode
}

func (n *astList) WriteGo(w io.Writer, opts *GenGoOpts) {
	for _, child := range n.children {
		if child != nil {
			child.WriteGo(w, opts)
		}
	}
}

type astLoop struct {
	localVariable string
	loopVariable  iAstNode
	body          iAstNode
}

func (n *astLoop) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "for _, "+n.localVariable+":= range ")
	n.loopVariable.WriteGo(w, opts)
	io.WriteString(w, "{\n")
	n.body.WriteGo(w, opts)
	io.WriteString(w, "}\n")
}

type astCondition struct {
	condition iAstNode
	ifBody    iAstNode
	elseBody  iAstNode
}

func (n *astCondition) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "//<CONDITION>\n")
}

type astVariableDef struct {
	vName, vType string
}

func (n *astVariableDef) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, n.vName+" "+n.vType)
}

type astHeader struct {
	imports iAstNode
	vars    iAstNode
}

func (n *astHeader) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "import (\n")
	io.WriteString(w, "\"fmt\"\n")
	io.WriteString(w, "\"io\"\n")
	io.WriteString(w, "\"github.com/go-qbit/template/filter\"\n")
	if n.imports != nil {
		for _, pkgName := range n.imports.(*astList).children {
			pkgName.WriteGo(w, opts)
		}
	}
	io.WriteString(w, ")\n")

	io.WriteString(w, "func Process"+opts.TemplateId+"(w io.Writer")
	for _, v := range n.vars.(*astList).children {
		io.WriteString(w, ","+v.(*astVariableDef).vName+" "+v.(*astVariableDef).vType)
	}
	io.WriteString(w, "){ \n")

}

type astTemplate struct {
	header, exprs iAstNode
}

func (n *astTemplate) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "package "+opts.PackageName+"\n")

	if n.header != nil {
		n.header.WriteGo(w, opts)
	}

	if n.exprs != nil {
		n.exprs.WriteGo(w, opts)
	}

	io.WriteString(w, "}\n")
}

type astImport struct {
	pkgName string
}

func (n *astImport) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, n.pkgName+"\n")
}

type astOutput struct {
	value iAstNode
}

func (n *astOutput) WriteGo(w io.Writer, opts *GenGoOpts) {
	n.value.WriteGo(w, opts)
}

type astFilter struct {
	name  string
	value iAstNode
}

func (n *astFilter) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "filter."+n.name+"(")
	n.value.WriteGo(w, opts)
	io.WriteString(w, ")")
}
