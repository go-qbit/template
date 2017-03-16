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
	sce845cca7dc3b95da880942e9e8077a8 = []byte{0x22, 0x20, 0x20, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x20, 0x20, 0x22}
	s22d7ca7e4f80646da34a473adf334e2f = []byte{0x22, 0x72, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x20, 0x20, 0x22}
	s0d69e99d6883ff07a2d160f9584fb428 = []byte{0x22, 0x20, 0x20, 0x20, 0x6c, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x22}
	sceff3f4d917917c3d26ac3fe3507d897 = []byte{0x22, 0x3a, 0x20, 0x22}
)

func ProcessTestExprSyntax(ctx context.Context, w io.Writer, i int, b bool, s string, t []TestType, ptr *TestType, buf *bytes.Buffer) {
	extwrapper.WrapperWrapper(ctx, w, func() {
		if i > 0 || i < 100 && !b || i <= 200 && i >= 150 || !(i == 0 && !b) {
			io.WriteString(w, utils.ToString(len(s)))
			w.Write(sce845cca7dc3b95da880942e9e8077a8)
			w.Write(s22d7ca7e4f80646da34a473adf334e2f)
			w.Write(s0d69e99d6883ff07a2d160f9584fb428)
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
			w.Write(sceff3f4d917917c3d26ac3fe3507d897)
			io.WriteString(w, utils.ToString(v))
		}
		io.WriteString(w, utils.ToString(buf.Next(10)))
		a := i
		a = i
		io.WriteString(w, utils.ToString(a))
	})
}
