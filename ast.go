package template

import (
	"io"
	"sort"
	"strings"
)

type iAstNode interface {
	GetImports() []string
	WriteGo(io.Writer, *GenGoOpts)
}

type GenGoOpts struct {
	PackageName string
	Imports     []string
}

type astList struct {
	children []iAstNode
}

func (n *astList) GetImports() []string {
	res := []string{}
	for _, child := range n.children {
		if child != nil {
			res = append(res, child.GetImports()...)
		}
	}

	return res
}

func (n *astList) WriteGo(w io.Writer, opts *GenGoOpts) {
	for _, child := range n.children {
		if child != nil {
			child.WriteGo(w, opts)
		}
	}
}

func (n *astList) Add(node ...iAstNode) {
	n.children = append(n.children, node...)
}

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

func (n *astFile) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "package "+opts.PackageName+"\n")

	if n.header == nil {
		n.header = &astHeader{}
	}

	opts.Imports = append(opts.Imports, n.GetImports()...)

	n.header.WriteGo(w, opts)

	if n.macroses != nil {
		n.macroses.WriteGo(w, opts)
	}
}

type astHeader struct {
	imports *astList
}

func (*astHeader) GetImports() []string { return []string{} }

func (n *astHeader) WriteGo(w io.Writer, opts *GenGoOpts) {
	importsMap := make(map[string]struct{}, len(opts.Imports))

	for _, name := range opts.Imports {
		importsMap[name] = struct{}{}
	}

	if n.imports != nil {
		for _, child := range n.imports.children {
			if child != nil {
				importsMap[strings.Trim(child.(*astImport).pkgName, `"`)] = struct{}{}
			}
		}
	}

	imports := make([]string, 0, len(importsMap))
	for name, _ := range importsMap {
		imports = append(imports, name)
	}
	sort.Strings(imports)

	io.WriteString(w, "import (\n")
	for _, name := range imports {
		io.WriteString(w, `"`+name+`"`+"\n")
	}
	io.WriteString(w, ")\n")
}

type astImport struct {
	pkgName string
}

func (*astImport) GetImports() []string { return []string{} }

func (n *astImport) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, n.pkgName+"\n")
}

type astTemplate struct {
	name    string
	vars    *astList
	wrapper *astUseWrapper
	body    *astList
}

func (n *astTemplate) GetImports() []string {
	res := []string{}
	if n.body != nil {
		res = append(res, n.body.GetImports()...)
	}

	return res
}

func (n *astTemplate) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "func Process"+n.name+"(w io.Writer")

	for _, v := range n.vars.children {
		io.WriteString(w, ", ")
		v.WriteGo(w, opts)
	}

	io.WriteString(w, ") {\n")

	if n.wrapper != nil {
		io.WriteString(w, "Wrapper"+n.wrapper.name+"(w, func() {\n")
	}

	if n.body != nil {
		n.body.WriteGo(w, opts)
	}

	if n.wrapper != nil {
		io.WriteString(w, "}")
		for _, p := range n.wrapper.params.children {
			io.WriteString(w, ", ")
			p.WriteGo(w, opts)
		}
		io.WriteString(w, ")\n")
	}

	io.WriteString(w, "}\n\n")
}

type astUseWrapper struct {
	name   string
	params *astList
}

func (*astUseWrapper) GetImports() []string                   { return []string{} }
func (n *astUseWrapper) WriteGo(w io.Writer, opts *GenGoOpts) {}

type astWrapper struct {
	name string
	vars *astList
	body iAstNode
}

func (n *astWrapper) GetImports() []string {
	res := []string{}
	if n.body != nil {
		res = append(res, n.body.GetImports()...)
	}

	return res
}

func (n *astWrapper) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "func Wrapper"+n.name+"(w io.Writer, tplClbF func()")
	for _, v := range n.vars.children {
		io.WriteString(w, ", ")
		v.WriteGo(w, opts)
	}
	io.WriteString(w, ") {\n")

	if n.body != nil {
		n.body.WriteGo(w, opts)
	}

	io.WriteString(w, "}\n\n")
}

type astWriteContent struct{}

func (*astWriteContent) GetImports() []string { return []string{} }
func (n *astWriteContent) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "tplClbF()\n")
}

type astVariableDef struct {
	vName, vType string
}

func (*astVariableDef) GetImports() []string { return []string{} }
func (n *astVariableDef) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, n.vName+" "+n.vType)
}

type astWriteValue struct {
	value iAstNode
}

func (n *astWriteValue) GetImports() []string {
	res := append([]string{"io"}, n.value.GetImports()...)

	if _, ok := n.value.(*astString); !ok {
		res = append(res, "fmt")
	}

	return res
}
func (n *astWriteValue) WriteGo(w io.Writer, opts *GenGoOpts) {
	switch n.value.(type) {
	case *astString:
		io.WriteString(w, "io.WriteString(w, ")
		n.value.WriteGo(w, opts)
		io.WriteString(w, ")\n")
	default:
		io.WriteString(w, "io.WriteString(w, fmt.Sprint(")
		n.value.WriteGo(w, opts)
		io.WriteString(w, "))\n")
	}
}

type astWriteString struct {
	value iAstNode
}

func (n *astWriteString) GetImports() []string { return append([]string{"io"}, n.value.GetImports()...) }
func (n *astWriteString) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "io.WriteString(w, ")
	n.value.WriteGo(w, opts)
	io.WriteString(w, ")\n")
}

type astString struct {
	value string
}

func (*astString) GetImports() []string { return []string{} }
func (n *astString) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, strings.Replace(n.value, "\n", `\n`, -1))
}

type astValue struct {
	name string
}

func (*astValue) GetImports() []string { return []string{} }
func (n *astValue) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, n.name)
}

type astExpr struct {
	operator           string
	operand1, operand2 iAstNode
}

func (n *astExpr) GetImports() []string {
	if n.operand1 != nil {
		return append(n.operand1.GetImports(), n.operand2.GetImports()...)
	} else {
		return n.operand2.GetImports()
	}
}
func (n *astExpr) WriteGo(w io.Writer, opts *GenGoOpts) {
	if n.operand1 != nil {
		n.operand1.WriteGo(w, opts)
	}
	io.WriteString(w, n.operator)
	n.operand2.WriteGo(w, opts)
}

type astParenthesis struct {
	expr iAstNode
}

func (n *astParenthesis) GetImports() []string { return n.expr.GetImports() }
func (n *astParenthesis) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "(")
	n.expr.WriteGo(w, opts)
	io.WriteString(w, ")")
}

type astFunc struct {
	name   string
	params *astList
}

func (*astFunc) GetImports() []string { return []string{} }
func (n *astFunc) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, n.name+"(")
	for i, param := range n.params.children {
		if i > 0 {
			io.WriteString(w, ", ")
		}
		param.WriteGo(w, opts)
	}
	io.WriteString(w, ")")
}

type astFilter struct {
	name  string
	value iAstNode
}

func (n *astFilter) GetImports() []string {
	return append([]string{"github.com/go-qbit/template/filter"}, n.value.GetImports()...)
}
func (n *astFilter) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "filter."+n.name+"(")
	n.value.WriteGo(w, opts)
	io.WriteString(w, ")")
}

type astLoop struct {
	localVariable string
	loopVariable  iAstNode
	body          iAstNode
}

func (n *astLoop) GetImports() []string {
	res := []string{}

	if n.loopVariable != nil {
		res = append(res, n.loopVariable.GetImports()...)
	}

	if n.body != nil {
		res = append(res, n.body.GetImports()...)
	}

	return res
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

func (n *astCondition) GetImports() []string {
	res := []string{}

	if n.condition != nil {
		res = append(res, n.condition.GetImports()...)
	}

	if n.ifBody != nil {
		res = append(res, n.ifBody.GetImports()...)
	}

	if n.elseBody != nil {
		res = append(res, n.elseBody.GetImports()...)
	}

	return res

}
func (n *astCondition) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "if ")
	n.condition.WriteGo(w, opts)
	io.WriteString(w, " {\n")
	n.ifBody.WriteGo(w, opts)
	io.WriteString(w, "}")

	if n.elseBody != nil {
		io.WriteString(w, " else {\n")
		n.elseBody.WriteGo(w, opts)
		io.WriteString(w, "}")
	}

	io.WriteString(w, "\n")
}

type astProcessTemplate struct {
	name   string
	params *astList
}

func (*astProcessTemplate) GetImports() []string { return []string{} }
func (n *astProcessTemplate) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "Process"+n.name+"(w")
	for _, param := range n.params.children {
		io.WriteString(w, ", ")
		param.WriteGo(w, opts)
	}
	io.WriteString(w, ")\n")
}
