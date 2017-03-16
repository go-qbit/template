package extwrapper

import (
	"context"
	"io"
)

var (
	s8b1a9953c4611296a827abf8c47804d7 = []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f}
)

func WrapperWrapper(ctx context.Context, w io.Writer, tplClbF func()) {
	tplClbF()
}

func ProcessExtTemplate(ctx context.Context, w io.Writer) {
	w.Write(s8b1a9953c4611296a827abf8c47804d7)
}
