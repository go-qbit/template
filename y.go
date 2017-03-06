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
const PROCESS = 57360

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
	"PROCESS",
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

//line template.y:110

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 43
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 130

var yyAct = [...]int{

	43, 74, 42, 73, 66, 23, 45, 44, 56, 50,
	55, 51, 88, 87, 37, 31, 82, 81, 80, 48,
	49, 93, 90, 45, 44, 79, 50, 80, 51, 33,
	94, 29, 75, 89, 32, 36, 48, 49, 45, 44,
	69, 50, 52, 51, 28, 91, 29, 20, 35, 21,
	84, 48, 49, 59, 83, 65, 71, 45, 44, 62,
	50, 68, 51, 65, 72, 54, 39, 38, 26, 12,
	48, 49, 77, 76, 78, 45, 44, 22, 50, 63,
	51, 19, 85, 18, 11, 65, 86, 40, 48, 49,
	65, 10, 5, 92, 65, 45, 44, 25, 50, 8,
	51, 9, 64, 7, 70, 60, 61, 27, 48, 49,
	16, 67, 58, 57, 53, 41, 17, 24, 14, 13,
	34, 6, 15, 4, 47, 46, 3, 2, 1, 30,
}
var yyPact = [...]int{

	78, -1000, -1000, 86, -1000, 65, 48, -1000, 115, 114,
	-1000, 105, 86, 64, 62, 27, -1000, -1000, 113, 113,
	47, 102, 24, -1000, 11, 9, -1000, -1000, 32, 113,
	-1000, -1000, -10, 46, 45, 72, -1000, 111, 71, 71,
	110, -1000, 44, -1000, -15, -17, -1000, -1000, -1000, 109,
	108, 101, 38, 60, 91, 107, 107, 21, 96, 35,
	-1000, -1000, 53, 101, -1000, -1000, -1000, 13, -1000, 101,
	101, 71, -1000, 5, -1000, -3, -4, 33, 29, -1000,
	101, -1000, -1000, 71, 2, -1000, 12, -1000, 1, 34,
	71, -1000, 0, 19, -1000,
}
var yyPgo = [...]int{

	0, 129, 128, 127, 126, 103, 5, 0, 1, 125,
	124, 123, 122, 121, 3, 77, 2, 120, 4,
}
var yyR1 = [...]int{

	0, 2, 3, 4, 4, 11, 12, 12, 13, 13,
	5, 5, 5, 5, 17, 17, 14, 14, 14, 15,
	15, 6, 6, 1, 1, 16, 16, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 8, 8, 18, 18,
	9, 10, 10,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 1, 5, 1, 3, 1, 3,
	0, 10, 9, 1, 0, 6, 0, 1, 3, 1,
	3, 0, 2, 1, 3, 1, 3, 0, 1, 1,
	3, 3, 1, 1, 1, 5, 1, 1, 1, 3,
	8, 6, 10,
}
var yyChk = [...]int{

	-1000, -2, -3, -4, -11, 14, -13, -5, 13, 15,
	5, 19, 21, 4, 4, -12, 5, -5, 19, 19,
	20, 22, -15, -6, 4, -15, 21, 5, 20, 22,
	-1, 4, 23, 20, -17, 16, -6, 24, 21, 21,
	15, 4, -16, -7, 5, 4, -9, -10, 17, 18,
	7, 9, -16, 4, 21, 25, 25, 4, 4, -8,
	4, 5, 21, 19, 11, -7, -18, 4, -18, 19,
	8, 21, 11, -14, -8, 19, -14, -8, -16, 20,
	22, 20, 20, 21, 21, -8, -16, 11, 10, 21,
	21, 11, -16, 21, 11,
}
var yyDef = [...]int{

	3, -2, 1, 10, 4, 0, 2, 8, 0, 0,
	13, 0, 10, 0, 0, 0, 6, 9, 21, 21,
	0, 0, 0, 19, 0, 0, 5, 7, 14, 21,
	22, 23, 0, 0, 0, 0, 20, 0, 27, 27,
	0, 24, 0, 25, 28, 29, 32, 33, 34, 0,
	0, 0, 0, 0, 27, 0, 0, 0, 0, 0,
	36, 37, 27, 16, 12, 26, 31, 38, 30, 16,
	0, 27, 11, 0, 17, 0, 0, 0, 0, 15,
	0, 39, 35, 27, 27, 18, 0, 41, 0, 27,
	27, 40, 0, 27, 42,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	19, 20, 3, 3, 22, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 21,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 23, 3, 24, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 25,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18,
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
		//line template.y:41
		{
			yylex.(*exprLex).result = yyVAL.iAstNode
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:43
		{
			yyVAL.iAstNode = &astFile{yyDollar[1].iAstNode, yyDollar[2].astList}
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:45
		{
			yyVAL.iAstNode = nil
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:46
		{
			yyVAL.iAstNode = &astHeader{yyDollar[1].astList}
		}
	case 5:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:49
		{
			yyVAL.astList = yyDollar[3].astList
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:51
		{
			yyVAL.astList = &astList{[]iAstNode{&astImport{yyDollar[1].string}}}
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:52
		{
			yyVAL.astList.Add(&astImport{yyDollar[3].string})
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:54
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:55
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 10:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:57
		{
			yyVAL.iAstNode = nil
		}
	case 11:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:59
		{
			yyVAL.iAstNode = &astTemplate{yyDollar[2].string, yyDollar[4].astList, yyDollar[6].astUseWrapper, yyDollar[8].astList}
		}
	case 12:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line template.y:61
		{
			yyVAL.iAstNode = &astWrapper{yyDollar[2].string, yyDollar[4].astList, yyDollar[7].astList}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:62
		{
			yyVAL.iAstNode = nil
		}
	case 14:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:64
		{
			yyVAL.astUseWrapper = nil
		}
	case 15:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:66
		{
			yyVAL.astUseWrapper = &astUseWrapper{yyDollar[3].string, yyDollar[5].astList}
		}
	case 16:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:68
		{
			yyVAL.astList = nil
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:69
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:70
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:72
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:73
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 21:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:75
		{
			yyVAL.iAstNode = nil
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:76
		{
			yyVAL.iAstNode = &astVariableDef{yyDollar[1].string, yyDollar[2].string}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:78
		{
			yyVAL.string = yyDollar[1].string
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:79
		{
			yyVAL.string = "[]" + yyDollar[3].string
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:81
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:82
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 27:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:84
		{
			yyVAL.iAstNode = nil
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:85
		{
			yyVAL.iAstNode = &astWriteString{&astString{yyDollar[1].string}}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:86
		{
			yyVAL.iAstNode = &astWriteValue{&astValue{yyDollar[1].string}}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:87
		{
			yyDollar[3].astFilter.value = &astValue{yyDollar[1].string}
			yyVAL.iAstNode = &astWriteString{yyDollar[3].astFilter}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:88
		{
			yyDollar[3].astFilter.value = &astString{yyDollar[1].string}
			yyVAL.iAstNode = &astWriteString{yyDollar[3].astFilter}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:89
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:90
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:91
		{
			yyVAL.iAstNode = &astWriteContent{}
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:93
		{
			yyVAL.iAstNode = &astProcessTemplate{yyDollar[2].string, yyDollar[4].astList}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:96
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:97
		{
			yyVAL.iAstNode = &astString{yyDollar[1].string}
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:99
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:100
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 40:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line template.y:104
		{
			yyVAL.iAstNode = &astLoop{yyDollar[2].string, yyDollar[4].iAstNode, yyDollar[6].astList}
		}
	case 41:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:106
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, nil}
		}
	case 42:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:108
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, yyDollar[8].astList}
		}
	}
	goto yystack /* stack new state and value */
}
