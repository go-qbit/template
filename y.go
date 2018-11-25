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

//line template.y:164

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 71,
	26, 47,
	40, 47,
	46, 47,
	-2, 78,
}

const yyNprod = 88
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 502

var yyAct = [...]int{

	57, 56, 137, 55, 66, 34, 129, 114, 86, 87,
	88, 89, 90, 91, 96, 50, 44, 35, 110, 81,
	82, 83, 84, 85, 79, 80, 86, 87, 88, 89,
	90, 91, 93, 18, 17, 53, 78, 81, 82, 83,
	84, 85, 79, 80, 153, 37, 94, 100, 101, 151,
	159, 42, 92, 152, 43, 36, 164, 102, 76, 103,
	104, 103, 104, 203, 133, 97, 99, 22, 134, 23,
	188, 106, 156, 108, 98, 67, 68, 109, 199, 113,
	116, 117, 118, 119, 120, 121, 122, 123, 124, 125,
	126, 127, 128, 130, 130, 64, 186, 170, 156, 156,
	131, 197, 72, 138, 135, 141, 165, 65, 156, 138,
	192, 147, 146, 113, 189, 86, 87, 88, 89, 90,
	91, 177, 157, 142, 156, 111, 81, 82, 83, 84,
	85, 79, 80, 77, 138, 198, 154, 71, 67, 68,
	69, 144, 70, 160, 204, 162, 158, 54, 47, 163,
	61, 62, 30, 12, 169, 143, 39, 172, 64, 38,
	32, 33, 33, 187, 173, 72, 167, 138, 179, 178,
	65, 139, 138, 171, 180, 166, 182, 183, 181, 113,
	107, 21, 20, 11, 98, 51, 40, 48, 5, 25,
	27, 113, 195, 26, 194, 46, 196, 175, 168, 113,
	161, 113, 201, 7, 140, 113, 86, 87, 88, 89,
	90, 91, 72, 52, 155, 29, 19, 81, 82, 83,
	84, 85, 79, 80, 41, 10, 191, 86, 87, 88,
	89, 90, 91, 8, 149, 9, 18, 17, 81, 82,
	83, 84, 85, 79, 80, 148, 115, 190, 86, 87,
	88, 89, 90, 91, 60, 132, 105, 95, 75, 81,
	82, 83, 84, 85, 79, 80, 74, 73, 176, 86,
	87, 88, 89, 90, 91, 6, 49, 45, 28, 14,
	81, 82, 83, 84, 85, 79, 80, 13, 15, 174,
	86, 87, 88, 89, 90, 91, 4, 59, 58, 63,
	3, 81, 82, 83, 84, 85, 79, 80, 2, 1,
	145, 86, 87, 88, 89, 90, 91, 0, 0, 16,
	0, 0, 81, 82, 83, 84, 85, 79, 80, 0,
	136, 71, 67, 68, 69, 24, 70, 185, 184, 0,
	0, 0, 0, 31, 61, 62, 0, 0, 0, 0,
	0, 0, 64, 0, 0, 0, 0, 0, 0, 72,
	0, 0, 0, 0, 65, 86, 87, 88, 89, 90,
	91, 0, 0, 0, 0, 0, 81, 82, 83, 84,
	85, 79, 80, 71, 67, 68, 69, 0, 70, 0,
	202, 0, 0, 0, 0, 0, 61, 62, 0, 0,
	71, 67, 68, 69, 64, 70, 0, 200, 0, 0,
	0, 72, 0, 61, 62, 0, 65, 71, 67, 68,
	69, 64, 70, 0, 193, 0, 0, 0, 72, 0,
	61, 62, 0, 65, 71, 67, 68, 69, 64, 70,
	0, 150, 0, 0, 0, 72, 0, 61, 62, 0,
	65, 71, 67, 68, 69, 64, 70, 0, 112, 0,
	0, 0, 72, 0, 61, 62, 0, 65, 71, 67,
	68, 69, 64, 70, 0, 0, 0, 0, 0, 72,
	0, 61, 62, 0, 65, 0, 0, 0, 0, 64,
	0, 0, 0, 0, 0, 0, 72, 0, 0, 0,
	0, 65,
}
var yyPact = [...]int{

	174, -1000, -1000, 220, -1000, 146, 114, -1000, 283, 275,
	-1000, 232, 220, 145, 144, 29, -1000, -1000, 184, -1000,
	274, 274, 113, 232, -1000, -1000, 122, -1000, 13, 121,
	-1000, -1000, 170, 274, -1000, 10, -27, 273, 170, 109,
	172, -1000, 272, -30, 181, -6, 108, 464, 263, -1000,
	-1000, -1000, 262, 254, 464, 94, -1000, -11, -1000, -1000,
	6, -1000, 253, -33, 70, 70, 20, -1000, -1000, 252,
	70, -1000, 143, 36, -23, -1000, 86, 447, 242, 70,
	70, 70, 70, 70, 70, 70, 70, 70, 70, 70,
	70, 70, 70, 70, 251, 27, 242, -1000, -1000, 292,
	-1000, -1000, 70, 167, 70, 115, 271, 180, 70, 241,
	230, 430, -1000, -1000, -1000, 12, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 4,
	346, 4, -1000, 70, 210, -1000, -1000, 84, 346, -1000,
	13, 7, 70, 196, 70, 464, 18, 68, 138, -1000,
	-1000, 128, 194, 70, 59, 136, 70, -1000, 126, -1000,
	250, 189, 229, 82, -1000, -1000, 70, -1000, 131, 346,
	-1000, 70, 346, -1000, 464, 70, 70, 327, 58, 125,
	32, 75, 208, 187, -1000, 71, -1000, -1000, -1000, 413,
	464, 70, 464, -1000, 62, 96, 39, 396, 464, 379,
	-1000, 24, -1000, 133, -1000,
}
var yyPgo = [...]int{

	0, 5, 309, 308, 300, 319, 203, 190, 1, 299,
	0, 298, 297, 4, 296, 288, 275, 2, 193, 3,
	254, 6, 156, 7,
}
var yyR1 = [...]int{

	0, 2, 3, 4, 4, 14, 15, 15, 15, 5,
	5, 16, 16, 6, 6, 6, 6, 22, 22, 22,
	17, 17, 17, 18, 18, 7, 7, 1, 1, 1,
	1, 1, 1, 1, 1, 19, 19, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 20, 20, 21,
	21, 9, 9, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 13, 13, 13, 13, 13, 23,
	23, 23, 23, 11, 11, 11, 12, 12,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 1, 5, 1, 3, 2, 1,
	2, 1, 3, 0, 10, 10, 1, 0, 6, 8,
	0, 1, 3, 1, 3, 0, 2, 1, 3, 2,
	3, 4, 3, 4, 6, 1, 3, 0, 1, 1,
	1, 3, 3, 1, 5, 7, 1, 1, 3, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 3, 2, 2,
	4, 1, 1, 1, 3, 4, 4, 5, 1, 1,
	3, 3, 5, 8, 10, 12, 6, 10,
}
var yyChk = [...]int{

	-1000, -2, -3, -4, -14, 14, -16, -6, 13, 15,
	5, 37, 39, 4, 4, -15, -5, 5, 4, -6,
	37, 37, 38, 40, -5, 5, -18, -7, 4, -18,
	39, -5, 38, 40, -1, 4, 42, 32, 38, -22,
	16, -7, 41, 44, 43, 4, -22, 39, 15, 4,
	45, 4, 32, 41, 39, -19, -8, -10, -11, -12,
	-20, 17, 18, -9, 25, 37, -13, 5, 6, 7,
	9, 4, 32, 4, 4, 4, -19, 39, 47, 35,
	36, 30, 31, 32, 33, 34, 19, 20, 21, 22,
	23, 24, 46, 26, 40, 4, 47, -10, 4, -10,
	27, 28, 37, 41, 42, 4, -10, 37, 37, 41,
	41, 39, 11, -8, -23, 4, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, -10, -10, -10, -21,
	-10, -21, 4, 37, 41, -23, 38, -17, -10, 4,
	37, -10, 8, 40, 26, 39, -13, -17, 4, 4,
	11, 37, 41, 40, -17, 4, 40, 38, -1, 43,
	-10, 4, -10, -19, 38, 38, 37, 38, 4, -10,
	38, 37, -10, 38, 39, 8, 39, 39, -17, 37,
	-17, -19, -10, -10, 11, 10, 38, 38, 38, 39,
	39, 39, 39, 11, -19, -10, -19, 39, 39, 39,
	11, -19, 11, 39, 11,
}
var yyDef = [...]int{

	3, -2, 1, 13, 4, 0, 2, 11, 0, 0,
	16, 0, 13, 0, 0, 0, 6, 9, 0, 12,
	25, 25, 0, 0, 8, 10, 0, 23, 0, 0,
	5, 7, 17, 25, 26, 27, 0, 0, 17, 0,
	0, 24, 0, 0, 0, 29, 0, 37, 0, 30,
	32, 28, 0, 0, 37, 0, 35, 38, 39, 40,
	0, 43, 0, 46, 0, 0, 71, 72, 73, 0,
	0, -2, 0, 0, 33, 31, 0, 37, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 66, 78, 0,
	68, 69, 20, 0, 0, 0, 0, 0, 20, 0,
	0, 37, 14, 36, 51, 79, 53, 54, 55, 56,
	57, 58, 59, 60, 61, 62, 63, 64, 65, 41,
	49, 42, 48, 20, 0, 52, 67, 0, 21, 74,
	0, 0, 0, 0, 0, 37, 0, 0, 0, 34,
	15, 0, 0, 0, 0, 0, 0, 70, 0, 75,
	0, 0, 0, 0, 76, 18, 20, 80, 81, 50,
	44, 20, 22, 77, 37, 0, 0, 37, 0, 0,
	0, 0, 0, 0, 86, 0, 19, 82, 45, 37,
	37, 0, 37, 83, 0, 0, 0, 37, 37, 37,
	84, 0, 87, 37, 85,
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
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:43
		{
			yyVAL.astList.Add(yyDollar[2].iAstNode)
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:45
		{
			yyVAL.iAstNode = &astImport{"", yyDollar[1].string}
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:46
		{
			yyVAL.iAstNode = &astImport{yyDollar[1].string, yyDollar[2].string}
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:48
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:49
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 13:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:51
		{
			yyVAL.iAstNode = nil
		}
	case 14:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:53
		{
			yyVAL.iAstNode = &astTemplate{yyDollar[2].string, yyDollar[4].astList, yyDollar[6].astUseWrapper, yyDollar[8].astList}
		}
	case 15:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:55
		{
			yyVAL.iAstNode = &astWrapper{yyDollar[2].string, yyDollar[4].astList, yyDollar[6].astUseWrapper, yyDollar[8].astList}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:56
		{
			yyVAL.iAstNode = nil
		}
	case 17:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:58
		{
			yyVAL.astUseWrapper = nil
		}
	case 18:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:60
		{
			yyVAL.astUseWrapper = &astUseWrapper{"", yyDollar[3].string, yyDollar[5].astList}
		}
	case 19:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line template.y:62
		{
			yyVAL.astUseWrapper = &astUseWrapper{yyDollar[3].string, yyDollar[5].string, yyDollar[7].astList}
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:64
		{
			yyVAL.astList = nil
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:65
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:66
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:68
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:69
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 25:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:71
		{
			yyVAL.iAstNode = nil
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:72
		{
			yyVAL.iAstNode = &astVariableDef{yyDollar[1].string, yyDollar[2].string}
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:74
		{
			yyVAL.string = yyDollar[1].string
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:75
		{
			yyVAL.string = "[]" + yyDollar[3].string
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:76
		{
			yyVAL.string = "*" + yyDollar[2].string
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:77
		{
			yyVAL.string = yyDollar[1].string + "." + yyDollar[3].string
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:79
		{
			yyVAL.string = "*" + yyDollar[2].string + "." + yyDollar[4].string
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:80
		{
			yyVAL.string = yyDollar[1].string + "{}"
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:81
		{
			yyVAL.string = "[]*" + yyDollar[4].string
		}
	case 34:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:83
		{
			yyVAL.string = "[]*" + yyDollar[4].string + "." + yyDollar[6].string
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:85
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:86
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 37:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:88
		{
			yyVAL.iAstNode = nil
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:89
		{
			yyVAL.iAstNode = &astWriteValue{yyDollar[1].iAstNode}
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:90
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:91
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:93
		{
			yyVAL.iAstNode = &astAssignment{"=", yyDollar[1].astList, yyDollar[3].astList}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:95
		{
			yyVAL.iAstNode = &astAssignment{":=", yyDollar[1].astList, yyDollar[3].astList}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:96
		{
			yyVAL.iAstNode = &astWriteContent{}
		}
	case 44:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:98
		{
			yyVAL.iAstNode = &astProcessTemplate{"", yyDollar[2].string, yyDollar[4].astList}
		}
	case 45:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line template.y:100
		{
			yyVAL.iAstNode = &astProcessTemplate{yyDollar[2].string, yyDollar[4].string, yyDollar[6].astList}
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:101
		{
			yyVAL.iAstNode = &astWriteString{yyDollar[1].iAstNode}
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:103
		{
			yyVAL.astList = &astList{[]iAstNode{&astValue{yyDollar[1].string}}}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:105
		{
			yyVAL.astList.Add(&astValue{yyDollar[3].string})
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:107
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:108
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
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
		//line template.y:111
		{
			yyDollar[3].astFilter.value = yyDollar[1].iAstNode
			yyVAL.iAstNode = yyDollar[3].astFilter
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:114
		{
			yyVAL.iAstNode = &astExpr{">", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:115
		{
			yyVAL.iAstNode = &astExpr{"<", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:116
		{
			yyVAL.iAstNode = &astExpr{"+", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:117
		{
			yyVAL.iAstNode = &astExpr{"-", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:118
		{
			yyVAL.iAstNode = &astExpr{"*", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:119
		{
			yyVAL.iAstNode = &astExpr{"/", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:120
		{
			yyVAL.iAstNode = &astExpr{"%", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:121
		{
			yyVAL.iAstNode = &astExpr{"==", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:122
		{
			yyVAL.iAstNode = &astExpr{"!=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:123
		{
			yyVAL.iAstNode = &astExpr{">=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:124
		{
			yyVAL.iAstNode = &astExpr{"<=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:125
		{
			yyVAL.iAstNode = &astExpr{"||", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:126
		{
			yyVAL.iAstNode = &astExpr{"&&", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:127
		{
			yyVAL.iAstNode = &astExpr{"!", nil, yyDollar[2].iAstNode}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:128
		{
			yyVAL.iAstNode = &astParenthesis{yyDollar[2].iAstNode}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:129
		{
			yyVAL.iAstNode = &astExpr{"++", yyDollar[1].iAstNode, nil}
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:130
		{
			yyVAL.iAstNode = &astExpr{"--", yyDollar[1].iAstNode, nil}
		}
	case 70:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:132
		{
			yyVAL.iAstNode = &astFunc{yyDollar[1].iAstNode, yyDollar[3].astList}
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:133
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:134
		{
			yyVAL.iAstNode = &astString{yyDollar[1].string}
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:135
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:137
		{
			yyVAL.iAstNode = &astStructField{yyDollar[1].iAstNode, yyDollar[3].string}
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:138
		{
			yyVAL.iAstNode = &astIndexedVar{yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 76:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:139
		{
			yyVAL.iAstNode = &astRef{yyDollar[3].iAstNode}
		}
	case 77:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:142
		{
			yyVAL.iAstNode = &astTypeConv{yyDollar[1].iAstNode, yyDollar[4].string}
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:143
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:146
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:147
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:148
		{
			yyVAL.astFilter = &astFilter{pkg: yyDollar[1].string, name: yyDollar[3].string}
		}
	case 82:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:150
		{
			yyVAL.astFilter = &astFilter{pkg: yyDollar[1].string, name: yyDollar[3].string}
		}
	case 83:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line template.y:154
		{
			yyVAL.iAstNode = &astRangeLoop{"_", yyDollar[2].string, yyDollar[4].iAstNode, yyDollar[6].astList}
		}
	case 84:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:156
		{
			yyVAL.iAstNode = &astRangeLoop{yyDollar[2].string, yyDollar[4].string, yyDollar[6].iAstNode, yyDollar[8].astList}
		}
	case 85:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line template.y:158
		{
			yyVAL.iAstNode = &astLoop{&astExpr{":=", &astValue{yyDollar[2].string}, yyDollar[4].iAstNode}, yyDollar[6].iAstNode, yyDollar[8].iAstNode, yyDollar[10].astList}
		}
	case 86:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:160
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, nil}
		}
	case 87:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:162
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, yyDollar[8].astList}
		}
	}
	goto yystack /* stack new state and value */
}
