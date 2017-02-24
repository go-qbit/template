package filter

import (
	"fmt"
	"html"
)

func HTML(v interface{}) string {
	return html.EscapeString(fmt.Sprint(v))
}
