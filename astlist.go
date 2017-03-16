package template

import "io"

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

func (n *astList) GetStrings() []string {
	res := []string{}
	for _, child := range n.children {
		if child != nil {
			res = append(res, child.GetStrings()...)
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
