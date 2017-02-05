package template

type iAstNode interface {
}

type astString struct {
	value string
}

type astList struct {
	children []iAstNode
}

type astLoop struct {
	localVariable string
	loopVariable  string
	body          iAstNode
}

type astCondition struct {
	condition iAstNode
	ifBody    iAstNode
	elseBody  iAstNode
}
