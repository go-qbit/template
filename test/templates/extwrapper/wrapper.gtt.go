package extwrapper

import (
	"context"
	"io"
)

var (
	s89f0d478f2efe8645effcb383ba348f7 = []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f}
)

func WrapperWrapper(ctx context.Context, w io.Writer, tplClbF func()) {
	tplClbF()
}

func ProcessExtTemplate(ctx context.Context, w io.Writer) {
	w.Write(s89f0d478f2efe8645effcb383ba348f7)
}
