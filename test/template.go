package test

import (
	"fmt"
	"github.com/go-qbit/template/filter"
	"io"
)

func ProcessTest(w io.Writer, header string, users []User) {
	io.WriteString(w, "\n<!DOCTYPE html>\n<html>\n<body>\n    <h1>")
	io.WriteString(w, fmt.Sprint(filter.HTML(header)))
	io.WriteString(w, "</h1>\n    <hr>\n    ")
	for _, user := range users {
		io.WriteString(w, "\n        <p>")
		io.WriteString(w, fmt.Sprint(filter.HTML(user.Name)))
		io.WriteString(w, " ")
		io.WriteString(w, fmt.Sprint(filter.HTML(user.Lastname)))
		io.WriteString(w, ": ")
		io.WriteString(w, fmt.Sprint(user.Age))
		io.WriteString(w, "</p>\n    ")
	}
	io.WriteString(w, "\n</<body>\n</html>\n")
}
