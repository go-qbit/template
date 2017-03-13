package templates

import (
	"context"
	"fmt"
	"github.com/go-qbit/template/filter"
	"io"
)

func ProcessTest(ctx context.Context, w io.Writer, header string, users []User) {
	Wrapperpage(ctx, w, func() {
		io.WriteString(w, "<h1>")
		io.WriteString(w, filter.HTML(header))
		io.WriteString(w, "</h1>\n    <hr>")
		for _, user := range users {
			io.WriteString(w, "<p>")
			ProcessUserName(ctx, w, user)
			io.WriteString(w, ":")
			io.WriteString(w, fmt.Sprint(user.Age))
			io.WriteString(w, "(")
			if user.IsMan {
				io.WriteString(w, "Man")
			} else {
				io.WriteString(w, "Woman")
			}
			io.WriteString(w, ")\n        </p>")
		}
	}, header)
}

func ProcessUserName(ctx context.Context, w io.Writer, user User) {
	io.WriteString(w, filter.HTML(user.Name))
	io.WriteString(w, " ")
	io.WriteString(w, filter.HTML(user.Lastname))
}
