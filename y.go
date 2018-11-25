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
const COMMENT = 57371

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
	"COMMENT",
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
	"'{'",
	"'}'",
	"'='",
	"'|'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line template.y:163

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 70,
	26, 46,
	40, 46,
	46, 46,
	-2, 77,
}

const yyNprod = 87
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 501

var yyAct = [...]int{

	56, 55, 136, 54, 65, 33, 128, 113, 85, 86,
	87, 88, 89, 90, 95, 49, 43, 141, 109, 80,
	81, 82, 83, 84, 78, 79, 85, 86, 87, 88,
	89, 90, 92, 34, 52, 143, 77, 80, 81, 82,
	83, 84, 78, 79, 99, 100, 93, 150, 152, 142,
	158, 151, 91, 41, 101, 202, 42, 75, 102, 103,
	187, 36, 155, 163, 96, 98, 102, 103, 132, 198,
	105, 35, 133, 97, 66, 67, 196, 191, 112, 115,
	116, 117, 118, 119, 120, 121, 122, 123, 124, 125,
	126, 127, 129, 129, 63, 185, 107, 155, 186, 130,
	108, 71, 137, 134, 140, 169, 64, 155, 137, 188,
	146, 145, 112, 176, 85, 86, 87, 88, 89, 90,
	164, 156, 155, 155, 110, 80, 81, 82, 83, 84,
	78, 79, 76, 137, 197, 153, 70, 66, 67, 68,
	53, 69, 159, 203, 161, 157, 46, 29, 162, 60,
	61, 12, 37, 168, 32, 172, 171, 63, 31, 22,
	32, 23, 166, 138, 71, 38, 137, 178, 177, 64,
	170, 137, 165, 179, 106, 181, 182, 180, 112, 21,
	20, 11, 26, 97, 50, 39, 47, 5, 24, 16,
	112, 194, 25, 193, 167, 195, 139, 174, 112, 160,
	112, 200, 154, 45, 112, 85, 86, 87, 88, 89,
	90, 71, 51, 30, 28, 40, 80, 81, 82, 83,
	84, 78, 79, 7, 10, 190, 85, 86, 87, 88,
	89, 90, 8, 148, 9, 147, 19, 80, 81, 82,
	83, 84, 78, 79, 18, 17, 189, 85, 86, 87,
	88, 89, 90, 59, 114, 131, 104, 94, 80, 81,
	82, 83, 84, 78, 79, 74, 73, 175, 85, 86,
	87, 88, 89, 90, 6, 72, 48, 44, 27, 80,
	81, 82, 83, 84, 78, 79, 14, 13, 173, 85,
	86, 87, 88, 89, 90, 15, 4, 58, 57, 62,
	80, 81, 82, 83, 84, 78, 79, 3, 2, 144,
	85, 86, 87, 88, 89, 90, 1, 0, 0, 0,
	0, 80, 81, 82, 83, 84, 78, 79, 0, 135,
	70, 66, 67, 68, 0, 69, 184, 183, 0, 0,
	0, 0, 0, 60, 61, 0, 0, 0, 0, 0,
	0, 63, 0, 0, 0, 0, 0, 0, 71, 0,
	0, 0, 0, 64, 85, 86, 87, 88, 89, 90,
	0, 0, 0, 0, 0, 80, 81, 82, 83, 84,
	78, 79, 70, 66, 67, 68, 0, 69, 0, 201,
	0, 0, 0, 0, 0, 60, 61, 0, 0, 70,
	66, 67, 68, 63, 69, 0, 199, 0, 0, 0,
	71, 0, 60, 61, 0, 64, 70, 66, 67, 68,
	63, 69, 0, 192, 0, 0, 0, 71, 0, 60,
	61, 0, 64, 70, 66, 67, 68, 63, 69, 0,
	149, 0, 0, 0, 71, 0, 60, 61, 0, 64,
	70, 66, 67, 68, 63, 69, 0, 111, 0, 0,
	0, 71, 0, 60, 61, 0, 64, 70, 66, 67,
	68, 63, 69, 0, 0, 0, 0, 0, 71, 0,
	60, 61, 0, 64, 0, 0, 0, 0, 63, 0,
	0, 0, 0, 0, 0, 71, 0, 0, 0, 0,
	64,
}
var yyPact = [...]int{

	173, -1000, -1000, 219, -1000, 144, 112, -1000, 283, 282,
	-1000, 240, 219, 143, 142, 121, -1000, -1000, 183, -1000,
	274, 274, 108, 240, -1000, 120, -1000, 29, 114, -1000,
	-1000, 169, 274, -1000, 12, -27, 273, 169, 107, 171,
	-1000, 272, -30, 180, -7, 101, 463, 271, -1000, -1000,
	-1000, 262, 261, 463, 93, -1000, -11, -1000, -1000, 6,
	-1000, 253, -33, 69, 69, 17, -1000, -1000, 252, 69,
	-1000, 137, 59, -23, -1000, 85, 446, 250, 69, 69,
	69, 69, 69, 69, 69, 69, 69, 69, 69, 69,
	69, 69, 69, 251, 31, 250, -1000, -1000, 291, -1000,
	-1000, 69, 159, 69, 9, 270, 179, 69, 231, 229,
	429, -1000, -1000, -1000, 10, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 8, 345,
	8, -1000, 69, 198, -1000, -1000, 83, 345, -1000, 29,
	7, 69, 195, 69, 463, 25, 82, 135, -1000, -1000,
	124, 190, 69, 67, 133, 69, -1000, 117, -1000, 249,
	189, 228, 74, -1000, -1000, 69, -1000, 130, 345, -1000,
	69, 345, -1000, 463, 69, 69, 326, 57, 60, 22,
	70, 207, 186, -1000, 38, -1000, -1000, -1000, 412, 463,
	69, 463, -1000, 37, 95, 30, 395, 463, 378, -1000,
	16, -1000, 132, -1000,
}
var yyPgo = [...]int{

	0, 5, 316, 308, 307, 189, 223, 182, 1, 299,
	0, 298, 297, 4, 296, 295, 274, 2, 192, 3,
	253, 6, 165, 7,
}
var yyR1 = [...]int{

	0, 2, 3, 4, 4, 14, 15, 15, 5, 5,
	16, 16, 6, 6, 6, 6, 22, 22, 22, 17,
	17, 17, 18, 18, 7, 7, 1, 1, 1, 1,
	1, 1, 1, 1, 19, 19, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 20, 20, 21, 21,
	9, 9, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 13, 13, 13, 13, 13, 23, 23,
	23, 23, 11, 11, 11, 12, 12,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 1, 5, 1, 3, 1, 2,
	1, 3, 0, 10, 10, 1, 0, 6, 8, 0,
	1, 3, 1, 3, 0, 2, 1, 3, 2, 3,
	4, 3, 4, 6, 1, 3, 0, 1, 1, 1,
	3, 3, 1, 5, 7, 1, 1, 3, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 3, 2, 2, 4,
	1, 1, 1, 3, 4, 4, 5, 1, 1, 3,
	3, 5, 8, 10, 12, 6, 10,
}
var yyChk = [...]int{

	-1000, -2, -3, -4, -14, 14, -16, -6, 13, 15,
	5, 37, 39, 4, 4, -15, -5, 5, 4, -6,
	37, 37, 38, 40, 5, -18, -7, 4, -18, 39,
	-5, 38, 40, -1, 4, 42, 32, 38, -22, 16,
	-7, 41, 44, 43, 4, -22, 39, 15, 4, 45,
	4, 32, 41, 39, -19, -8, -10, -11, -12, -20,
	17, 18, -9, 25, 37, -13, 5, 6, 7, 9,
	4, 32, 4, 4, 4, -19, 39, 47, 35, 36,
	30, 31, 32, 33, 34, 19, 20, 21, 22, 23,
	24, 46, 26, 40, 4, 47, -10, 4, -10, 27,
	28, 37, 41, 42, 4, -10, 37, 37, 41, 41,
	39, 11, -8, -23, 4, -10, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, -10, -10, -21, -10,
	-21, 4, 37, 41, -23, 38, -17, -10, 4, 37,
	-10, 8, 40, 26, 39, -13, -17, 4, 4, 11,
	37, 41, 40, -17, 4, 40, 38, -1, 43, -10,
	4, -10, -19, 38, 38, 37, 38, 4, -10, 38,
	37, -10, 38, 39, 8, 39, 39, -17, 37, -17,
	-19, -10, -10, 11, 10, 38, 38, 38, 39, 39,
	39, 39, 11, -19, -10, -19, 39, 39, 39, 11,
	-19, 11, 39, 11,
}
var yyDef = [...]int{

	3, -2, 1, 12, 4, 0, 2, 10, 0, 0,
	15, 0, 12, 0, 0, 0, 6, 8, 0, 11,
	24, 24, 0, 0, 9, 0, 22, 0, 0, 5,
	7, 16, 24, 25, 26, 0, 0, 16, 0, 0,
	23, 0, 0, 0, 28, 0, 36, 0, 29, 31,
	27, 0, 0, 36, 0, 34, 37, 38, 39, 0,
	42, 0, 45, 0, 0, 70, 71, 72, 0, 0,
	-2, 0, 0, 32, 30, 0, 36, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 65, 77, 0, 67,
	68, 19, 0, 0, 0, 0, 0, 19, 0, 0,
	36, 13, 35, 50, 78, 52, 53, 54, 55, 56,
	57, 58, 59, 60, 61, 62, 63, 64, 40, 48,
	41, 47, 19, 0, 51, 66, 0, 20, 73, 0,
	0, 0, 0, 0, 36, 0, 0, 0, 33, 14,
	0, 0, 0, 0, 0, 0, 69, 0, 74, 0,
	0, 0, 0, 75, 17, 19, 79, 80, 49, 43,
	19, 21, 76, 36, 0, 0, 36, 0, 0, 0,
	0, 0, 0, 85, 0, 18, 81, 44, 36, 36,
	0, 36, 82, 0, 0, 0, 36, 36, 36, 83,
	0, 86, 36, 84,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 34, 3, 3,
	37, 38, 32, 30, 40, 31, 41, 33, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 39,
	36, 46, 35, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 42, 3, 43, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 44, 47, 45,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29,
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
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:42
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:44
		{
			yyVAL.iAstNode = &astImport{"", yyDollar[1].string}
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:45
		{
			yyVAL.iAstNode = &astImport{yyDollar[1].string, yyDollar[2].string}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:47
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:48
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 12:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:50
		{
			yyVAL.iAstNode = nil
		}
	case 13:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:52
		{
			yyVAL.iAstNode = &astTemplate{yyDollar[2].string, yyDollar[4].astList, yyDollar[6].astUseWrapper, yyDollar[8].astList}
		}
	case 14:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:54
		{
			yyVAL.iAstNode = &astWrapper{yyDollar[2].string, yyDollar[4].astList, yyDollar[6].astUseWrapper, yyDollar[8].astList}
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:55
		{
			yyVAL.iAstNode = nil
		}
	case 16:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:57
		{
			yyVAL.astUseWrapper = nil
		}
	case 17:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:59
		{
			yyVAL.astUseWrapper = &astUseWrapper{"", yyDollar[3].string, yyDollar[5].astList}
		}
	case 18:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line template.y:61
		{
			yyVAL.astUseWrapper = &astUseWrapper{yyDollar[3].string, yyDollar[5].string, yyDollar[7].astList}
		}
	case 19:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:63
		{
			yyVAL.astList = nil
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
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:67
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:68
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:70
		{
			yyVAL.iAstNode = nil
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:71
		{
			yyVAL.iAstNode = &astVariableDef{yyDollar[1].string, yyDollar[2].string}
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:73
		{
			yyVAL.string = yyDollar[1].string
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:74
		{
			yyVAL.string = "[]" + yyDollar[3].string
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:75
		{
			yyVAL.string = "*" + yyDollar[2].string
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:76
		{
			yyVAL.string = yyDollar[1].string + "." + yyDollar[3].string
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:78
		{
			yyVAL.string = "*" + yyDollar[2].string + "." + yyDollar[4].string
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:79
		{
			yyVAL.string = yyDollar[1].string + "{}"
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:80
		{
			yyVAL.string = "[]*" + yyDollar[4].string
		}
	case 33:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:82
		{
			yyVAL.string = "[]*" + yyDollar[4].string + "." + yyDollar[6].string
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:84
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:85
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:87
		{
			yyVAL.iAstNode = nil
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:88
		{
			yyVAL.iAstNode = &astWriteValue{yyDollar[1].iAstNode}
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:89
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:90
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:92
		{
			yyVAL.iAstNode = &astAssignment{"=", yyDollar[1].astList, yyDollar[3].astList}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:94
		{
			yyVAL.iAstNode = &astAssignment{":=", yyDollar[1].astList, yyDollar[3].astList}
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:95
		{
			yyVAL.iAstNode = &astWriteContent{}
		}
	case 43:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:97
		{
			yyVAL.iAstNode = &astProcessTemplate{"", yyDollar[2].string, yyDollar[4].astList}
		}
	case 44:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line template.y:99
		{
			yyVAL.iAstNode = &astProcessTemplate{yyDollar[2].string, yyDollar[4].string, yyDollar[6].astList}
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:100
		{
			yyVAL.iAstNode = &astWriteString{yyDollar[1].iAstNode}
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:102
		{
			yyVAL.astList = &astList{[]iAstNode{&astValue{yyDollar[1].string}}}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:104
		{
			yyVAL.astList.Add(&astValue{yyDollar[3].string})
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:106
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:107
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:109
		{
			yyDollar[3].astFilter.value = yyDollar[1].iAstNode
			yyVAL.iAstNode = yyDollar[3].astFilter
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:110
		{
			yyDollar[3].astFilter.value = yyDollar[1].iAstNode
			yyVAL.iAstNode = yyDollar[3].astFilter
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:113
		{
			yyVAL.iAstNode = &astExpr{">", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:114
		{
			yyVAL.iAstNode = &astExpr{"<", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:115
		{
			yyVAL.iAstNode = &astExpr{"+", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:116
		{
			yyVAL.iAstNode = &astExpr{"-", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:117
		{
			yyVAL.iAstNode = &astExpr{"*", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:118
		{
			yyVAL.iAstNode = &astExpr{"/", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:119
		{
			yyVAL.iAstNode = &astExpr{"%", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:120
		{
			yyVAL.iAstNode = &astExpr{"==", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:121
		{
			yyVAL.iAstNode = &astExpr{"!=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:122
		{
			yyVAL.iAstNode = &astExpr{">=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:123
		{
			yyVAL.iAstNode = &astExpr{"<=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:124
		{
			yyVAL.iAstNode = &astExpr{"||", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:125
		{
			yyVAL.iAstNode = &astExpr{"&&", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:126
		{
			yyVAL.iAstNode = &astExpr{"!", nil, yyDollar[2].iAstNode}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:127
		{
			yyVAL.iAstNode = &astParenthesis{yyDollar[2].iAstNode}
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:128
		{
			yyVAL.iAstNode = &astExpr{"++", yyDollar[1].iAstNode, nil}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:129
		{
			yyVAL.iAstNode = &astExpr{"--", yyDollar[1].iAstNode, nil}
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:131
		{
			yyVAL.iAstNode = &astFunc{yyDollar[1].iAstNode, yyDollar[3].astList}
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:132
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:133
		{
			yyVAL.iAstNode = &astString{yyDollar[1].string}
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:134
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:136
		{
			yyVAL.iAstNode = &astStructField{yyDollar[1].iAstNode, yyDollar[3].string}
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:137
		{
			yyVAL.iAstNode = &astIndexedVar{yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:138
		{
			yyVAL.iAstNode = &astRef{yyDollar[3].iAstNode}
		}
	case 76:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:141
		{
			yyVAL.iAstNode = &astTypeConv{yyDollar[1].iAstNode, yyDollar[4].string}
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:142
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:145
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:146
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:147
		{
			yyVAL.astFilter = &astFilter{pkg: yyDollar[1].string, name: yyDollar[3].string}
		}
	case 81:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:149
		{
			yyVAL.astFilter = &astFilter{pkg: yyDollar[1].string, name: yyDollar[3].string}
		}
	case 82:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line template.y:153
		{
			yyVAL.iAstNode = &astRangeLoop{"_", yyDollar[2].string, yyDollar[4].iAstNode, yyDollar[6].astList}
		}
	case 83:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:155
		{
			yyVAL.iAstNode = &astRangeLoop{yyDollar[2].string, yyDollar[4].string, yyDollar[6].iAstNode, yyDollar[8].astList}
		}
	case 84:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line template.y:157
		{
			yyVAL.iAstNode = &astLoop{&astExpr{":=", &astValue{yyDollar[2].string}, yyDollar[4].iAstNode}, yyDollar[6].iAstNode, yyDollar[8].iAstNode, yyDollar[10].astList}
		}
	case 85:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:159
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, nil}
		}
	case 86:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:161
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, yyDollar[8].astList}
		}
	}
	goto yystack /* stack new state and value */
}
