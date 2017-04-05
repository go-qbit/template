package templates

import (
	"bytes"
	"context"
	"github.com/go-qbit/template/filter"
	"github.com/go-qbit/template/test/templates/extwrapper"
	"github.com/go-qbit/template/utils"
	"io"
)

var (
	s5a3580e377ef6fb1ad2ef729617e4c30 = []byte{0x20, 0x20, 0x20, 0x6C, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73}
	s499a8aee50bd398a3939a0f359d93145 = []byte{0x20, 0x20, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x20, 0x20}
	sd4c4fe320d69ba258449839fc65a4f93 = []byte{0x3A, 0x20}
	s5d98090b412b8b80b94b6c9ffeea34c2 = []byte{0x22, 0x51, 0x75, 0x6F, 0x74, 0x65, 0x64, 0x20, 0x74, 0x65, 0x78, 0x74, 0x22, 0xA, 0x20, 0x20, 0x20, 0x20, 0x57, 0x69, 0x74, 0x68, 0x20, 0x6E, 0x65, 0x77, 0x20, 0x6C, 0x69, 0x6E, 0x65}
	s6869935619bd5c0e40b40d7b054fb833 = []byte{0x72, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x20, 0x20}
)

func ProcessTestExprSyntax(ctx context.Context, w io.Writer, i int, b bool, s string, t []TestType, ptr *TestType, buf *bytes.Buffer) {
	extwrapper.WrapperWrapper(ctx, w, func() {
		if i > 0 || i < 100 && !b || i <= 200 && i >= 150 || !(i == 0 && !b) {
			io.WriteString(w, utils.ToString(len(s)))
			w.Write(s499a8aee50bd398a3939a0f359d93145)
			w.Write(s6869935619bd5c0e40b40d7b054fb833)
			w.Write(s5a3580e377ef6fb1ad2ef729617e4c30)
		}
		for i := 0; i < 10; i++ {
			io.WriteString(w, utils.ToString(i))
		}
		extwrapper.ProcessExtTemplate(ctx, w)
		io.WriteString(w, filter.Filterhtml(s[10]))
		io.WriteString(w, filter.Filterhtml(s[10]))
		io.WriteString(w, utils.ToString(t[0].StructField.F1["test"]))
		io.WriteString(w, utils.ToString(5+10*15/20%4))
		for i, v := range t {
			io.WriteString(w, utils.ToString(i))
			w.Write(sd4c4fe320d69ba258449839fc65a4f93)
			io.WriteString(w, utils.ToString(v))
		}
		io.WriteString(w, utils.ToString(buf.Next(10)))
		a := i
		a = i
		io.WriteString(w, utils.ToString(a))
		Processtest1(ctx, w, "test")
		w.Write(s5d98090b412b8b80b94b6c9ffeea34c2)
	})
}

func Processtest1(ctx context.Context, w io.Writer, s string) {
	io.WriteString(w, filter.Filterhtml(filter.Filterhtml(s)))
}

func Wrapperwrapper1(ctx context.Context, w io.Writer, tplClbF func()) {
	tplClbF()
}

func Wrapperwrapper2(ctx context.Context, w io.Writer, tplClbF func()) {
	Wrapperwrapper1(ctx, w, func() {
		tplClbF()
	})
}
