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
const INC = 57369
const DEC = 57370

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
	"INC",
	"DEC",
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

//line template.y:146

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 78
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 492

var yyAct = [...]int{

	50, 49, 104, 48, 91, 150, 143, 142, 87, 88,
	46, 120, 77, 78, 79, 80, 81, 82, 39, 31,
	93, 38, 72, 73, 74, 75, 76, 70, 71, 77,
	78, 79, 80, 81, 82, 187, 172, 69, 140, 72,
	73, 74, 75, 76, 70, 71, 33, 183, 182, 181,
	170, 67, 140, 84, 85, 176, 32, 138, 173, 92,
	95, 139, 86, 53, 97, 161, 87, 88, 53, 83,
	103, 106, 107, 108, 109, 110, 111, 112, 113, 114,
	115, 116, 117, 118, 119, 101, 130, 121, 68, 53,
	125, 126, 84, 85, 128, 127, 47, 94, 60, 61,
	121, 86, 42, 103, 132, 87, 88, 54, 60, 61,
	62, 135, 63, 169, 168, 26, 171, 131, 58, 134,
	55, 56, 53, 156, 64, 140, 99, 121, 58, 59,
	100, 146, 12, 148, 64, 153, 163, 149, 144, 59,
	157, 155, 151, 141, 140, 140, 34, 28, 29, 29,
	20, 152, 21, 121, 53, 98, 19, 18, 121, 11,
	166, 167, 165, 103, 162, 35, 94, 89, 23, 164,
	22, 36, 43, 5, 7, 103, 179, 159, 178, 53,
	180, 27, 53, 103, 16, 103, 185, 17, 154, 103,
	25, 10, 147, 64, 53, 53, 145, 53, 37, 8,
	41, 9, 53, 53, 53, 124, 123, 136, 53, 77,
	78, 79, 80, 81, 82, 105, 122, 96, 90, 72,
	73, 74, 75, 76, 70, 71, 66, 65, 175, 77,
	78, 79, 80, 81, 82, 45, 44, 40, 24, 72,
	73, 74, 75, 76, 70, 71, 14, 13, 174, 77,
	78, 79, 80, 81, 82, 6, 15, 4, 52, 72,
	73, 74, 75, 76, 70, 71, 51, 57, 160, 77,
	78, 79, 80, 81, 82, 3, 2, 1, 30, 72,
	73, 74, 75, 76, 70, 71, 0, 0, 158, 77,
	78, 79, 80, 81, 82, 0, 0, 0, 0, 72,
	73, 74, 75, 76, 70, 71, 0, 0, 133, 77,
	78, 79, 80, 81, 82, 0, 0, 0, 0, 72,
	73, 74, 75, 76, 70, 71, 0, 129, 77, 78,
	79, 80, 81, 82, 0, 0, 0, 0, 72, 73,
	74, 75, 76, 70, 71, 54, 60, 61, 62, 0,
	63, 0, 188, 0, 0, 0, 0, 0, 55, 56,
	54, 60, 61, 62, 0, 63, 58, 186, 0, 0,
	0, 0, 64, 55, 56, 0, 0, 59, 0, 0,
	0, 58, 0, 54, 60, 61, 62, 64, 63, 0,
	184, 0, 59, 0, 0, 0, 55, 56, 54, 60,
	61, 62, 0, 63, 58, 177, 0, 0, 0, 0,
	64, 55, 56, 0, 0, 59, 0, 0, 0, 58,
	0, 54, 60, 61, 62, 64, 63, 0, 137, 0,
	59, 0, 0, 0, 55, 56, 54, 60, 61, 62,
	0, 63, 58, 102, 0, 0, 0, 0, 64, 55,
	56, 0, 0, 59, 0, 0, 0, 58, 0, 54,
	60, 61, 62, 64, 63, 0, 0, 0, 59, 0,
	0, 0, 55, 56, 0, 0, 0, 0, 0, 0,
	58, 0, 0, 0, 0, 0, 64, 0, 0, 0,
	0, 59,
}
var yyPact = [...]int{

	159, -1000, -1000, 186, -1000, 123, 94, -1000, 243, 242,
	-1000, 179, 186, 121, 120, 113, -1000, -1000, 234, 234,
	77, 176, 110, -1000, 15, 109, -1000, -1000, 155, 234,
	-1000, -19, -24, 233, 155, 64, 157, -1000, 232, 231,
	-30, 58, 455, 223, -1000, -1000, 222, 455, 50, -1000,
	-7, -1000, -1000, 26, 141, -1000, 214, -40, 93, 93,
	-1000, -1000, 213, 93, 119, 90, -1000, 47, 432, 211,
	93, 93, 93, 93, 93, 93, 93, 93, 93, 93,
	93, 93, 93, 93, -1000, -1000, 93, 212, 200, 93,
	55, 211, -1000, 65, -1000, 290, 78, 270, 162, 93,
	203, 417, -1000, -1000, -1000, 21, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 309,
	106, 309, -1000, -35, -36, 309, 93, 192, -1000, -1000,
	93, 188, 93, 455, -32, 105, 115, -1000, 98, 184,
	93, -1000, -1000, -1000, 86, 104, 250, 169, 230, 27,
	-1000, -1000, 93, -1000, 100, 309, -1000, 93, 455, 93,
	93, 103, 13, 79, -1, 20, 210, 190, -1000, 17,
	-1000, -1000, -1000, 394, 455, 93, 455, -1000, 11, 10,
	9, 379, 455, 356, -1000, -3, -1000, 341, -1000,
}
var yyPgo = [...]int{

	0, 278, 20, 277, 276, 275, 174, 168, 1, 267,
	0, 266, 258, 257, 256, 255, 11, 170, 3, 165,
	2,
}
var yyR1 = [...]int{

	0, 3, 4, 5, 5, 13, 14, 14, 15, 15,
	6, 6, 6, 6, 19, 19, 19, 16, 16, 16,
	17, 17, 7, 7, 1, 1, 1, 1, 1, 18,
	18, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 9, 9, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 2, 2, 2, 2, 2, 20,
	20, 20, 20, 11, 11, 11, 12, 12,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 1, 5, 1, 3, 1, 3,
	0, 10, 10, 1, 0, 6, 8, 0, 1, 3,
	1, 3, 0, 2, 1, 3, 2, 3, 4, 1,
	3, 0, 1, 1, 1, 3, 3, 1, 5, 7,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 3, 2, 2,
	4, 1, 1, 1, 3, 4, 4, 4, 1, 1,
	3, 3, 5, 8, 10, 12, 6, 10,
}
var yyChk = [...]int{

	-1000, -3, -4, -5, -13, 14, -15, -6, 13, 15,
	5, 36, 38, 4, 4, -14, 5, -6, 36, 36,
	37, 39, -17, -7, 4, -17, 38, 5, 37, 39,
	-1, 4, 41, 31, 37, -19, 16, -7, 40, 42,
	4, -19, 38, 15, 4, 4, 40, 38, -18, -8,
	-10, -11, -12, -2, 4, 17, 18, -9, 25, 36,
	5, 6, 7, 9, 31, 4, 4, -18, 38, 44,
	34, 35, 29, 30, 31, 32, 33, 19, 20, 21,
	22, 23, 24, 43, 27, 28, 36, 40, 41, 26,
	4, 44, -10, -2, 4, -10, 4, -10, 36, 36,
	40, 38, 11, -8, -20, 4, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, -10, -10, -10, -10,
	-16, -10, 4, 6, 5, -10, 36, 40, -20, 37,
	8, 39, 26, 38, -2, -16, 4, 11, 36, 40,
	39, 37, 42, 42, -16, 4, -10, 4, -10, -18,
	37, 37, 36, 37, 4, -10, 37, 36, 38, 8,
	38, 38, -16, 36, -16, -18, -10, -10, 11, 10,
	37, 37, 37, 38, 38, 38, 38, 11, -18, -10,
	-18, 38, 38, 38, 11, -18, 11, 38, 11,
}
var yyDef = [...]int{

	3, -2, 1, 10, 4, 0, 2, 8, 0, 0,
	13, 0, 10, 0, 0, 0, 6, 9, 22, 22,
	0, 0, 0, 20, 0, 0, 5, 7, 14, 22,
	23, 24, 0, 0, 14, 0, 0, 21, 0, 0,
	26, 0, 31, 0, 27, 25, 0, 31, 0, 29,
	32, 33, 34, 61, 68, 37, 0, 40, 0, 0,
	62, 63, 0, 0, 0, 0, 28, 0, 31, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 58, 59, 17, 0, 0, 0,
	0, 0, 56, 61, 68, 0, 0, 0, 0, 17,
	0, 31, 11, 30, 41, 69, 43, 44, 45, 46,
	47, 48, 49, 50, 51, 52, 53, 54, 55, 35,
	0, 18, 64, 0, 0, 36, 17, 0, 42, 57,
	0, 0, 0, 31, 0, 0, 0, 12, 0, 0,
	0, 60, 65, 66, 0, 0, 0, 0, 0, 0,
	67, 15, 17, 70, 71, 19, 38, 17, 31, 0,
	0, 31, 0, 0, 0, 0, 0, 0, 76, 0,
	16, 72, 39, 31, 31, 0, 31, 73, 0, 0,
	0, 31, 31, 31, 74, 0, 77, 31, 75,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 33, 3, 3,
	36, 37, 31, 29, 39, 30, 40, 32, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 38,
	35, 43, 34, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 41, 3, 42, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 44,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28,
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
		//line template.y:31
		{
			yylex.(*exprLex).result = yyVAL.iAstNode
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:33
		{
			yyVAL.iAstNode = &astFile{yyDollar[1].iAstNode, yyDollar[2].astList}
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:35
		{
			yyVAL.iAstNode = nil
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:36
		{
			yyVAL.iAstNode = &astHeader{yyDollar[1].astList}
		}
	case 5:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:39
		{
			yyVAL.astList = yyDollar[3].astList
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:41
		{
			yyVAL.astList = &astList{[]iAstNode{&astImport{yyDollar[1].string}}}
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:42
		{
			yyVAL.astList.Add(&astImport{yyDollar[3].string})
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:44
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:45
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 10:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:47
		{
			yyVAL.iAstNode = nil
		}
	case 11:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:49
		{
			yyVAL.iAstNode = &astTemplate{yyDollar[2].string, yyDollar[4].astList, yyDollar[6].astUseWrapper, yyDollar[8].astList}
		}
	case 12:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:51
		{
			yyVAL.iAstNode = &astWrapper{yyDollar[2].string, yyDollar[4].astList, yyDollar[6].astUseWrapper, yyDollar[8].astList}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:52
		{
			yyVAL.iAstNode = nil
		}
	case 14:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:54
		{
			yyVAL.astUseWrapper = nil
		}
	case 15:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:56
		{
			yyVAL.astUseWrapper = &astUseWrapper{"", yyDollar[3].string, yyDollar[5].astList}
		}
	case 16:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line template.y:58
		{
			yyVAL.astUseWrapper = &astUseWrapper{yyDollar[3].string, yyDollar[5].string, yyDollar[7].astList}
		}
	case 17:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:60
		{
			yyVAL.astList = nil
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:61
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:62
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:64
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:65
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 22:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:67
		{
			yyVAL.iAstNode = nil
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:68
		{
			yyVAL.iAstNode = &astVariableDef{yyDollar[1].string, yyDollar[2].string}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:70
		{
			yyVAL.string = yyDollar[1].string
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:71
		{
			yyVAL.string = "[]" + yyDollar[3].string
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:72
		{
			yyVAL.string = "*" + yyDollar[2].string
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:73
		{
			yyVAL.string = yyDollar[1].string + "." + yyDollar[3].string
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:75
		{
			yyVAL.string = "*" + yyDollar[2].string + "." + yyDollar[4].string
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:77
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:78
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 31:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:80
		{
			yyVAL.iAstNode = nil
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:81
		{
			yyVAL.iAstNode = &astWriteValue{yyDollar[1].iAstNode}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:82
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:83
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:84
		{
			yyVAL.iAstNode = &astAssignment{"=", &astValue{yyDollar[1].string}, yyDollar[3].iAstNode}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:85
		{
			yyVAL.iAstNode = &astAssignment{":=", &astValue{yyDollar[1].string}, yyDollar[3].iAstNode}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:86
		{
			yyVAL.iAstNode = &astWriteContent{}
		}
	case 38:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:88
		{
			yyVAL.iAstNode = &astProcessTemplate{"", yyDollar[2].string, yyDollar[4].astList}
		}
	case 39:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line template.y:90
		{
			yyVAL.iAstNode = &astProcessTemplate{yyDollar[2].string, yyDollar[4].string, yyDollar[6].astList}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:91
		{
			yyVAL.iAstNode = &astWriteString{yyDollar[1].iAstNode}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:93
		{
			yyDollar[3].astFilter.value = yyDollar[1].iAstNode
			yyVAL.iAstNode = yyDollar[3].astFilter
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:94
		{
			yyDollar[3].astFilter.value = yyDollar[1].iAstNode
			yyVAL.iAstNode = yyDollar[3].astFilter
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:97
		{
			yyVAL.iAstNode = &astExpr{">", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:98
		{
			yyVAL.iAstNode = &astExpr{"<", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:99
		{
			yyVAL.iAstNode = &astExpr{"+", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:100
		{
			yyVAL.iAstNode = &astExpr{"-", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:101
		{
			yyVAL.iAstNode = &astExpr{"*", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:102
		{
			yyVAL.iAstNode = &astExpr{"/", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:103
		{
			yyVAL.iAstNode = &astExpr{"%", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:104
		{
			yyVAL.iAstNode = &astExpr{"==", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:105
		{
			yyVAL.iAstNode = &astExpr{"!=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:106
		{
			yyVAL.iAstNode = &astExpr{">=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:107
		{
			yyVAL.iAstNode = &astExpr{"<=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:108
		{
			yyVAL.iAstNode = &astExpr{"||", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:109
		{
			yyVAL.iAstNode = &astExpr{"&&", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:110
		{
			yyVAL.iAstNode = &astExpr{"!", nil, yyDollar[2].iAstNode}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:111
		{
			yyVAL.iAstNode = &astParenthesis{yyDollar[2].iAstNode}
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:112
		{
			yyVAL.iAstNode = &astExpr{"++", &astValue{yyDollar[1].string}, nil}
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:113
		{
			yyVAL.iAstNode = &astExpr{"--", &astValue{yyDollar[1].string}, nil}
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:115
		{
			yyVAL.iAstNode = &astFunc{yyDollar[1].string, yyDollar[3].astList}
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:116
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:117
		{
			yyVAL.iAstNode = &astString{yyDollar[1].string}
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:118
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:120
		{
			yyVAL.string = yyDollar[1].string + "." + yyDollar[3].string
		}
	case 65:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:121
		{
			yyVAL.string = yyDollar[1].string + "[" + yyDollar[3].string + "]"
		}
	case 66:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:122
		{
			yyVAL.string = yyDollar[1].string + "[" + yyDollar[3].string + "]"
		}
	case 67:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:123
		{
			yyVAL.string = "*(" + yyDollar[3].string + ")"
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:125
		{
			yyVAL.string = yyDollar[1].string
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:128
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:129
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:130
		{
			yyVAL.astFilter = &astFilter{pkg: yyDollar[1].string, name: yyDollar[3].string}
		}
	case 72:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:132
		{
			yyVAL.astFilter = &astFilter{pkg: yyDollar[1].string, name: yyDollar[3].string}
		}
	case 73:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line template.y:136
		{
			yyVAL.iAstNode = &astRangeLoop{"_", yyDollar[2].string, yyDollar[4].iAstNode, yyDollar[6].astList}
		}
	case 74:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:138
		{
			yyVAL.iAstNode = &astRangeLoop{yyDollar[2].string, yyDollar[4].string, yyDollar[6].iAstNode, yyDollar[8].astList}
		}
	case 75:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line template.y:140
		{
			yyVAL.iAstNode = &astLoop{&astExpr{":=", &astValue{yyDollar[2].string}, yyDollar[4].iAstNode}, yyDollar[6].iAstNode, yyDollar[8].iAstNode, yyDollar[10].astList}
		}
	case 76:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:142
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, nil}
		}
	case 77:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:144
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, yyDollar[8].astList}
		}
	}
	goto yystack /* stack new state and value */
}
