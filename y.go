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
	-1, 70,
	26, 46,
	40, 46,
	46, 46,
	-2, 78,
}

const yyNprod = 88
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 499

var yyAct = [...]int{

	56, 55, 136, 54, 33, 113, 65, 128, 85, 86,
	87, 88, 89, 90, 95, 49, 160, 34, 159, 80,
	81, 82, 83, 84, 78, 79, 85, 86, 87, 88,
	89, 90, 41, 43, 92, 42, 77, 80, 81, 82,
	83, 84, 78, 79, 109, 36, 199, 165, 93, 204,
	102, 103, 99, 100, 91, 35, 151, 75, 52, 132,
	152, 153, 101, 133, 96, 98, 102, 103, 107, 200,
	105, 189, 108, 156, 187, 188, 156, 142, 112, 115,
	116, 117, 118, 119, 120, 121, 122, 123, 124, 125,
	126, 127, 129, 129, 171, 144, 156, 166, 174, 156,
	130, 134, 137, 157, 37, 156, 32, 198, 137, 143,
	147, 193, 112, 146, 85, 86, 87, 88, 89, 90,
	97, 66, 67, 190, 178, 80, 81, 82, 83, 84,
	78, 79, 110, 137, 192, 154, 31, 22, 32, 23,
	76, 63, 53, 161, 158, 163, 46, 29, 71, 164,
	12, 168, 138, 64, 170, 180, 38, 173, 172, 167,
	106, 21, 20, 11, 16, 26, 39, 25, 137, 47,
	179, 97, 5, 137, 10, 181, 50, 183, 184, 182,
	112, 176, 8, 7, 9, 139, 141, 140, 30, 28,
	18, 17, 112, 196, 45, 195, 19, 197, 40, 71,
	112, 24, 112, 202, 51, 169, 112, 85, 86, 87,
	88, 89, 90, 59, 162, 155, 149, 148, 80, 81,
	82, 83, 84, 78, 79, 114, 131, 191, 85, 86,
	87, 88, 89, 90, 6, 104, 94, 74, 73, 80,
	81, 82, 83, 84, 78, 79, 72, 48, 177, 85,
	86, 87, 88, 89, 90, 15, 44, 27, 14, 13,
	80, 81, 82, 83, 84, 78, 79, 4, 58, 175,
	85, 86, 87, 88, 89, 90, 57, 62, 3, 2,
	1, 80, 81, 82, 83, 84, 78, 79, 0, 0,
	145, 85, 86, 87, 88, 89, 90, 0, 0, 0,
	0, 0, 80, 81, 82, 83, 84, 78, 79, 0,
	135, 70, 66, 67, 68, 0, 69, 186, 185, 0,
	0, 0, 0, 0, 60, 61, 0, 0, 0, 0,
	0, 0, 63, 0, 0, 0, 0, 0, 0, 71,
	0, 0, 0, 0, 64, 85, 86, 87, 88, 89,
	90, 0, 0, 0, 0, 0, 80, 81, 82, 83,
	84, 78, 79, 70, 66, 67, 68, 0, 69, 0,
	205, 0, 0, 0, 0, 0, 60, 61, 0, 0,
	70, 66, 67, 68, 63, 69, 0, 203, 0, 0,
	0, 71, 0, 60, 61, 0, 64, 70, 66, 67,
	68, 63, 69, 0, 201, 0, 0, 0, 71, 0,
	60, 61, 0, 64, 70, 66, 67, 68, 63, 69,
	0, 194, 0, 0, 0, 71, 0, 60, 61, 0,
	64, 70, 66, 67, 68, 63, 69, 0, 150, 0,
	0, 0, 71, 0, 60, 61, 0, 64, 70, 66,
	67, 68, 63, 69, 0, 111, 0, 0, 0, 71,
	0, 60, 61, 0, 64, 70, 66, 67, 68, 63,
	69, 0, 0, 0, 0, 0, 71, 0, 60, 61,
	0, 64, 0, 0, 0, 0, 63, 0, 0, 0,
	0, 0, 0, 71, 0, 0, 0, 0, 64,
}
var yyPact = [...]int{

	158, -1000, -1000, 169, -1000, 126, 111, -1000, 255, 254,
	-1000, 186, 169, 125, 124, 99, -1000, -1000, 196, -1000,
	253, 253, 108, 186, -1000, 98, -1000, 13, 66, -1000,
	-1000, 150, 253, -1000, -9, -10, 252, 150, 107, 154,
	-1000, 243, -30, 172, 17, 103, 461, 242, -1000, -1000,
	-1000, 234, 233, 461, 101, -1000, -11, -1000, -1000, 8,
	-1000, 232, -33, 116, 116, 25, -1000, -1000, 231, 116,
	-1000, 123, 31, 3, -1000, 93, 444, 221, 116, 116,
	116, 116, 116, 116, 116, 116, 116, 116, 116, 116,
	116, 116, 116, 222, 22, 221, -1000, -1000, 272, -1000,
	-1000, 116, 148, 181, 69, 251, 167, 116, 213, 212,
	427, -1000, -1000, -1000, 19, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 21, 326,
	21, -1000, 116, 211, -1000, -1000, 65, 326, -1000, 13,
	-25, -27, 116, 210, 116, 461, 9, 59, 122, -1000,
	-1000, 113, 201, 116, 56, 121, 116, -1000, 60, -1000,
	-1000, 230, 173, 209, 85, -1000, -1000, 116, -1000, 118,
	326, -1000, 116, 326, -1000, 461, 116, 116, 307, 36,
	37, 33, 84, 188, 95, -1000, 72, -1000, -1000, -1000,
	410, 461, 116, 461, -1000, 68, 7, 30, 393, 461,
	376, -1000, 10, -1000, 359, -1000,
}
var yyPgo = [...]int{

	0, 4, 6, 280, 279, 278, 164, 183, 165, 1,
	277, 0, 276, 268, 267, 255, 234, 2, 167, 3,
	213, 7, 156, 5,
}
var yyR1 = [...]int{

	0, 3, 4, 5, 5, 14, 15, 15, 6, 6,
	16, 16, 7, 7, 7, 7, 22, 22, 22, 17,
	17, 17, 18, 18, 8, 8, 1, 1, 1, 1,
	1, 1, 1, 1, 19, 19, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 20, 20, 21, 21,
	10, 10, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 2, 2, 2, 2, 2, 2, 23,
	23, 23, 23, 12, 12, 12, 13, 13,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 1, 5, 1, 3, 1, 2,
	1, 3, 0, 10, 10, 1, 0, 6, 8, 0,
	1, 3, 1, 3, 0, 2, 1, 3, 2, 3,
	4, 3, 4, 6, 1, 3, 0, 1, 1, 1,
	3, 3, 1, 5, 7, 1, 1, 3, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 3, 2, 2, 4,
	1, 1, 1, 3, 4, 4, 4, 5, 1, 1,
	3, 3, 5, 8, 10, 12, 6, 10,
}
var yyChk = [...]int{

	-1000, -3, -4, -5, -14, 14, -16, -7, 13, 15,
	5, 37, 39, 4, 4, -15, -6, 5, 4, -7,
	37, 37, 38, 40, 5, -18, -8, 4, -18, 39,
	-6, 38, 40, -1, 4, 42, 32, 38, -22, 16,
	-8, 41, 44, 43, 4, -22, 39, 15, 4, 45,
	4, 32, 41, 39, -19, -9, -11, -12, -13, -20,
	17, 18, -10, 25, 37, -2, 5, 6, 7, 9,
	4, 32, 4, 4, 4, -19, 39, 47, 35, 36,
	30, 31, 32, 33, 34, 19, 20, 21, 22, 23,
	24, 46, 26, 40, 4, 47, -11, 4, -11, 27,
	28, 37, 41, 42, 4, -11, 37, 37, 41, 41,
	39, 11, -9, -23, 4, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -21, -11,
	-21, 4, 37, 41, -23, 38, -17, -11, 4, 37,
	6, 5, 8, 40, 26, 39, -2, -17, 4, 4,
	11, 37, 41, 40, -17, 4, 40, 38, -1, 43,
	43, -11, 4, -11, -19, 38, 38, 37, 38, 4,
	-11, 38, 37, -11, 38, 39, 8, 39, 39, -17,
	37, -17, -19, -11, -11, 11, 10, 38, 38, 38,
	39, 39, 39, 39, 11, -19, -11, -19, 39, 39,
	39, 11, -19, 11, 39, 11,
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
	0, 0, 0, 0, 0, 0, 65, 78, 0, 67,
	68, 19, 0, 0, 0, 0, 0, 19, 0, 0,
	36, 13, 35, 50, 79, 52, 53, 54, 55, 56,
	57, 58, 59, 60, 61, 62, 63, 64, 40, 48,
	41, 47, 19, 0, 51, 66, 0, 20, 73, 0,
	0, 0, 0, 0, 0, 36, 0, 0, 0, 33,
	14, 0, 0, 0, 0, 0, 0, 69, 0, 74,
	75, 0, 0, 0, 0, 76, 17, 19, 80, 81,
	49, 43, 19, 21, 77, 36, 0, 0, 36, 0,
	0, 0, 0, 0, 0, 86, 0, 18, 82, 44,
	36, 36, 0, 36, 83, 0, 0, 0, 36, 36,
	36, 84, 0, 87, 36, 85,
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
			yyVAL.iAstNode = &astExpr{"++", &astValue{yyDollar[1].string}, nil}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:129
		{
			yyVAL.iAstNode = &astExpr{"--", &astValue{yyDollar[1].string}, nil}
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:131
		{
			yyVAL.iAstNode = &astFunc{yyDollar[1].string, yyDollar[3].astList}
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:132
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
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
			yyVAL.string = yyDollar[1].string + "." + yyDollar[3].string
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:137
		{
			yyVAL.string = yyDollar[1].string + "[" + yyDollar[3].string + "]"
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:138
		{
			yyVAL.string = yyDollar[1].string + "[" + yyDollar[3].string + "]"
		}
	case 76:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:139
		{
			yyVAL.string = "*(" + yyDollar[3].string + ")"
		}
	case 77:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:142
		{
			yyVAL.string = yyDollar[1].string + ".(" + yyDollar[4].string + ")"
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:143
		{
			yyVAL.string = yyDollar[1].string
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
