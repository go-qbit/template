package test

import (
	"fmt"
	"io"
)

func ProcessTest(w io.Writer, header string, users []User) {
	io.WriteString(w, "\n<!DOCTYPE html>\n<html>\n<body>\n    ")
	io.WriteString(w, fmt.Sprint(header))
	io.WriteString(w, "\n    <hr>\n    ")
	for _, user := range users {
		io.WriteString(w, "\n        <p>")
		io.WriteString(w, fmt.Sprint(user.Name))
		io.WriteString(w, " ")
		io.WriteString(w, fmt.Sprint(user.Lastname))
		io.WriteString(w, "</p>\n    ")
	}
	io.WriteString(w, "\n</<body>\n</html>\n")
}
