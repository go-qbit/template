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
	sa86183743bc7d15fab82684be5946930 = []byte{0x20, 0x20, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x20, 0x20}
	s51ab1aa55be8d44f72c53c1a2a0787ca = []byte{0x72, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x20, 0x20}
	sca818fd8734e8b9278f32876fc4b0fc5 = []byte{0x20, 0x20, 0x20, 0x6c, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73}
	se265492707dd0595fd1f399bb8a2690e = []byte{0x3a, 0x20}
)

func ProcessTestExprSyntax(ctx context.Context, w io.Writer, i int, b bool, s string, t []TestType, ptr *TestType, buf *bytes.Buffer) {
	extwrapper.WrapperWrapper(ctx, w, func() {
		if i > 0 || i < 100 && !b || i <= 200 && i >= 150 || !(i == 0 && !b) {
			io.WriteString(w, utils.ToString(len(s)))
			w.Write(sa86183743bc7d15fab82684be5946930)
			w.Write(s51ab1aa55be8d44f72c53c1a2a0787ca)
			w.Write(sca818fd8734e8b9278f32876fc4b0fc5)
		}
		for i := 0; i < 10; i++ {
			io.WriteString(w, utils.ToString(i))
		}
		extwrapper.ProcessExtTemplate(ctx, w)
		io.WriteString(w, filter.HTML(s[10]))
		io.WriteString(w, utils.ToString(t[0].StructField.F1["test"]))
		io.WriteString(w, utils.ToString(5+10*15/20%4))
		for i, v := range t {
			io.WriteString(w, utils.ToString(i))
			w.Write(se265492707dd0595fd1f399bb8a2690e)
			io.WriteString(w, utils.ToString(v))
		}
		io.WriteString(w, utils.ToString(buf.Next(10)))
		a := i
		a = i
		io.WriteString(w, utils.ToString(a))
	})
}
