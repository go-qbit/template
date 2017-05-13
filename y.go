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

//line template.y:161

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 69,
	26, 44,
	40, 44,
	46, 44,
	-2, 76,
}

const yyNprod = 86
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 478

var yyAct = [...]int{

	55, 54, 133, 53, 33, 64, 125, 110, 83, 84,
	85, 86, 87, 88, 93, 49, 156, 155, 43, 78,
	79, 80, 81, 82, 76, 77, 83, 84, 85, 86,
	87, 88, 41, 51, 90, 42, 75, 78, 79, 80,
	81, 82, 76, 77, 149, 200, 195, 184, 91, 147,
	196, 97, 98, 148, 89, 161, 73, 194, 100, 101,
	185, 99, 152, 94, 96, 100, 101, 34, 189, 103,
	129, 95, 65, 66, 130, 186, 109, 112, 113, 114,
	115, 116, 117, 118, 119, 120, 121, 122, 123, 124,
	126, 126, 62, 174, 105, 36, 170, 127, 106, 70,
	134, 131, 176, 107, 63, 35, 134, 74, 144, 109,
	143, 83, 84, 85, 86, 87, 88, 183, 164, 152,
	139, 52, 78, 79, 80, 81, 82, 76, 77, 46,
	134, 188, 150, 29, 69, 65, 66, 67, 141, 68,
	157, 154, 159, 167, 135, 152, 160, 59, 60, 162,
	166, 152, 140, 169, 153, 62, 152, 37, 31, 32,
	32, 22, 70, 23, 134, 12, 175, 63, 168, 134,
	163, 177, 38, 179, 180, 178, 109, 136, 104, 21,
	20, 11, 39, 95, 47, 16, 5, 7, 109, 192,
	25, 191, 24, 193, 172, 165, 109, 26, 109, 198,
	19, 158, 109, 83, 84, 85, 86, 87, 88, 30,
	45, 70, 28, 151, 78, 79, 80, 81, 82, 76,
	77, 138, 137, 187, 83, 84, 85, 86, 87, 88,
	40, 18, 17, 145, 111, 78, 79, 80, 81, 82,
	76, 77, 58, 10, 173, 83, 84, 85, 86, 87,
	88, 8, 128, 9, 102, 92, 78, 79, 80, 81,
	82, 76, 77, 72, 71, 171, 83, 84, 85, 86,
	87, 88, 6, 50, 48, 44, 27, 78, 79, 80,
	81, 82, 76, 77, 14, 13, 142, 83, 84, 85,
	86, 87, 88, 15, 4, 57, 56, 61, 78, 79,
	80, 81, 82, 76, 77, 3, 132, 69, 65, 66,
	67, 2, 68, 182, 181, 1, 0, 0, 0, 0,
	59, 60, 0, 0, 0, 0, 0, 0, 62, 0,
	0, 0, 0, 0, 0, 70, 0, 0, 0, 0,
	63, 83, 84, 85, 86, 87, 88, 0, 0, 0,
	0, 0, 78, 79, 80, 81, 82, 76, 77, 69,
	65, 66, 67, 0, 68, 0, 201, 0, 0, 0,
	0, 0, 59, 60, 0, 0, 69, 65, 66, 67,
	62, 68, 0, 199, 0, 0, 0, 70, 0, 59,
	60, 0, 63, 69, 65, 66, 67, 62, 68, 0,
	197, 0, 0, 0, 70, 0, 59, 60, 0, 63,
	69, 65, 66, 67, 62, 68, 0, 190, 0, 0,
	0, 70, 0, 59, 60, 0, 63, 69, 65, 66,
	67, 62, 68, 0, 146, 0, 0, 0, 70, 0,
	59, 60, 0, 63, 69, 65, 66, 67, 62, 68,
	0, 108, 0, 0, 0, 70, 0, 59, 60, 0,
	63, 0, 0, 0, 0, 62, 0, 0, 0, 0,
	0, 0, 70, 0, 0, 0, 0, 63,
}
var yyPact = [...]int{

	172, -1000, -1000, 238, -1000, 144, 126, -1000, 281, 280,
	-1000, 227, 238, 143, 142, 123, -1000, -1000, 187, -1000,
	272, 272, 94, 227, -1000, 120, -1000, 63, 119, -1000,
	-1000, 166, 272, -1000, -9, -25, 271, 166, 90, 169,
	-1000, 270, -30, 269, -8, 82, 130, 260, -1000, -1000,
	-1000, 259, 130, 68, -1000, -11, -1000, -1000, 8, -1000,
	251, -33, 67, 67, 24, -1000, -1000, 250, 67, -1000,
	141, 57, -1000, 64, 440, 230, 67, 67, 67, 67,
	67, 67, 67, 67, 67, 67, 67, 67, 67, 67,
	67, 248, 33, 230, -1000, -1000, 268, -1000, -1000, 67,
	140, 216, 112, 247, 179, 67, 229, 423, -1000, -1000,
	-1000, 12, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 4, 322, 4, -1000, 67,
	209, -1000, -1000, 116, 322, -1000, 63, -26, -27, 67,
	197, 67, 130, 17, 111, 133, -1000, 80, 191, 67,
	105, 131, 67, -1000, 58, -1000, -1000, 226, 186, 205,
	54, -1000, -1000, 67, -1000, 65, 322, -1000, 67, 322,
	-1000, 130, 67, 67, 303, 79, 9, 22, 36, 184,
	92, -1000, 29, -1000, -1000, -1000, 406, 130, 67, 130,
	-1000, 18, 7, 11, 389, 130, 372, -1000, 6, -1000,
	355, -1000,
}
var yyPgo = [...]int{

	0, 4, 5, 315, 311, 305, 185, 187, 197, 1,
	297, 0, 296, 295, 294, 293, 272, 2, 190, 3,
	242, 6, 172, 7,
}
var yyR1 = [...]int{

	0, 3, 4, 5, 5, 14, 15, 15, 6, 6,
	16, 16, 7, 7, 7, 7, 22, 22, 22, 17,
	17, 17, 18, 18, 8, 8, 1, 1, 1, 1,
	1, 1, 19, 19, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 20, 20, 21, 21, 10, 10,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 2, 2, 2, 2, 2, 2, 23, 23, 23,
	23, 12, 12, 12, 13, 13,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 1, 5, 1, 3, 1, 2,
	1, 3, 0, 10, 10, 1, 0, 6, 8, 0,
	1, 3, 1, 3, 0, 2, 1, 3, 2, 3,
	4, 3, 1, 3, 0, 1, 1, 1, 3, 3,
	1, 5, 7, 1, 1, 3, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 3, 2, 2, 4, 1, 1,
	1, 3, 4, 4, 4, 5, 1, 1, 3, 3,
	5, 8, 10, 12, 6, 10,
}
var yyChk = [...]int{

	-1000, -3, -4, -5, -14, 14, -16, -7, 13, 15,
	5, 37, 39, 4, 4, -15, -6, 5, 4, -7,
	37, 37, 38, 40, 5, -18, -8, 4, -18, 39,
	-6, 38, 40, -1, 4, 42, 32, 38, -22, 16,
	-8, 41, 44, 43, 4, -22, 39, 15, 4, 45,
	4, 41, 39, -19, -9, -11, -12, -13, -20, 17,
	18, -10, 25, 37, -2, 5, 6, 7, 9, 4,
	32, 4, 4, -19, 39, 47, 35, 36, 30, 31,
	32, 33, 34, 19, 20, 21, 22, 23, 24, 46,
	26, 40, 4, 47, -11, 4, -11, 27, 28, 37,
	41, 42, 4, -11, 37, 37, 41, 39, 11, -9,
	-23, 4, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -21, -11, -21, 4, 37,
	41, -23, 38, -17, -11, 4, 37, 6, 5, 8,
	40, 26, 39, -2, -17, 4, 11, 37, 41, 40,
	-17, 4, 40, 38, -1, 43, 43, -11, 4, -11,
	-19, 38, 38, 37, 38, 4, -11, 38, 37, -11,
	38, 39, 8, 39, 39, -17, 37, -17, -19, -11,
	-11, 11, 10, 38, 38, 38, 39, 39, 39, 39,
	11, -19, -11, -19, 39, 39, 39, 11, -19, 11,
	39, 11,
}
var yyDef = [...]int{

	3, -2, 1, 12, 4, 0, 2, 10, 0, 0,
	15, 0, 12, 0, 0, 0, 6, 8, 0, 11,
	24, 24, 0, 0, 9, 0, 22, 0, 0, 5,
	7, 16, 24, 25, 26, 0, 0, 16, 0, 0,
	23, 0, 0, 0, 28, 0, 34, 0, 29, 31,
	27, 0, 34, 0, 32, 35, 36, 37, 0, 40,
	0, 43, 0, 0, 68, 69, 70, 0, 0, -2,
	0, 0, 30, 0, 34, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 63, 76, 0, 65, 66, 19,
	0, 0, 0, 0, 0, 19, 0, 34, 13, 33,
	48, 77, 50, 51, 52, 53, 54, 55, 56, 57,
	58, 59, 60, 61, 62, 38, 46, 39, 45, 19,
	0, 49, 64, 0, 20, 71, 0, 0, 0, 0,
	0, 0, 34, 0, 0, 0, 14, 0, 0, 0,
	0, 0, 0, 67, 0, 72, 73, 0, 0, 0,
	0, 74, 17, 19, 78, 79, 47, 41, 19, 21,
	75, 34, 0, 0, 34, 0, 0, 0, 0, 0,
	0, 84, 0, 18, 80, 42, 34, 34, 0, 34,
	81, 0, 0, 0, 34, 34, 34, 82, 0, 85,
	34, 83,
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
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:81
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:82
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 34:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:84
		{
			yyVAL.iAstNode = nil
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:85
		{
			yyVAL.iAstNode = &astWriteValue{yyDollar[1].iAstNode}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:86
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:87
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:89
		{
			yyVAL.iAstNode = &astAssignment{"=", yyDollar[1].astList, yyDollar[3].astList}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:91
		{
			yyVAL.iAstNode = &astAssignment{":=", yyDollar[1].astList, yyDollar[3].astList}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:92
		{
			yyVAL.iAstNode = &astWriteContent{}
		}
	case 41:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:94
		{
			yyVAL.iAstNode = &astProcessTemplate{"", yyDollar[2].string, yyDollar[4].astList}
		}
	case 42:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line template.y:96
		{
			yyVAL.iAstNode = &astProcessTemplate{yyDollar[2].string, yyDollar[4].string, yyDollar[6].astList}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:97
		{
			yyVAL.iAstNode = &astWriteString{yyDollar[1].iAstNode}
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:99
		{
			yyVAL.astList = &astList{[]iAstNode{&astValue{yyDollar[1].string}}}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:101
		{
			yyVAL.astList.Add(&astValue{yyDollar[3].string})
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:103
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:104
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:106
		{
			yyDollar[3].astFilter.value = yyDollar[1].iAstNode
			yyVAL.iAstNode = yyDollar[3].astFilter
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:107
		{
			yyDollar[3].astFilter.value = yyDollar[1].iAstNode
			yyVAL.iAstNode = yyDollar[3].astFilter
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:110
		{
			yyVAL.iAstNode = &astExpr{">", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:111
		{
			yyVAL.iAstNode = &astExpr{"<", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:112
		{
			yyVAL.iAstNode = &astExpr{"+", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:113
		{
			yyVAL.iAstNode = &astExpr{"-", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:114
		{
			yyVAL.iAstNode = &astExpr{"*", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:115
		{
			yyVAL.iAstNode = &astExpr{"/", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:116
		{
			yyVAL.iAstNode = &astExpr{"%", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:117
		{
			yyVAL.iAstNode = &astExpr{"==", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:118
		{
			yyVAL.iAstNode = &astExpr{"!=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:119
		{
			yyVAL.iAstNode = &astExpr{">=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:120
		{
			yyVAL.iAstNode = &astExpr{"<=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:121
		{
			yyVAL.iAstNode = &astExpr{"||", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:122
		{
			yyVAL.iAstNode = &astExpr{"&&", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:123
		{
			yyVAL.iAstNode = &astExpr{"!", nil, yyDollar[2].iAstNode}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:124
		{
			yyVAL.iAstNode = &astParenthesis{yyDollar[2].iAstNode}
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:125
		{
			yyVAL.iAstNode = &astExpr{"++", &astValue{yyDollar[1].string}, nil}
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:126
		{
			yyVAL.iAstNode = &astExpr{"--", &astValue{yyDollar[1].string}, nil}
		}
	case 67:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:128
		{
			yyVAL.iAstNode = &astFunc{yyDollar[1].string, yyDollar[3].astList}
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:129
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:130
		{
			yyVAL.iAstNode = &astString{yyDollar[1].string}
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:131
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:133
		{
			yyVAL.string = yyDollar[1].string + "." + yyDollar[3].string
		}
	case 72:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:134
		{
			yyVAL.string = yyDollar[1].string + "[" + yyDollar[3].string + "]"
		}
	case 73:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:135
		{
			yyVAL.string = yyDollar[1].string + "[" + yyDollar[3].string + "]"
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:136
		{
			yyVAL.string = "*(" + yyDollar[3].string + ")"
		}
	case 75:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:139
		{
			yyVAL.string = yyDollar[1].string + ".(" + yyDollar[4].string + ")"
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:140
		{
			yyVAL.string = yyDollar[1].string
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:143
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:144
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:145
		{
			yyVAL.astFilter = &astFilter{pkg: yyDollar[1].string, name: yyDollar[3].string}
		}
	case 80:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:147
		{
			yyVAL.astFilter = &astFilter{pkg: yyDollar[1].string, name: yyDollar[3].string}
		}
	case 81:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line template.y:151
		{
			yyVAL.iAstNode = &astRangeLoop{"_", yyDollar[2].string, yyDollar[4].iAstNode, yyDollar[6].astList}
		}
	case 82:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:153
		{
			yyVAL.iAstNode = &astRangeLoop{yyDollar[2].string, yyDollar[4].string, yyDollar[6].iAstNode, yyDollar[8].astList}
		}
	case 83:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line template.y:155
		{
			yyVAL.iAstNode = &astLoop{&astExpr{":=", &astValue{yyDollar[2].string}, yyDollar[4].iAstNode}, yyDollar[6].iAstNode, yyDollar[8].iAstNode, yyDollar[10].astList}
		}
	case 84:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:157
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, nil}
		}
	case 85:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:159
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, yyDollar[8].astList}
		}
	}
	goto yystack /* stack new state and value */
}
