package template

import (
	"io"
)

type iAstNode interface {
	GetImports() []string
	GetStrings() []string
	WriteGo(io.Writer, *GenGoOpts)
}

type GenGoOpts struct {
	PackageName string
	Imports     []string
	FileName    string
}
