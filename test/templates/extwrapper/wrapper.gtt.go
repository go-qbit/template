package extwrapper

import (
	"context"
	"io"
)

func WrapperWrapper(ctx context.Context, w io.Writer, tplClbF func()) {
	tplClbF()
}

func ProcessExtTemplate(ctx context.Context, w io.Writer) {
	io.WriteString(w, "Hello")
}