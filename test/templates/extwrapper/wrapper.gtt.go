package extwrapper

import (
	"context"
	"io"
)

var (
	s89f0d478f2efe8645effcb383ba348f7 = []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F}
)

func WrapperWrapper(ctx context.Context, w io.Writer, tplClbF func()) {
	tplClbF()
}

func ProcessExtTemplate(ctx context.Context, w io.Writer) {
	w.Write(s89f0d478f2efe8645effcb383ba348f7)
}
