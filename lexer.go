package template

import (
	"fmt"
	"regexp"
	"strings"
)

var yyLexDebug = false //|| true

type exprLex struct {
	text    string
	curLine int
	err     error
	result  iAstNode
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
	{"TEMPLATE", TEMPLATE},
	{"WRAPPER", WRAPPER},
	{"PROCESS", PROCESS},
	{"CONTENT", CONTENT_MARKER},
	{"IMPORT", IMPORT},
	{"VARS", VARS},
	{"ELSE", ELSE},
	{"FOR", FOR},
	{"END", END},
	{"USE", USE},
	{"NOT", NOT},
	{"AND", AND},
	{"OR", OR},
	{"IF", IF},
	{"IN", IN},
	{":=", ASSIGNMENT},
	{"==", EQ},
	{"!=", NE},
	{">=", GE},
	{"<=", LE},
	{"&&", AND},
	{"||", OR},
	{"++", INC},
	{"--", DEC},
	{"!", NOT},
}

var reTokens = []reToken{
	{`^(?:\")(?:[^\\\"]*(?:\\.[^\\\"]*)*)(?:\")`, STRING},
	{`^-?\d+(?:[.,]\d+)?`, NUMBER},
	{`^[a-zA-Z_][a-zA-Z0-9\_]*`, IDENTIFIER},
	{`(?s)^/\*.+?\*/`, COMMENT},
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
L1:
	for {
		if len(x.text) == 0 {
			return 0
		}

		for _, token := range simpleTokens {
			if strings.HasPrefix(x.text, token.token) {
				x.text = x.text[len(token.token):]
				yylval.string = token.token

				if yyLexDebug {
					fmt.Println("SIMPLE TOKEN ", tokenName(token.value))
				}

				return token.value
			}
		}

		for _, token := range compiledReTokens {
			if m := token.re.FindString(x.text); m != "" {
				x.text = x.text[len(m):]

				if token.value == COMMENT {
					continue L1
				}

				yylval.string = string(m)

				x.curLine += strings.Count(yylval.string, "\n")

				if yyLexDebug {
					fmt.Println("TOKEN ", tokenName(token.value), yylval.string)
				}

				return token.value
			}
		}

		c := x.text[0]
		x.text = x.text[1:]

		switch c {
		case '\n':
			x.curLine++
			continue
		case ' ', '\t', '\r':
			continue
		default:
			if yyLexDebug {
				fmt.Println("TOKEN ", string(c))
			}
			return int(c)
		}
	}

	return 0
}

func (x *exprLex) Error(s string) {
	x.err = fmt.Errorf("%s at line %d", s, x.curLine+1)
}

func tokenName(c int) string {
	if c > yyPrivate {
		c -= yyPrivate - 1
	}
	if c >= 0 && c < len(yyToknames) {
		if yyToknames[c] != "" {
			return yyToknames[c]
		}
	}
	return fmt.Sprintf("tok-%v", c)

}
