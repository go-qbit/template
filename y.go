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
	"'.'",
	"'|'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line template.y:143

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 66
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 371

var yyAct = [...]int{

	49, 109, 47, 85, 57, 84, 46, 38, 74, 75,
	76, 77, 78, 79, 48, 70, 71, 72, 73, 68,
	69, 126, 74, 75, 76, 77, 78, 79, 67, 70,
	71, 72, 73, 68, 69, 31, 125, 143, 39, 74,
	75, 76, 77, 78, 79, 63, 70, 71, 72, 73,
	68, 69, 149, 148, 135, 81, 82, 144, 130, 33,
	142, 85, 87, 84, 137, 132, 133, 32, 123, 95,
	96, 97, 98, 99, 100, 101, 102, 103, 104, 105,
	106, 92, 89, 131, 110, 123, 56, 58, 59, 114,
	66, 110, 120, 117, 124, 34, 123, 29, 28, 20,
	29, 21, 42, 41, 92, 26, 12, 54, 110, 122,
	62, 121, 107, 23, 55, 127, 115, 90, 88, 129,
	56, 58, 59, 60, 134, 61, 141, 140, 83, 19,
	18, 11, 118, 52, 53, 22, 36, 139, 138, 43,
	5, 54, 136, 37, 62, 27, 146, 147, 55, 7,
	10, 128, 92, 113, 112, 25, 62, 92, 8, 16,
	9, 111, 17, 92, 92, 74, 75, 76, 77, 78,
	79, 94, 70, 71, 72, 73, 68, 69, 86, 80,
	116, 74, 75, 76, 77, 78, 79, 65, 70, 71,
	72, 73, 68, 69, 64, 108, 56, 58, 59, 60,
	45, 61, 44, 151, 40, 24, 14, 13, 93, 52,
	53, 35, 6, 15, 4, 51, 50, 54, 3, 2,
	62, 1, 30, 0, 55, 56, 58, 59, 60, 0,
	61, 0, 150, 0, 0, 0, 0, 0, 52, 53,
	0, 0, 0, 0, 0, 0, 54, 0, 0, 62,
	0, 0, 0, 55, 56, 58, 59, 60, 0, 61,
	0, 145, 0, 0, 0, 0, 0, 52, 53, 0,
	0, 0, 0, 0, 0, 54, 0, 0, 62, 0,
	0, 0, 55, 74, 75, 76, 77, 78, 79, 0,
	70, 71, 72, 73, 68, 69, 56, 58, 59, 60,
	0, 61, 0, 119, 0, 0, 0, 0, 0, 52,
	53, 0, 0, 0, 0, 0, 0, 54, 0, 0,
	62, 0, 0, 0, 55, 56, 58, 59, 60, 0,
	61, 0, 91, 0, 0, 0, 0, 0, 52, 53,
	0, 0, 56, 58, 59, 60, 54, 61, 0, 62,
	0, 0, 0, 55, 0, 52, 53, 0, 0, 0,
	0, 0, 0, 54, 0, 0, 62, 0, 0, 0,
	55,
}
var yyPact = [...]int{

	126, -1000, -1000, 145, -1000, 99, 72, -1000, 203, 202,
	-1000, 154, 145, 98, 97, 66, -1000, -1000, 201, 201,
	71, 140, 65, -1000, 31, 62, -1000, -1000, 120, 201,
	-1000, -31, 1, 200, 69, 68, 124, -1000, 198, 196,
	-32, 338, 338, 190, -1000, -1000, 183, 56, -1000, -11,
	-1000, -1000, -1000, 175, 82, 82, 96, -33, -1000, -1000,
	174, 82, 86, 48, 85, -1000, 321, 167, 82, 82,
	82, 82, 82, 82, 82, 82, 82, 82, 82, 82,
	80, -1000, 162, 82, 157, 148, 81, 146, 128, 292,
	82, -1000, -1000, -1000, 79, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 82, -1000, 61,
	264, -1000, -1, -16, 82, 147, 338, 25, -1000, -1000,
	50, 32, 33, 82, -1000, -1000, -1000, 20, 134, 30,
	-1000, -1000, -1000, -1000, 264, 338, 82, 116, 26, 3,
	-1000, 23, 250, 338, 338, -1000, 19, 18, 221, 192,
	-1000, -1000,
}
var yyPgo = [...]int{

	0, 222, 4, 221, 219, 218, 149, 113, 14, 0,
	216, 215, 214, 213, 212, 1, 135, 2, 211, 208,
}
var yyR1 = [...]int{

	0, 3, 4, 5, 5, 12, 13, 13, 14, 14,
	6, 6, 6, 6, 18, 18, 15, 15, 15, 16,
	16, 7, 7, 1, 1, 1, 1, 1, 17, 17,
	8, 8, 8, 8, 8, 8, 8, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 2, 2, 2, 2, 2,
	19, 19, 10, 10, 11, 11,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 1, 5, 1, 3, 1, 3,
	0, 10, 9, 1, 0, 6, 0, 1, 3, 1,
	3, 0, 2, 1, 3, 2, 3, 4, 1, 3,
	0, 1, 1, 1, 1, 5, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	3, 4, 1, 1, 1, 3, 4, 4, 4, 1,
	1, 3, 8, 10, 6, 10,
}
var yyChk = [...]int{

	-1000, -3, -4, -5, -12, 14, -14, -6, 13, 15,
	5, 32, 34, 4, 4, -13, 5, -6, 32, 32,
	33, 35, -16, -7, 4, -16, 34, 5, 33, 35,
	-1, 4, 36, 28, 33, -18, 16, -7, 38, 37,
	4, 34, 34, 15, 4, 4, 38, -17, -8, -9,
	-10, -11, 17, 18, 25, 32, 4, -2, 5, 6,
	7, 9, 28, -17, 4, 4, 34, 39, 30, 31,
	26, 27, 28, 29, 19, 20, 21, 22, 23, 24,
	4, -9, -9, 32, 38, 36, 4, -9, 32, 34,
	32, 11, -8, -19, 4, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, 32, 33, -15,
	-9, 4, 6, 5, 8, 35, 34, -2, 4, 11,
	-15, 32, -15, 35, 33, 37, 37, -9, 4, -17,
	33, 33, 33, 33, -9, 34, 8, 34, -17, -9,
	11, 10, 34, 34, 34, 11, -17, -17, 34, 34,
	11, 11,
}
var yyDef = [...]int{

	3, -2, 1, 10, 4, 0, 2, 8, 0, 0,
	13, 0, 10, 0, 0, 0, 6, 9, 21, 21,
	0, 0, 0, 19, 0, 0, 5, 7, 14, 21,
	22, 23, 0, 0, 0, 0, 0, 20, 0, 0,
	25, 30, 30, 0, 26, 24, 0, 0, 28, 31,
	32, 33, 34, 0, 0, 0, 59, 52, 53, 54,
	0, 0, 0, 0, 0, 27, 30, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 49, 0, 16, 0, 0, 0, 0, 0, 30,
	16, 12, 29, 36, 60, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 16, 50, 0,
	17, 55, 0, 0, 0, 0, 30, 0, 59, 11,
	0, 0, 0, 0, 51, 56, 57, 0, 0, 0,
	58, 15, 61, 35, 18, 30, 0, 30, 0, 0,
	64, 0, 30, 30, 30, 62, 0, 0, 30, 30,
	63, 65,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	32, 33, 28, 26, 35, 27, 38, 29, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 34,
	31, 3, 30, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 36, 3, 37, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 39,
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
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:85
		{
			yyVAL.string = yyDollar[1].string + "." + yyDollar[3].string
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:87
		{
			yyVAL.string = "*" + yyDollar[2].string + "." + yyDollar[4].string
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:89
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:90
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 30:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:92
		{
			yyVAL.iAstNode = nil
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:93
		{
			yyVAL.iAstNode = &astWriteValue{yyDollar[1].iAstNode}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:94
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:95
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:96
		{
			yyVAL.iAstNode = &astWriteContent{}
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:98
		{
			yyVAL.iAstNode = &astProcessTemplate{yyDollar[2].string, yyDollar[4].astList}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:99
		{
			yyDollar[3].astFilter.value = yyDollar[1].iAstNode
			yyVAL.iAstNode = &astWriteString{yyDollar[3].astFilter}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:102
		{
			yyVAL.iAstNode = &astExpr{">", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:103
		{
			yyVAL.iAstNode = &astExpr{"<", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:104
		{
			yyVAL.iAstNode = &astExpr{"+", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:105
		{
			yyVAL.iAstNode = &astExpr{"-", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:106
		{
			yyVAL.iAstNode = &astExpr{"*", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:107
		{
			yyVAL.iAstNode = &astExpr{"/", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:108
		{
			yyVAL.iAstNode = &astExpr{"==", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:109
		{
			yyVAL.iAstNode = &astExpr{"!=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:110
		{
			yyVAL.iAstNode = &astExpr{">=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:111
		{
			yyVAL.iAstNode = &astExpr{"<=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:112
		{
			yyVAL.iAstNode = &astExpr{"||", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:113
		{
			yyVAL.iAstNode = &astExpr{"&&", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:114
		{
			yyVAL.iAstNode = &astExpr{"!", nil, yyDollar[2].iAstNode}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:115
		{
			yyVAL.iAstNode = &astParenthesis{yyDollar[2].iAstNode}
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:117
		{
			yyVAL.iAstNode = &astFunc{yyDollar[1].string, yyDollar[3].astList}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:118
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:119
		{
			yyVAL.iAstNode = &astString{yyDollar[1].string}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:120
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:122
		{
			yyVAL.string = yyDollar[1].string + "." + yyDollar[3].string
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:123
		{
			yyVAL.string = yyDollar[1].string + "[" + yyDollar[3].string + "]"
		}
	case 57:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:124
		{
			yyVAL.string = yyDollar[1].string + "[" + yyDollar[3].string + "]"
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:125
		{
			yyVAL.string = "*(" + yyDollar[3].string + ")"
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:127
		{
			yyVAL.string = yyDollar[1].string
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:130
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:131
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 62:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line template.y:135
		{
			yyVAL.iAstNode = &astLoop{"_", yyDollar[2].string, yyDollar[4].iAstNode, yyDollar[6].astList}
		}
	case 63:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:137
		{
			yyVAL.iAstNode = &astLoop{yyDollar[2].string, yyDollar[4].string, yyDollar[6].iAstNode, yyDollar[8].astList}
		}
	case 64:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:139
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, nil}
		}
	case 65:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:141
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, yyDollar[8].astList}
		}
	}
	goto yystack /* stack new state and value */
}
