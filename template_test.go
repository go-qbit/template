package template_test

import (
	"github.com/go-qbit/template"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTemplate_Parse(t *testing.T) {
	tpl := template.New()

	text := `
<!DOCTYPE html>
<html>
	<body>
		[% header %]
		<hr>
		[% FOR user IN users %]
			<p>[% user.Name %]</p>
		[% END %]
	</<body>
</html>
`
	assert.NoError(t, tpl.Parse(text))
}
