package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-qbit/template"
)

func main() {
	fileName := flag.String("filename", "", "Path to template")
	outFileName := flag.String("out", "template.go", "Write to file")
	pkgName := flag.String("package", "template", "Package name")

	flag.Parse()
	if *fileName == "" {
		fmt.Println("-filename is missed")
		os.Exit(1)
	}

	t := template.New()
	err := t.ParseFile(*fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	out, err := os.Create(*outFileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer out.Close()

	t.WriteGo(out, *pkgName, "Test")
}
