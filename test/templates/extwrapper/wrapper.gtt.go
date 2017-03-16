package extwrapper

import (
	"context"
	"io"
)

var (
	sde8870e0a51013c257e3752f33feae93 = []byte{0x22, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x22}
)

func WrapperWrapper(ctx context.Context, w io.Writer, tplClbF func()) {
	tplClbF()
}

func ProcessExtTemplate(ctx context.Context, w io.Writer) {
	w.Write(sde8870e0a51013c257e3752f33feae93)
}
