package template

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type exprLex struct {
	text string
	err  error
}

type simpleToken struct {
	token string
	value int
}

type reToken struct {
	re    string
	value int
}

type compiledReToken struct {
	re    *regexp.Regexp
	value int
}

var simpleTokens = []simpleToken{
	{"IF", IF},
	{"IN", IN},
	{"FOR", FOR},
	{"ELSE", ELSE},
	{"END", END},
}

var reTokens = []reToken{
	{`^(?:\")(?:[^\\\"]*(?:\\.[^\\\"]*)*)(?:\")`, STRING},
	{`^-?\d+(?:[.,]\d+)?`, NUMBER},
	{`^\*?[a-zA-Z_][a-zA-Z0-9\._]+`, IDENTIFIER},
}

var compiledReTokens = getCompiledReTokens()

func getCompiledReTokens() []compiledReToken {
	result := make([]compiledReToken, len(reTokens))
	var err error
	for i := range reTokens {
		result[i].re, err = regexp.Compile(reTokens[i].re)
		if err != nil {
			panic(err)
		}
		result[i].value = reTokens[i].value
	}
	return result
}

func (x *exprLex) Lex(yylval *yySymType) int {
	for {
		if len(x.text) == 0 {
			return 0
		}

		for _, token := range simpleTokens {
			if strings.HasPrefix(x.text, token.token) {
				x.text = x.text[len(token.token):]
				yylval.string = token.token

				fmt.Println("SIMPLE TOKEN ", tokenName(token.value))

				return token.value
			}
		}

		for _, token := range compiledReTokens {
			if m := token.re.FindString(x.text); m != "" {
				x.text = x.text[len(m):]
				yylval.string = string(m)

				fmt.Println("TOKEN ", tokenName(token.value), yylval.string)

				return token.value
			}
		}

		c := x.text[0]
		x.text = x.text[1:]

		switch c {
		case ' ', '\t', '\n', '\r':
			continue
		default:
			fmt.Println("TOKEN ", string(c))
			return int(c)
		}
	}

	return 0
}

func (x *exprLex) Error(s string) {
	x.err = errors.New(s)
}

func tokenName(c int) string {
	if c > yyPrivate {
		c -= yyPrivate -1
	}
	if c >= 0 && c < len(yyToknames) {
		if yyToknames[c] != "" {
			return yyToknames[c]
		}
	}
	return fmt.Sprintf("tok-%v", c)

}
