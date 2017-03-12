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
const ASSIGNMENT = 57368

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
	"ASSIGNMENT",
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
	"'='",
	"'|'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line template.y:148

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 70
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 379

var yyAct = [...]int{

	88, 115, 139, 49, 82, 83, 84, 47, 83, 84,
	161, 81, 134, 75, 76, 77, 78, 79, 80, 133,
	48, 70, 71, 72, 73, 74, 68, 69, 39, 31,
	123, 46, 53, 58, 59, 60, 67, 61, 152, 151,
	38, 160, 52, 52, 82, 54, 55, 95, 83, 84,
	63, 96, 156, 56, 33, 154, 153, 62, 131, 124,
	87, 90, 57, 147, 32, 92, 144, 52, 131, 140,
	142, 131, 101, 102, 103, 104, 105, 106, 107, 108,
	109, 110, 111, 112, 113, 114, 116, 98, 132, 120,
	131, 34, 141, 29, 126, 52, 28, 128, 29, 116,
	75, 76, 77, 78, 79, 80, 94, 66, 70, 71,
	72, 73, 74, 68, 69, 98, 42, 155, 41, 89,
	58, 59, 20, 135, 21, 116, 52, 136, 26, 12,
	130, 121, 93, 138, 19, 143, 18, 11, 85, 23,
	56, 89, 36, 148, 62, 116, 52, 22, 52, 57,
	150, 43, 5, 149, 146, 52, 52, 52, 7, 119,
	118, 52, 52, 158, 159, 27, 62, 25, 98, 37,
	10, 17, 16, 137, 129, 98, 117, 100, 8, 91,
	9, 98, 98, 75, 76, 77, 78, 79, 80, 86,
	65, 70, 71, 72, 73, 74, 68, 69, 64, 45,
	145, 75, 76, 77, 78, 79, 80, 44, 40, 70,
	71, 72, 73, 74, 68, 69, 24, 14, 125, 75,
	76, 77, 78, 79, 80, 13, 99, 70, 71, 72,
	73, 74, 68, 69, 35, 122, 75, 76, 77, 78,
	79, 80, 6, 15, 70, 71, 72, 73, 74, 68,
	69, 53, 58, 59, 60, 4, 61, 51, 163, 50,
	3, 2, 1, 30, 54, 55, 53, 58, 59, 60,
	0, 61, 56, 162, 0, 0, 62, 0, 0, 54,
	55, 57, 0, 0, 0, 0, 0, 56, 0, 0,
	0, 62, 53, 58, 59, 60, 57, 61, 0, 157,
	0, 0, 0, 0, 0, 54, 55, 53, 58, 59,
	60, 0, 61, 56, 127, 0, 0, 62, 0, 0,
	54, 55, 57, 0, 0, 0, 0, 0, 56, 0,
	0, 0, 62, 53, 58, 59, 60, 57, 61, 0,
	97, 0, 0, 0, 0, 0, 54, 55, 53, 58,
	59, 60, 0, 61, 56, 0, 0, 0, 62, 0,
	0, 54, 55, 57, 0, 0, 0, 0, 0, 56,
	0, 0, 0, 62, 0, 0, 0, 0, 57,
}
var yyPact = [...]int{

	138, -1000, -1000, 165, -1000, 103, 93, -1000, 221, 213,
	-1000, 167, 165, 102, 100, 87, -1000, -1000, 212, 212,
	92, 160, 61, -1000, 25, 56, -1000, -1000, 126, 212,
	-1000, 2, -12, 204, 82, 80, 136, -1000, 203, 195,
	-7, 344, 344, 194, -1000, -1000, 186, 71, -1000, -6,
	-1000, -1000, -30, 112, -1000, 185, 115, 115, -1000, -1000,
	175, 115, 98, 70, 13, -1000, 329, 173, 115, 115,
	115, 115, 115, 115, 115, 115, 115, 115, 115, 115,
	115, 115, 115, 172, 154, 115, 97, -1000, 10, -1000,
	200, 22, 182, 137, 303, 115, 170, -1000, -1000, -1000,
	96, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 217, 53, 217, -1000, -21, -28,
	217, 115, -1000, 115, 169, 344, -33, -1000, 34, 58,
	35, 115, -1000, -1000, -1000, 31, 164, 146, 27, -1000,
	-1000, 115, -1000, 217, -1000, 344, 115, 28, 21, 19,
	81, -1000, 16, -1000, 288, 344, 344, -1000, 5, -26,
	262, 247, -1000, -1000,
}
var yyPgo = [...]int{

	0, 263, 0, 262, 261, 260, 158, 139, 20, 3,
	259, 257, 255, 243, 242, 1, 147, 7, 234, 226,
}
var yyR1 = [...]int{

	0, 3, 4, 5, 5, 12, 13, 13, 14, 14,
	6, 6, 6, 6, 18, 18, 18, 15, 15, 15,
	16, 16, 7, 7, 1, 1, 1, 1, 1, 17,
	17, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 2,
	2, 2, 2, 2, 19, 19, 10, 10, 11, 11,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 1, 5, 1, 3, 1, 3,
	0, 10, 9, 1, 0, 6, 8, 0, 1, 3,
	1, 3, 0, 2, 1, 3, 2, 3, 4, 1,
	3, 0, 1, 1, 1, 3, 3, 1, 5, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 3, 4, 1, 1, 1, 3,
	4, 4, 4, 1, 1, 3, 8, 10, 6, 10,
}
var yyChk = [...]int{

	-1000, -3, -4, -5, -12, 14, -14, -6, 13, 15,
	5, 34, 36, 4, 4, -13, 5, -6, 34, 34,
	35, 37, -16, -7, 4, -16, 36, 5, 35, 37,
	-1, 4, 39, 29, 35, -18, 16, -7, 38, 40,
	4, 36, 36, 15, 4, 4, 38, -17, -8, -9,
	-10, -11, -2, 4, 17, 18, 25, 34, 5, 6,
	7, 9, 29, -17, 4, 4, 36, 42, 32, 33,
	27, 28, 29, 30, 31, 19, 20, 21, 22, 23,
	24, 41, 34, 38, 39, 26, 4, -9, -2, 4,
	-9, 4, -9, 34, 36, 34, 38, 11, -8, -19,
	4, -9, -9, -9, -9, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -15, -9, 4, 6, 5,
	-9, 34, 35, 8, 37, 36, -2, 11, -15, 4,
	34, 37, 35, 40, 40, -15, -9, 4, -17, 35,
	35, 34, 35, -9, 35, 36, 8, 36, -15, -17,
	-9, 11, 10, 35, 36, 36, 36, 11, -17, -17,
	36, 36, 11, 11,
}
var yyDef = [...]int{

	3, -2, 1, 10, 4, 0, 2, 8, 0, 0,
	13, 0, 10, 0, 0, 0, 6, 9, 22, 22,
	0, 0, 0, 20, 0, 0, 5, 7, 14, 22,
	23, 24, 0, 0, 0, 0, 0, 21, 0, 0,
	26, 31, 31, 0, 27, 25, 0, 0, 29, 32,
	33, 34, 56, 63, 37, 0, 0, 0, 57, 58,
	0, 0, 0, 0, 0, 28, 31, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 17, 0, 0, 0, 0, 53, 56, 63,
	0, 0, 0, 0, 31, 17, 0, 12, 30, 39,
	64, 40, 41, 42, 43, 44, 45, 46, 47, 48,
	49, 50, 51, 52, 35, 0, 18, 59, 0, 0,
	36, 17, 54, 0, 0, 31, 0, 11, 0, 0,
	0, 0, 55, 60, 61, 0, 0, 0, 0, 62,
	15, 17, 65, 19, 38, 31, 0, 31, 0, 0,
	0, 68, 0, 16, 31, 31, 31, 66, 0, 0,
	31, 31, 67, 69,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 31, 3, 3,
	34, 35, 29, 27, 37, 28, 38, 30, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 36,
	33, 41, 32, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 39, 3, 40, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 42,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26,
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
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:98
		{
			yyVAL.iAstNode = &astAssignment{"=", &astValue{yyDollar[1].string}, yyDollar[3].iAstNode}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:99
		{
			yyVAL.iAstNode = &astAssignment{":=", &astValue{yyDollar[1].string}, yyDollar[3].iAstNode}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:100
		{
			yyVAL.iAstNode = &astWriteContent{}
		}
	case 38:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:102
		{
			yyVAL.iAstNode = &astProcessTemplate{yyDollar[2].string, yyDollar[4].astList}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:103
		{
			yyDollar[3].astFilter.value = yyDollar[1].iAstNode
			yyVAL.iAstNode = &astWriteString{yyDollar[3].astFilter}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:106
		{
			yyVAL.iAstNode = &astExpr{">", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:107
		{
			yyVAL.iAstNode = &astExpr{"<", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:108
		{
			yyVAL.iAstNode = &astExpr{"+", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:109
		{
			yyVAL.iAstNode = &astExpr{"-", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:110
		{
			yyVAL.iAstNode = &astExpr{"*", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:111
		{
			yyVAL.iAstNode = &astExpr{"/", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:112
		{
			yyVAL.iAstNode = &astExpr{"%", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:113
		{
			yyVAL.iAstNode = &astExpr{"==", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:114
		{
			yyVAL.iAstNode = &astExpr{"!=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:115
		{
			yyVAL.iAstNode = &astExpr{">=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:116
		{
			yyVAL.iAstNode = &astExpr{"<=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:117
		{
			yyVAL.iAstNode = &astExpr{"||", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:118
		{
			yyVAL.iAstNode = &astExpr{"&&", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:119
		{
			yyVAL.iAstNode = &astExpr{"!", nil, yyDollar[2].iAstNode}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:120
		{
			yyVAL.iAstNode = &astParenthesis{yyDollar[2].iAstNode}
		}
	case 55:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:122
		{
			yyVAL.iAstNode = &astFunc{yyDollar[1].string, yyDollar[3].astList}
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:123
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:124
		{
			yyVAL.iAstNode = &astString{yyDollar[1].string}
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:125
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:127
		{
			yyVAL.string = yyDollar[1].string + "." + yyDollar[3].string
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:128
		{
			yyVAL.string = yyDollar[1].string + "[" + yyDollar[3].string + "]"
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:129
		{
			yyVAL.string = yyDollar[1].string + "[" + yyDollar[3].string + "]"
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:130
		{
			yyVAL.string = "*(" + yyDollar[3].string + ")"
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:132
		{
			yyVAL.string = yyDollar[1].string
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:135
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:136
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 66:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line template.y:140
		{
			yyVAL.iAstNode = &astLoop{"_", yyDollar[2].string, yyDollar[4].iAstNode, yyDollar[6].astList}
		}
	case 67:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:142
		{
			yyVAL.iAstNode = &astLoop{yyDollar[2].string, yyDollar[4].string, yyDollar[6].iAstNode, yyDollar[8].astList}
		}
	case 68:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:144
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, nil}
		}
	case 69:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:146
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, yyDollar[8].astList}
		}
	}
	goto yystack /* stack new state and value */
}
