package templates

import (
	"fmt"
	"github.com/go-qbit/template/filter"
	"io"
)

func Wrapperpage(w io.Writer, tplClbF func(), caption string) {
	io.WriteString(w, "\n    <!DOCTYPE html>\n    <html>\n    <head>\n        <title>")
	io.WriteString(w, filter.HTML(caption))
	io.WriteString(w, "</title>\n    </head>\n    <body>\n        ")
	tplClbF()
	io.WriteString(w, "\n    </<body>\n    </html>\n")
}

func ProcessTest(w io.Writer, header string, users []User) {
	Wrapperpage(w, func() {
		io.WriteString(w, "\n    <h1>")
		io.WriteString(w, filter.HTML(header))
		io.WriteString(w, "</h1>\n    <hr>\n    ")
		for _, user := range users {
			io.WriteString(w, "\n        <p>\n            ")
			ProcessUserName(w, user)
			io.WriteString(w, ":\n            ")
			io.WriteString(w, fmt.Sprint(user.Age))
			io.WriteString(w, " (")
			if user.IsMan {
				io.WriteString(w, "Man")
			} else {
				io.WriteString(w, "Woman")
			}
			io.WriteString(w, ")\n        </p>\n    ")
		}
		io.WriteString(w, "\n")
	}, header)
}

func ProcessUserName(w io.Writer, user User) {
	io.WriteString(w, filter.HTML(user.Name))
	io.WriteString(w, " ")
	io.WriteString(w, filter.HTML(user.Lastname))
}

func ProcessTestExprSyntax(w io.Writer, i int, b bool, s string) {
	io.WriteString(w, "\n    ")
	if i > 0 || i < 100 && !b || i <= 200 && i >= 150 || !(i == 0 && !b) {
		io.WriteString(w, "\n        ")
		io.WriteString(w, fmt.Sprint(len(s)))
		io.WriteString(w, "\n    ")
	}
	io.WriteString(w, "\n")
}
