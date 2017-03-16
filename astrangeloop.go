package template

import "io"

type astRangeLoop struct {
	indexVariable string
	localVariable string
	loopVariable  iAstNode
	body          iAstNode
}

func (n *astRangeLoop) GetImports() []string {
	res := []string{}

	if n.loopVariable != nil {
		res = append(res, n.loopVariable.GetImports()...)
	}

	if n.body != nil {
		res = append(res, n.body.GetImports()...)
	}

	return res
}

func (n *astRangeLoop) GetStrings() []string {
	res := []string{}

	if n.loopVariable != nil {
		res = append(res, n.loopVariable.GetStrings()...)
	}

	if n.body != nil {
		res = append(res, n.body.GetStrings()...)
	}

	return res
}

func (n *astRangeLoop) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "for "+n.indexVariable+","+n.localVariable+":= range ")
	n.loopVariable.WriteGo(w, opts)
	io.WriteString(w, "{\n")
	n.body.WriteGo(w, opts)
	io.WriteString(w, "}\n")
}
