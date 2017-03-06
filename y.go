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
const EQ = 57361
const NE = 57362
const GE = 57363
const LE = 57364
const OR = 57365
const AND = 57366
const NOT = 57367

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
	"EQ",
	"NE",
	"GE",
	"LE",
	"OR",
	"AND",
	"NOT",
	"'>'",
	"'<'",
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

//line template.y:126

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 53
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 229

var yyAct = [...]int{

	43, 44, 42, 90, 62, 63, 64, 65, 66, 67,
	37, 60, 61, 31, 115, 51, 52, 53, 54, 59,
	55, 110, 109, 51, 52, 53, 112, 111, 47, 48,
	107, 74, 51, 52, 53, 54, 49, 55, 58, 50,
	39, 32, 56, 38, 49, 47, 48, 50, 26, 96,
	12, 69, 70, 49, 103, 88, 50, 73, 104, 77,
	98, 75, 80, 81, 82, 83, 84, 85, 86, 87,
	102, 71, 98, 91, 99, 77, 98, 91, 19, 95,
	62, 63, 64, 65, 66, 67, 18, 60, 61, 35,
	91, 106, 97, 33, 100, 29, 101, 28, 11, 29,
	105, 20, 40, 21, 5, 23, 22, 92, 77, 108,
	7, 27, 77, 16, 79, 114, 77, 62, 63, 64,
	65, 66, 67, 17, 60, 61, 25, 72, 93, 51,
	52, 53, 54, 68, 55, 36, 116, 57, 41, 24,
	14, 13, 47, 48, 51, 52, 53, 54, 78, 55,
	49, 113, 34, 50, 6, 15, 4, 47, 48, 51,
	52, 53, 54, 46, 55, 49, 94, 10, 50, 45,
	3, 2, 47, 48, 1, 8, 30, 9, 0, 0,
	49, 0, 0, 50, 62, 63, 64, 65, 66, 67,
	0, 60, 61, 0, 89, 51, 52, 53, 54, 0,
	55, 0, 76, 0, 0, 0, 0, 0, 47, 48,
	0, 0, 0, 0, 0, 0, 49, 0, 0, 50,
	62, 63, 64, 65, 66, 67, 0, 60, 61,
}
var yyPact = [...]int{

	90, -1000, -1000, 162, -1000, 70, 20, -1000, 137, 136,
	-1000, 108, 162, 58, 50, 72, -1000, -1000, 135, 135,
	18, 106, 68, -1000, 9, 64, -1000, -1000, 73, 135,
	-1000, -1000, -23, 13, 10, 87, -1000, 134, 28, 28,
	133, -1000, 8, -1000, -15, -1000, -1000, -1000, 129, 19,
	19, 43, -1000, -1000, 123, 19, 1, 33, 191, 110,
	19, 19, 19, 19, 19, 19, 19, 19, 27, -1000,
	165, 19, 99, 98, 155, 19, -1000, -1000, -1000, 21,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 19, -1000,
	45, 201, 19, 28, -1000, 41, 25, 29, 19, -1000,
	61, 0, -1000, -1000, -1000, 201, 28, 11, -3, -1000,
	-4, 140, 28, -1000, -16, 125, -1000,
}
var yyPgo = [...]int{

	0, 176, 174, 171, 170, 110, 105, 0, 1, 169,
	163, 156, 155, 154, 3, 106, 2, 152, 148,
}
var yyR1 = [...]int{

	0, 2, 3, 4, 4, 11, 12, 12, 13, 13,
	5, 5, 5, 5, 17, 17, 14, 14, 14, 15,
	15, 6, 6, 1, 1, 16, 16, 7, 7, 7,
	7, 7, 7, 7, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 18, 18,
	9, 10, 10,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 1, 5, 1, 3, 1, 3,
	0, 10, 9, 1, 0, 6, 0, 1, 3, 1,
	3, 0, 2, 1, 3, 1, 3, 0, 1, 1,
	1, 1, 5, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 3, 4, 1, 1, 1, 1, 3,
	8, 6, 10,
}
var yyChk = [...]int{

	-1000, -2, -3, -4, -11, 14, -13, -5, 13, 15,
	5, 28, 30, 4, 4, -12, 5, -5, 28, 28,
	29, 31, -15, -6, 4, -15, 30, 5, 29, 31,
	-1, 4, 32, 29, -17, 16, -6, 33, 30, 30,
	15, 4, -16, -7, -8, -9, -10, 17, 18, 25,
	28, 4, 5, 6, 7, 9, -16, 4, 30, 34,
	26, 27, 19, 20, 21, 22, 23, 24, 4, -8,
	-8, 28, 4, -8, 30, 28, 11, -7, -18, 4,
	-8, -8, -8, -8, -8, -8, -8, -8, 28, 29,
	-14, -8, 8, 30, 11, -14, 28, -14, 31, 29,
	-8, -16, 29, 29, 29, -8, 30, 30, -16, 11,
	10, 30, 30, 11, -16, 30, 11,
}
var yyDef = [...]int{

	3, -2, 1, 10, 4, 0, 2, 8, 0, 0,
	13, 0, 10, 0, 0, 0, 6, 9, 21, 21,
	0, 0, 0, 19, 0, 0, 5, 7, 14, 21,
	22, 23, 0, 0, 0, 0, 20, 0, 27, 27,
	0, 24, 0, 25, 28, 29, 30, 31, 0, 0,
	0, 45, 46, 47, 0, 0, 0, 0, 27, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 42,
	0, 16, 0, 0, 27, 16, 12, 26, 33, 48,
	34, 35, 36, 37, 38, 39, 40, 41, 16, 43,
	0, 17, 0, 27, 11, 0, 0, 0, 0, 44,
	0, 0, 15, 49, 32, 18, 27, 27, 0, 51,
	0, 27, 27, 50, 0, 27, 52,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	28, 29, 3, 3, 31, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 30,
	27, 3, 26, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 32, 3, 33, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 34,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25,
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
		//line template.y:45
		{
			yylex.(*exprLex).result = yyVAL.iAstNode
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:47
		{
			yyVAL.iAstNode = &astFile{yyDollar[1].iAstNode, yyDollar[2].astList}
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:49
		{
			yyVAL.iAstNode = nil
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:50
		{
			yyVAL.iAstNode = &astHeader{yyDollar[1].astList}
		}
	case 5:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:53
		{
			yyVAL.astList = yyDollar[3].astList
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:55
		{
			yyVAL.astList = &astList{[]iAstNode{&astImport{yyDollar[1].string}}}
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:56
		{
			yyVAL.astList.Add(&astImport{yyDollar[3].string})
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:58
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:59
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 10:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:61
		{
			yyVAL.iAstNode = nil
		}
	case 11:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:63
		{
			yyVAL.iAstNode = &astTemplate{yyDollar[2].string, yyDollar[4].astList, yyDollar[6].astUseWrapper, yyDollar[8].astList}
		}
	case 12:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line template.y:65
		{
			yyVAL.iAstNode = &astWrapper{yyDollar[2].string, yyDollar[4].astList, yyDollar[7].astList}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:66
		{
			yyVAL.iAstNode = nil
		}
	case 14:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:68
		{
			yyVAL.astUseWrapper = nil
		}
	case 15:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:70
		{
			yyVAL.astUseWrapper = &astUseWrapper{yyDollar[3].string, yyDollar[5].astList}
		}
	case 16:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:72
		{
			yyVAL.astList = nil
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:73
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:74
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:76
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:77
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 21:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:79
		{
			yyVAL.iAstNode = nil
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:80
		{
			yyVAL.iAstNode = &astVariableDef{yyDollar[1].string, yyDollar[2].string}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:82
		{
			yyVAL.string = yyDollar[1].string
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:83
		{
			yyVAL.string = "[]" + yyDollar[3].string
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:85
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:86
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 27:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:88
		{
			yyVAL.iAstNode = nil
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:89
		{
			yyVAL.iAstNode = &astWriteValue{yyDollar[1].iAstNode}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:90
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:91
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:92
		{
			yyVAL.iAstNode = &astWriteContent{}
		}
	case 32:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:94
		{
			yyVAL.iAstNode = &astProcessTemplate{yyDollar[2].string, yyDollar[4].astList}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:95
		{
			yyDollar[3].astFilter.value = yyDollar[1].iAstNode
			yyVAL.iAstNode = &astWriteString{yyDollar[3].astFilter}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:98
		{
			yyVAL.iAstNode = &astExpr{">", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:99
		{
			yyVAL.iAstNode = &astExpr{"<", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:100
		{
			yyVAL.iAstNode = &astExpr{"==", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:101
		{
			yyVAL.iAstNode = &astExpr{"!=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:102
		{
			yyVAL.iAstNode = &astExpr{">=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:103
		{
			yyVAL.iAstNode = &astExpr{"<=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:104
		{
			yyVAL.iAstNode = &astExpr{"||", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:105
		{
			yyVAL.iAstNode = &astExpr{"&&", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:106
		{
			yyVAL.iAstNode = &astExpr{"!", nil, yyDollar[2].iAstNode}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:107
		{
			yyVAL.iAstNode = &astParenthesis{yyDollar[2].iAstNode}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:109
		{
			yyVAL.iAstNode = &astFunc{yyDollar[1].string, yyDollar[3].astList}
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:110
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:111
		{
			yyVAL.iAstNode = &astString{yyDollar[1].string}
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:112
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:115
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:116
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 50:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line template.y:120
		{
			yyVAL.iAstNode = &astLoop{yyDollar[2].string, yyDollar[4].iAstNode, yyDollar[6].astList}
		}
	case 51:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:122
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, nil}
		}
	case 52:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:124
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, yyDollar[8].astList}
		}
	}
	goto yystack /* stack new state and value */
}
