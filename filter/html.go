package filter

import (
	"html"

	"github.com/go-qbit/template/utils"
)

func HTML(v interface{}) string {
	return html.EscapeString(utils.ToString(v))
}
