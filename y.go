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
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'>'",
	"'<'",
	"'('",
	"')'",
	"';'",
	"','",
	"'['",
	"']'",
	"'|'",
	"'.'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line template.y:138

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 63
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 291

var yyAct = [...]int{

	45, 46, 44, 105, 54, 70, 71, 72, 73, 74,
	75, 121, 66, 67, 68, 69, 64, 65, 70, 71,
	72, 73, 74, 75, 63, 66, 67, 68, 69, 64,
	65, 124, 81, 129, 81, 80, 138, 80, 70, 71,
	72, 73, 74, 75, 60, 66, 67, 68, 69, 64,
	65, 120, 104, 77, 78, 31, 127, 38, 118, 125,
	83, 118, 119, 88, 118, 135, 91, 92, 93, 94,
	95, 96, 97, 98, 99, 100, 101, 102, 34, 33,
	29, 106, 28, 20, 29, 21, 88, 32, 106, 112,
	115, 70, 71, 72, 73, 74, 75, 134, 66, 67,
	68, 69, 64, 65, 130, 106, 111, 117, 85, 62,
	41, 40, 122, 26, 123, 12, 53, 55, 56, 57,
	128, 58, 133, 132, 126, 116, 103, 86, 84, 49,
	50, 88, 131, 79, 19, 88, 18, 51, 137, 88,
	59, 11, 23, 113, 52, 53, 55, 56, 57, 36,
	58, 22, 139, 42, 5, 110, 109, 108, 49, 50,
	10, 53, 55, 56, 27, 7, 51, 59, 8, 59,
	9, 25, 37, 52, 53, 55, 56, 57, 17, 58,
	16, 136, 51, 107, 90, 59, 82, 49, 50, 52,
	76, 61, 43, 39, 24, 51, 14, 13, 59, 89,
	35, 6, 52, 70, 71, 72, 73, 74, 75, 15,
	66, 67, 68, 69, 64, 65, 53, 55, 56, 57,
	4, 58, 48, 114, 47, 3, 2, 1, 30, 49,
	50, 0, 0, 0, 0, 0, 0, 51, 0, 0,
	59, 0, 0, 0, 52, 53, 55, 56, 57, 0,
	58, 0, 87, 0, 0, 0, 0, 0, 49, 50,
	0, 0, 53, 55, 56, 57, 51, 58, 0, 59,
	0, 0, 0, 52, 0, 49, 50, 0, 0, 0,
	0, 0, 0, 51, 0, 0, 59, 0, 0, 0,
	52,
}
var yyPact = [...]int{

	140, -1000, -1000, 155, -1000, 109, 81, -1000, 193, 192,
	-1000, 175, 155, 104, 102, 50, -1000, -1000, 190, 190,
	79, 159, 49, -1000, 51, 45, -1000, -1000, 133, 190,
	-1000, -1000, 20, 189, 77, 76, 138, -1000, 188, -1000,
	258, 258, 187, -1000, 75, -1000, -14, -1000, -1000, -1000,
	186, 157, 157, 101, -4, -1000, -1000, 182, 157, 96,
	74, 95, 241, 180, 157, 157, 157, 157, 157, 157,
	157, 157, 157, 157, 157, 157, 94, -1000, 19, 157,
	179, 151, 147, 72, 139, 212, 157, -1000, -1000, -1000,
	93, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 157, -1000, 29, 184, -1000, 14, -26,
	157, 258, -2, -1000, -1000, 26, 91, 23, 157, -1000,
	-1000, -1000, -1, 70, -1000, -1000, -1000, -1000, 184, 258,
	112, 63, -1000, 31, 170, 258, -1000, 2, 141, -1000,
}
var yyPgo = [...]int{

	0, 228, 4, 227, 226, 225, 165, 142, 0, 1,
	224, 222, 220, 209, 201, 3, 151, 2, 200, 199,
}
var yyR1 = [...]int{

	0, 3, 4, 5, 5, 12, 13, 13, 14, 14,
	6, 6, 6, 6, 18, 18, 15, 15, 15, 16,
	16, 7, 7, 1, 1, 1, 17, 17, 8, 8,
	8, 8, 8, 8, 8, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 2, 2, 2, 2, 2, 19, 19,
	10, 11, 11,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 1, 5, 1, 3, 1, 3,
	0, 10, 9, 1, 0, 6, 0, 1, 3, 1,
	3, 0, 2, 1, 3, 2, 1, 3, 0, 1,
	1, 1, 1, 5, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 3, 4,
	1, 1, 1, 3, 4, 4, 4, 1, 1, 3,
	8, 6, 10,
}
var yyChk = [...]int{

	-1000, -3, -4, -5, -12, 14, -14, -6, 13, 15,
	5, 32, 34, 4, 4, -13, 5, -6, 32, 32,
	33, 35, -16, -7, 4, -16, 34, 5, 33, 35,
	-1, 4, 36, 28, 33, -18, 16, -7, 37, 4,
	34, 34, 15, 4, -17, -8, -9, -10, -11, 17,
	18, 25, 32, 4, -2, 5, 6, 7, 9, 28,
	-17, 4, 34, 38, 30, 31, 26, 27, 28, 29,
	19, 20, 21, 22, 23, 24, 4, -9, -9, 32,
	39, 36, 4, -9, 32, 34, 32, 11, -8, -19,
	4, -9, -9, -9, -9, -9, -9, -9, -9, -9,
	-9, -9, -9, 32, 33, -15, -9, 4, 6, 5,
	8, 34, -2, 4, 11, -15, 32, -15, 35, 33,
	37, 37, -9, -17, 33, 33, 33, 33, -9, 34,
	34, -17, 11, 10, 34, 34, 11, -17, 34, 11,
}
var yyDef = [...]int{

	3, -2, 1, 10, 4, 0, 2, 8, 0, 0,
	13, 0, 10, 0, 0, 0, 6, 9, 21, 21,
	0, 0, 0, 19, 0, 0, 5, 7, 14, 21,
	22, 23, 0, 0, 0, 0, 0, 20, 0, 25,
	28, 28, 0, 24, 0, 26, 29, 30, 31, 32,
	0, 0, 0, 57, 50, 51, 52, 0, 0, 0,
	0, 0, 28, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 47, 0, 16,
	0, 0, 0, 0, 0, 28, 16, 12, 27, 34,
	58, 35, 36, 37, 38, 39, 40, 41, 42, 43,
	44, 45, 46, 16, 48, 0, 17, 53, 0, 0,
	0, 28, 0, 57, 11, 0, 0, 0, 0, 49,
	54, 55, 0, 0, 56, 15, 59, 33, 18, 28,
	28, 0, 61, 0, 28, 28, 60, 0, 28, 62,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	32, 33, 28, 26, 35, 27, 39, 29, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 34,
	31, 3, 30, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 36, 3, 37, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 38,
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
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:84
		{
			yyVAL.string = "*" + yyDollar[2].string
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:86
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:87
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:89
		{
			yyVAL.iAstNode = nil
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:90
		{
			yyVAL.iAstNode = &astWriteValue{yyDollar[1].iAstNode}
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
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:93
		{
			yyVAL.iAstNode = &astWriteContent{}
		}
	case 33:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:95
		{
			yyVAL.iAstNode = &astProcessTemplate{yyDollar[2].string, yyDollar[4].astList}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:96
		{
			yyDollar[3].astFilter.value = yyDollar[1].iAstNode
			yyVAL.iAstNode = &astWriteString{yyDollar[3].astFilter}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:99
		{
			yyVAL.iAstNode = &astExpr{">", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:100
		{
			yyVAL.iAstNode = &astExpr{"<", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:101
		{
			yyVAL.iAstNode = &astExpr{"+", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:102
		{
			yyVAL.iAstNode = &astExpr{"-", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:103
		{
			yyVAL.iAstNode = &astExpr{"*", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:104
		{
			yyVAL.iAstNode = &astExpr{"/", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:105
		{
			yyVAL.iAstNode = &astExpr{"==", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:106
		{
			yyVAL.iAstNode = &astExpr{"!=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:107
		{
			yyVAL.iAstNode = &astExpr{">=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:108
		{
			yyVAL.iAstNode = &astExpr{"<=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:109
		{
			yyVAL.iAstNode = &astExpr{"||", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:110
		{
			yyVAL.iAstNode = &astExpr{"&&", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:111
		{
			yyVAL.iAstNode = &astExpr{"!", nil, yyDollar[2].iAstNode}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:112
		{
			yyVAL.iAstNode = &astParenthesis{yyDollar[2].iAstNode}
		}
	case 49:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:114
		{
			yyVAL.iAstNode = &astFunc{yyDollar[1].string, yyDollar[3].astList}
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:115
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:116
		{
			yyVAL.iAstNode = &astString{yyDollar[1].string}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:117
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:119
		{
			yyVAL.string = yyDollar[1].string + "." + yyDollar[3].string
		}
	case 54:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:120
		{
			yyVAL.string = yyDollar[1].string + "[" + yyDollar[3].string + "]"
		}
	case 55:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:121
		{
			yyVAL.string = yyDollar[1].string + "[" + yyDollar[3].string + "]"
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:122
		{
			yyVAL.string = "*(" + yyDollar[3].string + ")"
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:124
		{
			yyVAL.string = yyDollar[1].string
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:127
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:128
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 60:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line template.y:132
		{
			yyVAL.iAstNode = &astLoop{yyDollar[2].string, yyDollar[4].iAstNode, yyDollar[6].astList}
		}
	case 61:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:134
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, nil}
		}
	case 62:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:136
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, yyDollar[8].astList}
		}
	}
	goto yystack /* stack new state and value */
}
