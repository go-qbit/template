package filter

import (
	"html"

	"github.com/go-qbit/template/utils"
)

func Filterhtml(v interface{}) string {
	return html.EscapeString(utils.ToString(v))
}
