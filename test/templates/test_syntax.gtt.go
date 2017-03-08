package templates

import (
	"bytes"
	"fmt"
	"github.com/go-qbit/template/filter"
	"io"
)

func ProcessTestExprSyntax(w io.Writer, i int, b bool, s string, t []TestType, ptr *TestType, buf *bytes.Buffer) {
	if i > 0 || i < 100 && !b || i <= 200 && i >= 150 || !(i == 0 && !b) {
		io.WriteString(w, fmt.Sprint(len(s)))
		io.WriteString(w, "   spaces  ")
		io.WriteString(w, "rspaces  ")
		io.WriteString(w, "   lspaces")
	}
	io.WriteString(w, filter.HTML(s[10]))
	io.WriteString(w, fmt.Sprint(t[0].StructField.F1["test"]))
	io.WriteString(w, fmt.Sprint(5+10*15/20))
	for i, v := range t {
		io.WriteString(w, fmt.Sprint(i))
		io.WriteString(w, ": ")
		io.WriteString(w, fmt.Sprint(v))
	}
}
