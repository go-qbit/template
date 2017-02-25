//line template.y:2
package template

import (
	__yyfmt__ "fmt"
)

//line template.y:9
type yySymType struct {
	yys       int
	string    string
	iAstNode  iAstNode
	astFilter *astFilter
}

const IDENTIFIER = 57346
const STRING = 57347
const NUMBER = 57348
const FOR = 57349
const IN = 57350
const IF = 57351
const ELSE = 57352
const END = 57353
const VARS = 57354
const IMPORT = 57355

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENTIFIER",
	"STRING",
	"NUMBER",
	"FOR",
	"IN",
	"IF",
	"ELSE",
	"END",
	"VARS",
	"IMPORT",
	"'('",
	"')'",
	"';'",
	"','",
	"'['",
	"']'",
	"'|'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line template.y:82

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 32
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 78

var yyAct = [...]int{

	7, 34, 6, 19, 21, 18, 28, 53, 44, 41,
	36, 42, 37, 62, 59, 58, 51, 50, 27, 49,
	38, 46, 45, 32, 9, 8, 30, 12, 17, 13,
	56, 55, 48, 24, 5, 40, 39, 16, 9, 8,
	15, 12, 47, 13, 52, 63, 31, 22, 23, 26,
	57, 27, 54, 9, 8, 35, 12, 29, 13, 27,
	60, 20, 61, 27, 9, 8, 11, 12, 10, 13,
	33, 14, 25, 4, 3, 2, 1, 43,
}
var yyPact = [...]int{

	21, -1000, -1000, 60, 28, 23, 12, -1000, -15, -17,
	-1000, -1000, 57, 43, -1000, 19, 44, 60, 53, 53,
	38, 7, -1000, -1000, 51, -5, -1000, -1000, -1000, 6,
	-1000, 43, 60, -6, -1000, 4, 5, 37, 17, 3,
	1, 0, 51, -1000, -1000, -12, -1000, -1000, -1000, 60,
	20, -1000, -1000, 46, -1, -1000, -2, -1000, 49, 60,
	-1000, -3, 34, -1000,
}
var yyPgo = [...]int{

	0, 77, 76, 75, 74, 73, 72, 71, 70, 1,
	2, 0, 4, 68, 66, 6,
}
var yyR1 = [...]int{

	0, 2, 3, 4, 5, 5, 6, 6, 7, 7,
	8, 8, 9, 9, 1, 1, 10, 10, 11, 11,
	11, 11, 11, 11, 11, 12, 12, 15, 15, 13,
	14, 14,
}
var yyR2 = [...]int{

	0, 1, 2, 2, 0, 5, 1, 3, 0, 5,
	1, 3, 0, 2, 1, 3, 1, 3, 0, 1,
	1, 3, 3, 1, 1, 1, 1, 1, 3, 8,
	6, 10,
}
var yyChk = [...]int{

	-1000, -2, -3, -4, -5, 13, -10, -11, 5, 4,
	-13, -14, 7, 9, -7, 12, 14, 16, 20, 20,
	4, -12, 4, 5, 14, -6, 5, -11, -15, 4,
	-15, 8, 16, -8, -9, 4, 15, 17, 14, -12,
	-10, 15, 17, -1, 4, 18, 16, 5, 15, 16,
	16, 16, -9, 19, -10, 11, 10, 4, 16, 16,
	11, -10, 16, 11,
}
var yyDef = [...]int{

	4, -2, 1, 18, 8, 0, 2, 16, 19, 20,
	23, 24, 0, 0, 3, 0, 0, 18, 0, 0,
	0, 0, 25, 26, 12, 0, 6, 17, 22, 27,
	21, 0, 18, 0, 10, 0, 0, 0, 0, 0,
	0, 0, 12, 13, 14, 0, 5, 7, 28, 18,
	18, 9, 11, 0, 0, 30, 0, 15, 18, 18,
	29, 0, 18, 31,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	14, 15, 3, 3, 17, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 16,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 18, 3, 19, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 20,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:32
		{
			yylex.(*exprLex).result = yyVAL.iAstNode
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:34
		{
			yyVAL.iAstNode = &astTemplate{yyDollar[1].iAstNode, yyDollar[2].iAstNode}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:36
		{
			yyVAL.iAstNode = &astHeader{yyDollar[1].iAstNode, yyDollar[2].iAstNode}
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:38
		{
			yyVAL.iAstNode = nil
		}
	case 5:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:39
		{
			yyVAL.iAstNode = yyDollar[3].iAstNode
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:41
		{
			yyVAL.iAstNode = &astList{[]iAstNode{&astImport{yyDollar[1].string}}}
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:42
		{
			yyVAL.iAstNode.(*astList).children = append(yyVAL.iAstNode.(*astList).children, &astImport{yyDollar[3].string})
		}
	case 8:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:44
		{
			yyVAL.iAstNode = nil
		}
	case 9:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:45
		{
			yyVAL.iAstNode = yyDollar[3].iAstNode
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:47
		{
			yyVAL.iAstNode = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:48
		{
			yyVAL.iAstNode.(*astList).children = append(yyVAL.iAstNode.(*astList).children, yyDollar[3].iAstNode)
		}
	case 12:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:50
		{
			yyVAL.iAstNode = nil
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:51
		{
			yyVAL.iAstNode = &astVariableDef{yyDollar[1].string, yyDollar[2].string}
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:53
		{
			yyVAL.string = yyDollar[1].string
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:54
		{
			yyVAL.string = "[]" + yyDollar[3].string
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:56
		{
			yyVAL.iAstNode = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:57
		{
			yyVAL.iAstNode.(*astList).children = append(yyVAL.iAstNode.(*astList).children, yyDollar[3].iAstNode)
		}
	case 18:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:59
		{
			yyVAL.iAstNode = nil
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:60
		{
			yyVAL.iAstNode = &astWriteString{&astString{yyDollar[1].string}}
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:61
		{
			yyVAL.iAstNode = &astWriteValue{&astValue{yyDollar[1].string}}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:62
		{
			yyDollar[3].astFilter.value = &astValue{yyDollar[1].string}
			yyVAL.iAstNode = &astWriteString{yyDollar[3].astFilter}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:63
		{
			yyDollar[3].astFilter.value = &astString{yyDollar[1].string}
			yyVAL.iAstNode = &astWriteString{yyDollar[3].astFilter}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:64
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:65
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:68
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:69
		{
			yyVAL.iAstNode = &astString{yyDollar[1].string}
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:71
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:72
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 29:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line template.y:76
		{
			yyVAL.iAstNode = &astLoop{yyDollar[2].string, yyDollar[4].iAstNode, yyDollar[6].iAstNode}
		}
	case 30:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:78
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].iAstNode, nil}
		}
	case 31:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:80
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].iAstNode, yyDollar[8].iAstNode}
		}
	}
	goto yystack /* stack new state and value */
}
