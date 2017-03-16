package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-qbit/template"
)

func main() {
	pkgName := flag.String("package", "template", "Package name")
	flag.Parse()

	for _, pattern := range flag.Args() {
		filesNames, err := filepath.Glob(pattern)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		for _, fileName := range filesNames {
			genFromFile(fileName, *pkgName)
		}
	}
}

func genFromFile(fileName, pkgName string) {
	t := template.New()
	err := t.ParseFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	out, err := os.Create(fileName + ".go")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer out.Close()

	t.WriteGo(out, pkgName, fileName)
}
