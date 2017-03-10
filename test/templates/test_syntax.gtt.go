package templates

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-qbit/template/filter"
	"github.com/go-qbit/template/test/templates/extwrapper"
	"io"
)

func ProcessTestExprSyntax(ctx context.Context, w io.Writer, i int, b bool, s string, t []TestType, ptr *TestType, buf *bytes.Buffer) {
	extwrapper.WrapperWrapper(ctx, w, func() {
		if i > 0 || i < 100 && !b || i <= 200 && i >= 150 || !(i == 0 && !b) {
			io.WriteString(w, fmt.Sprint(len(s)))
			io.WriteString(w, "   spaces  ")
			io.WriteString(w, "rspaces  ")
			io.WriteString(w, "   lspaces")
		}
		io.WriteString(w, filter.HTML(s[10]))
		io.WriteString(w, fmt.Sprint(t[0].StructField.F1["test"]))
		io.WriteString(w, fmt.Sprint(5+10*15/20%4))
		for i, v := range t {
			io.WriteString(w, fmt.Sprint(i))
			io.WriteString(w, ": ")
			io.WriteString(w, fmt.Sprint(v))
		}
	})
}
