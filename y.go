//line template.y:2
package template

import (
	__yyfmt__ "fmt"
)

//line template.y:9
type yySymType struct {
	yys           int
	string        string
	iAstNode      iAstNode
	astList       *astList
	astUseWrapper *astUseWrapper
	astFilter     *astFilter
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
const TEMPLATE = 57355
const IMPORT = 57356
const WRAPPER = 57357
const USE = 57358
const CONTENT_MARKER = 57359

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
	"TEMPLATE",
	"IMPORT",
	"WRAPPER",
	"USE",
	"CONTENT_MARKER",
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

//line template.y:107

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 42
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 119

var yyAct = [...]int{

	43, 57, 42, 64, 23, 55, 54, 45, 44, 37,
	49, 31, 50, 83, 82, 75, 77, 76, 45, 44,
	48, 49, 88, 50, 33, 89, 29, 35, 28, 32,
	29, 48, 45, 44, 36, 49, 20, 50, 21, 86,
	85, 84, 51, 45, 44, 48, 49, 79, 50, 78,
	69, 68, 60, 53, 63, 39, 48, 45, 44, 66,
	49, 63, 50, 71, 62, 38, 26, 12, 72, 73,
	48, 74, 22, 61, 45, 44, 19, 49, 80, 50,
	63, 81, 18, 11, 40, 63, 10, 48, 87, 63,
	5, 7, 25, 67, 8, 27, 9, 58, 59, 16,
	65, 56, 52, 41, 17, 24, 14, 13, 34, 70,
	6, 15, 4, 47, 46, 3, 2, 1, 30,
}
var yyPact = [...]int{

	76, -1000, -1000, 81, -1000, 65, 47, -1000, 103, 102,
	-1000, 94, 81, 64, 58, 17, -1000, -1000, 101, 101,
	46, 90, 9, -1000, 7, 5, -1000, -1000, 11, 101,
	-1000, -1000, -14, 45, 35, 69, -1000, 99, 70, 70,
	98, -1000, 33, -1000, -18, -19, -1000, -1000, -1000, 97,
	93, 32, 55, 53, 96, 96, 85, 31, -1000, -1000,
	39, 93, -1000, -1000, -1000, 50, -1000, 93, 70, -1000,
	-4, -1000, -3, 29, 27, -1000, 93, -1000, 70, 3,
	-1000, 21, -1000, 20, 28, 70, -1000, 2, 14, -1000,
}
var yyPgo = [...]int{

	0, 118, 117, 116, 115, 91, 4, 0, 1, 114,
	113, 112, 111, 110, 109, 72, 2, 108, 3,
}
var yyR1 = [...]int{

	0, 2, 3, 4, 4, 11, 12, 12, 13, 13,
	5, 5, 5, 5, 17, 17, 14, 14, 14, 15,
	15, 6, 6, 1, 1, 16, 16, 7, 7, 7,
	7, 7, 7, 7, 7, 8, 8, 18, 18, 9,
	10, 10,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 1, 5, 1, 3, 1, 3,
	0, 10, 9, 1, 0, 6, 0, 1, 3, 1,
	3, 0, 2, 1, 3, 1, 3, 0, 1, 1,
	3, 3, 1, 1, 1, 1, 1, 1, 3, 8,
	6, 10,
}
var yyChk = [...]int{

	-1000, -2, -3, -4, -11, 14, -13, -5, 13, 15,
	5, 18, 20, 4, 4, -12, 5, -5, 18, 18,
	19, 21, -15, -6, 4, -15, 20, 5, 19, 21,
	-1, 4, 22, 19, -17, 16, -6, 23, 20, 20,
	15, 4, -16, -7, 5, 4, -9, -10, 17, 7,
	9, -16, 4, 20, 24, 24, 4, -8, 4, 5,
	20, 18, 11, -7, -18, 4, -18, 8, 20, 11,
	-14, -8, 18, -8, -16, 19, 21, 19, 20, 20,
	-8, -16, 11, 10, 20, 20, 11, -16, 20, 11,
}
var yyDef = [...]int{

	3, -2, 1, 10, 4, 0, 2, 8, 0, 0,
	13, 0, 10, 0, 0, 0, 6, 9, 21, 21,
	0, 0, 0, 19, 0, 0, 5, 7, 14, 21,
	22, 23, 0, 0, 0, 0, 20, 0, 27, 27,
	0, 24, 0, 25, 28, 29, 32, 33, 34, 0,
	0, 0, 0, 27, 0, 0, 0, 0, 35, 36,
	27, 16, 12, 26, 31, 37, 30, 0, 27, 11,
	0, 17, 0, 0, 0, 15, 0, 38, 27, 27,
	18, 0, 40, 0, 27, 27, 39, 0, 27, 41,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	18, 19, 3, 3, 21, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 20,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 22, 3, 23, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 24,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17,
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
		//line template.y:40
		{
			yylex.(*exprLex).result = yyVAL.iAstNode
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:42
		{
			yyVAL.iAstNode = &astFile{yyDollar[1].iAstNode, yyDollar[2].astList}
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:44
		{
			yyVAL.iAstNode = nil
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:45
		{
			yyVAL.iAstNode = &astHeader{yyDollar[1].astList}
		}
	case 5:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:48
		{
			yyVAL.astList = yyDollar[3].astList
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:50
		{
			yyVAL.astList = &astList{[]iAstNode{&astImport{yyDollar[1].string}}}
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:51
		{
			yyVAL.astList.Add(&astImport{yyDollar[3].string})
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:53
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:54
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 10:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:56
		{
			yyVAL.iAstNode = nil
		}
	case 11:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:58
		{
			yyVAL.iAstNode = &astTemplate{yyDollar[2].string, yyDollar[4].astList, yyDollar[6].astUseWrapper, yyDollar[8].astList}
		}
	case 12:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line template.y:60
		{
			yyVAL.iAstNode = &astWrapper{yyDollar[2].string, yyDollar[4].astList, yyDollar[7].astList}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:61
		{
			yyVAL.iAstNode = nil
		}
	case 14:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:63
		{
			yyVAL.astUseWrapper = nil
		}
	case 15:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:65
		{
			yyVAL.astUseWrapper = &astUseWrapper{yyDollar[3].string, yyDollar[5].astList}
		}
	case 16:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:67
		{
			yyVAL.astList = nil
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:68
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:69
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:71
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:72
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 21:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:74
		{
			yyVAL.iAstNode = nil
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:75
		{
			yyVAL.iAstNode = &astVariableDef{yyDollar[1].string, yyDollar[2].string}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:77
		{
			yyVAL.string = yyDollar[1].string
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:78
		{
			yyVAL.string = "[]" + yyDollar[3].string
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:80
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:81
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 27:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:83
		{
			yyVAL.iAstNode = nil
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:84
		{
			yyVAL.iAstNode = &astWriteString{&astString{yyDollar[1].string}}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:85
		{
			yyVAL.iAstNode = &astWriteValue{&astValue{yyDollar[1].string}}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:86
		{
			yyDollar[3].astFilter.value = &astValue{yyDollar[1].string}
			yyVAL.iAstNode = &astWriteString{yyDollar[3].astFilter}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:87
		{
			yyDollar[3].astFilter.value = &astString{yyDollar[1].string}
			yyVAL.iAstNode = &astWriteString{yyDollar[3].astFilter}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:88
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:89
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:90
		{
			yyVAL.iAstNode = &astWriteContent{}
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:93
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:94
		{
			yyVAL.iAstNode = &astString{yyDollar[1].string}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:96
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:97
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 39:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line template.y:101
		{
			yyVAL.iAstNode = &astLoop{yyDollar[2].string, yyDollar[4].iAstNode, yyDollar[6].astList}
		}
	case 40:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:103
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, nil}
		}
	case 41:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:105
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, yyDollar[8].astList}
		}
	}
	goto yystack /* stack new state and value */
}
