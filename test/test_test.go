package test

import (
	"bytes"
	"os"
	"testing"
)

func TestProcessTest(t *testing.T) {
	buf := &bytes.Buffer{}
	ProcessTest(buf, "Header1", []User{
		{"Ivan", "Sidorov"},
		{"Petr", "Ivanov"},
		{"James", "Bond"},
		{"John", "Connor"},
		{"Sara", "Connor"},
	})

	buf.WriteTo(os.Stdout)
}
