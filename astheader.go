package template

import (
	"io"
	"sort"
)

type astHeader struct {
	imports *astList
}

func (*astHeader) GetImports() []string { return []string{} }

func (*astHeader) GetStrings() []string { return []string{} }

func (n *astHeader) WriteGo(w io.Writer, opts *GenGoOpts) {
	importsMap := make(map[string]struct{}, len(opts.Imports))

	for _, name := range opts.Imports {
		importsMap[`"`+name+`"`] = struct{}{}
	}

	if n.imports != nil {
		for _, child := range n.imports.children {
			if child != nil {
				importStr := child.(*astImport).pkgName
				if child.(*astImport).alias != "" {
					importStr = child.(*astImport).alias + " " + importStr
				}
				importsMap[importStr] = struct{}{}
			}
		}
	}

	imports := make([]string, 0, len(importsMap))
	for name, _ := range importsMap {
		imports = append(imports, name)
	}
	sort.Strings(imports)

	io.WriteString(w, "import (\n")
	for _, name := range imports {
		io.WriteString(w, name+"\n")
	}
	io.WriteString(w, ")\n")
}
