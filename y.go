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
	"'%'",
	"'>'",
	"'<'",
	"'('",
	"')'",
	"';'",
	"','",
	"'.'",
	"'['",
	"']'",
	"'|'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line template.y:146

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 68
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 367

var yyAct = [...]int{

	49, 112, 47, 134, 57, 130, 85, 86, 46, 75,
	76, 77, 78, 79, 80, 48, 70, 71, 72, 73,
	74, 68, 69, 31, 75, 76, 77, 78, 79, 80,
	67, 70, 71, 72, 73, 74, 68, 69, 129, 39,
	150, 85, 86, 38, 91, 63, 137, 33, 92, 148,
	138, 127, 127, 156, 136, 82, 83, 32, 135, 128,
	127, 127, 88, 117, 34, 28, 29, 29, 155, 97,
	98, 99, 100, 101, 102, 103, 104, 105, 106, 107,
	108, 109, 94, 151, 149, 113, 56, 58, 59, 142,
	125, 118, 113, 123, 120, 56, 58, 59, 60, 90,
	61, 147, 146, 20, 110, 21, 94, 54, 52, 53,
	62, 113, 126, 66, 42, 55, 54, 89, 131, 62,
	41, 26, 133, 12, 55, 84, 19, 18, 139, 11,
	22, 56, 58, 59, 60, 121, 61, 113, 143, 23,
	36, 10, 145, 144, 52, 53, 141, 43, 5, 8,
	25, 9, 54, 153, 154, 62, 27, 7, 94, 62,
	55, 116, 115, 16, 132, 94, 124, 114, 96, 37,
	17, 94, 94, 75, 76, 77, 78, 79, 80, 87,
	70, 71, 72, 73, 74, 68, 69, 81, 65, 140,
	75, 76, 77, 78, 79, 80, 64, 70, 71, 72,
	73, 74, 68, 69, 45, 44, 119, 75, 76, 77,
	78, 79, 80, 40, 70, 71, 72, 73, 74, 68,
	69, 24, 111, 75, 76, 77, 78, 79, 80, 14,
	70, 71, 72, 73, 74, 68, 69, 56, 58, 59,
	60, 13, 61, 95, 158, 35, 6, 15, 4, 51,
	52, 53, 50, 3, 2, 1, 30, 0, 54, 0,
	0, 62, 56, 58, 59, 60, 55, 61, 0, 157,
	0, 0, 0, 0, 0, 52, 53, 0, 0, 0,
	0, 0, 0, 54, 0, 0, 62, 56, 58, 59,
	60, 55, 61, 0, 152, 0, 0, 0, 0, 0,
	52, 53, 0, 0, 0, 0, 0, 0, 54, 0,
	0, 62, 56, 58, 59, 60, 55, 61, 0, 122,
	0, 0, 0, 0, 0, 52, 53, 0, 0, 0,
	0, 0, 0, 54, 0, 0, 62, 56, 58, 59,
	60, 55, 61, 0, 93, 0, 0, 0, 0, 0,
	52, 53, 0, 0, 0, 0, 0, 0, 54, 0,
	0, 62, 0, 0, 0, 0, 55,
}
var yyPact = [...]int{

	134, -1000, -1000, 136, -1000, 96, 88, -1000, 237, 225,
	-1000, 158, 136, 94, 93, 69, -1000, -1000, 217, 217,
	86, 151, 31, -1000, 19, 30, -1000, -1000, 124, 217,
	-1000, 6, 0, 209, 85, 79, 132, -1000, 201, 200,
	-29, 127, 127, 192, -1000, -1000, 184, 78, -1000, -10,
	-1000, -1000, -1000, 183, 82, 82, 92, 4, -1000, -1000,
	175, 82, 84, 64, 11, -1000, 333, 164, 82, 82,
	82, 82, 82, 82, 82, 82, 82, 82, 82, 82,
	82, 71, -1000, 188, 82, 163, 156, 55, 171, 131,
	308, 82, 162, -1000, -1000, -1000, 57, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	82, -1000, 25, 204, -1000, -1, -34, 82, 160, 127,
	-31, -1000, -1000, 24, 21, 12, 16, 82, -1000, -1000,
	-1000, 154, 138, 54, -1000, -1000, 82, -1000, -1000, 204,
	127, 82, 91, 15, 49, 5, -1000, 48, -1000, 283,
	127, 127, -1000, 33, 18, 258, 233, -1000, -1000,
}
var yyPgo = [...]int{

	0, 256, 4, 255, 254, 253, 157, 139, 15, 0,
	252, 249, 248, 247, 246, 1, 130, 2, 245, 243,
}
var yyR1 = [...]int{

	0, 3, 4, 5, 5, 12, 13, 13, 14, 14,
	6, 6, 6, 6, 18, 18, 18, 15, 15, 15,
	16, 16, 7, 7, 1, 1, 1, 1, 1, 17,
	17, 8, 8, 8, 8, 8, 8, 8, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 2, 2, 2,
	2, 2, 19, 19, 10, 10, 11, 11,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 1, 5, 1, 3, 1, 3,
	0, 10, 9, 1, 0, 6, 8, 0, 1, 3,
	1, 3, 0, 2, 1, 3, 2, 3, 4, 1,
	3, 0, 1, 1, 1, 1, 5, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 3, 4, 1, 1, 1, 3, 4, 4,
	4, 1, 1, 3, 8, 10, 6, 10,
}
var yyChk = [...]int{

	-1000, -3, -4, -5, -12, 14, -14, -6, 13, 15,
	5, 33, 35, 4, 4, -13, 5, -6, 33, 33,
	34, 36, -16, -7, 4, -16, 35, 5, 34, 36,
	-1, 4, 38, 28, 34, -18, 16, -7, 37, 39,
	4, 35, 35, 15, 4, 4, 37, -17, -8, -9,
	-10, -11, 17, 18, 25, 33, 4, -2, 5, 6,
	7, 9, 28, -17, 4, 4, 35, 40, 31, 32,
	26, 27, 28, 29, 30, 19, 20, 21, 22, 23,
	24, 4, -9, -9, 33, 37, 38, 4, -9, 33,
	35, 33, 37, 11, -8, -19, 4, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -9, -9, -9,
	33, 34, -15, -9, 4, 6, 5, 8, 36, 35,
	-2, 4, 11, -15, 4, 33, -15, 36, 34, 39,
	39, -9, 4, -17, 34, 34, 33, 34, 34, -9,
	35, 8, 35, -15, -17, -9, 11, 10, 34, 35,
	35, 35, 11, -17, -17, 35, 35, 11, 11,
}
var yyDef = [...]int{

	3, -2, 1, 10, 4, 0, 2, 8, 0, 0,
	13, 0, 10, 0, 0, 0, 6, 9, 22, 22,
	0, 0, 0, 20, 0, 0, 5, 7, 14, 22,
	23, 24, 0, 0, 0, 0, 0, 21, 0, 0,
	26, 31, 31, 0, 27, 25, 0, 0, 29, 32,
	33, 34, 35, 0, 0, 0, 61, 54, 55, 56,
	0, 0, 0, 0, 0, 28, 31, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 51, 0, 17, 0, 0, 0, 0, 0,
	31, 17, 0, 12, 30, 37, 62, 38, 39, 40,
	41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
	17, 52, 0, 18, 57, 0, 0, 0, 0, 31,
	0, 61, 11, 0, 0, 0, 0, 0, 53, 58,
	59, 0, 0, 0, 60, 15, 17, 63, 36, 19,
	31, 0, 31, 0, 0, 0, 66, 0, 16, 31,
	31, 31, 64, 0, 0, 31, 31, 65, 67,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 30, 3, 3,
	33, 34, 28, 26, 36, 27, 37, 29, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 35,
	32, 3, 31, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 38, 3, 39, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 40,
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
			yyVAL.astUseWrapper = &astUseWrapper{"", yyDollar[3].string, yyDollar[5].astList}
		}
	case 16:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line template.y:72
		{
			yyVAL.astUseWrapper = &astUseWrapper{yyDollar[3].string, yyDollar[5].string, yyDollar[7].astList}
		}
	case 17:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:74
		{
			yyVAL.astList = nil
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:75
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:76
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:78
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:79
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 22:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:81
		{
			yyVAL.iAstNode = nil
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:82
		{
			yyVAL.iAstNode = &astVariableDef{yyDollar[1].string, yyDollar[2].string}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:84
		{
			yyVAL.string = yyDollar[1].string
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:85
		{
			yyVAL.string = "[]" + yyDollar[3].string
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:86
		{
			yyVAL.string = "*" + yyDollar[2].string
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:87
		{
			yyVAL.string = yyDollar[1].string + "." + yyDollar[3].string
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:89
		{
			yyVAL.string = "*" + yyDollar[2].string + "." + yyDollar[4].string
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:91
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:92
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 31:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:94
		{
			yyVAL.iAstNode = nil
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:95
		{
			yyVAL.iAstNode = &astWriteValue{yyDollar[1].iAstNode}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:96
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:97
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:98
		{
			yyVAL.iAstNode = &astWriteContent{}
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:100
		{
			yyVAL.iAstNode = &astProcessTemplate{yyDollar[2].string, yyDollar[4].astList}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:101
		{
			yyDollar[3].astFilter.value = yyDollar[1].iAstNode
			yyVAL.iAstNode = &astWriteString{yyDollar[3].astFilter}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:104
		{
			yyVAL.iAstNode = &astExpr{">", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:105
		{
			yyVAL.iAstNode = &astExpr{"<", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:106
		{
			yyVAL.iAstNode = &astExpr{"+", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:107
		{
			yyVAL.iAstNode = &astExpr{"-", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:108
		{
			yyVAL.iAstNode = &astExpr{"*", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:109
		{
			yyVAL.iAstNode = &astExpr{"/", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:110
		{
			yyVAL.iAstNode = &astExpr{"%", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:111
		{
			yyVAL.iAstNode = &astExpr{"==", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:112
		{
			yyVAL.iAstNode = &astExpr{"!=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:113
		{
			yyVAL.iAstNode = &astExpr{">=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:114
		{
			yyVAL.iAstNode = &astExpr{"<=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:115
		{
			yyVAL.iAstNode = &astExpr{"||", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:116
		{
			yyVAL.iAstNode = &astExpr{"&&", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:117
		{
			yyVAL.iAstNode = &astExpr{"!", nil, yyDollar[2].iAstNode}
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:118
		{
			yyVAL.iAstNode = &astParenthesis{yyDollar[2].iAstNode}
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:120
		{
			yyVAL.iAstNode = &astFunc{yyDollar[1].string, yyDollar[3].astList}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:121
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:122
		{
			yyVAL.iAstNode = &astString{yyDollar[1].string}
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:123
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:125
		{
			yyVAL.string = yyDollar[1].string + "." + yyDollar[3].string
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:126
		{
			yyVAL.string = yyDollar[1].string + "[" + yyDollar[3].string + "]"
		}
	case 59:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:127
		{
			yyVAL.string = yyDollar[1].string + "[" + yyDollar[3].string + "]"
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:128
		{
			yyVAL.string = "*(" + yyDollar[3].string + ")"
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:130
		{
			yyVAL.string = yyDollar[1].string
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:133
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:134
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 64:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line template.y:138
		{
			yyVAL.iAstNode = &astLoop{"_", yyDollar[2].string, yyDollar[4].iAstNode, yyDollar[6].astList}
		}
	case 65:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:140
		{
			yyVAL.iAstNode = &astLoop{yyDollar[2].string, yyDollar[4].string, yyDollar[6].iAstNode, yyDollar[8].astList}
		}
	case 66:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:142
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, nil}
		}
	case 67:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:144
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, yyDollar[8].astList}
		}
	}
	goto yystack /* stack new state and value */
}
