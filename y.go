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
	"'='",
	"'|'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line template.y:149

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 80
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 491

var yyAct = [...]int{

	53, 52, 107, 51, 94, 153, 146, 145, 90, 91,
	49, 123, 80, 81, 82, 83, 84, 85, 42, 41,
	96, 34, 190, 75, 76, 77, 78, 79, 73, 74,
	80, 81, 82, 83, 84, 85, 141, 186, 72, 184,
	142, 75, 76, 77, 78, 79, 73, 74, 179, 36,
	185, 176, 164, 129, 70, 87, 88, 130, 104, 35,
	102, 71, 95, 98, 103, 89, 56, 100, 50, 90,
	91, 56, 86, 106, 109, 110, 111, 112, 113, 114,
	115, 116, 117, 118, 119, 120, 121, 122, 45, 133,
	124, 29, 56, 128, 87, 88, 175, 131, 143, 97,
	63, 64, 12, 124, 89, 166, 106, 135, 90, 91,
	57, 63, 64, 65, 138, 66, 172, 171, 174, 160,
	61, 134, 137, 58, 59, 56, 173, 67, 143, 156,
	124, 61, 62, 38, 149, 159, 151, 143, 67, 155,
	152, 147, 101, 62, 158, 154, 144, 143, 143, 37,
	31, 32, 32, 22, 21, 23, 124, 56, 20, 11,
	26, 124, 97, 169, 170, 168, 106, 165, 92, 25,
	16, 44, 167, 39, 46, 5, 162, 24, 106, 182,
	157, 181, 56, 183, 150, 56, 106, 148, 106, 188,
	67, 28, 106, 40, 30, 7, 139, 56, 56, 108,
	56, 127, 126, 18, 17, 56, 56, 56, 19, 6,
	10, 56, 80, 81, 82, 83, 84, 85, 8, 125,
	9, 99, 93, 75, 76, 77, 78, 79, 73, 74,
	69, 68, 178, 80, 81, 82, 83, 84, 85, 15,
	48, 47, 43, 27, 75, 76, 77, 78, 79, 73,
	74, 14, 13, 177, 80, 81, 82, 83, 84, 85,
	4, 55, 54, 60, 3, 75, 76, 77, 78, 79,
	73, 74, 2, 1, 163, 80, 81, 82, 83, 84,
	85, 33, 0, 0, 0, 0, 75, 76, 77, 78,
	79, 73, 74, 0, 0, 161, 80, 81, 82, 83,
	84, 85, 0, 0, 0, 0, 0, 75, 76, 77,
	78, 79, 73, 74, 0, 0, 136, 80, 81, 82,
	83, 84, 85, 0, 0, 0, 0, 0, 75, 76,
	77, 78, 79, 73, 74, 0, 132, 80, 81, 82,
	83, 84, 85, 0, 0, 0, 0, 0, 75, 76,
	77, 78, 79, 73, 74, 57, 63, 64, 65, 0,
	66, 0, 191, 0, 0, 0, 0, 0, 58, 59,
	0, 0, 57, 63, 64, 65, 61, 66, 0, 189,
	0, 0, 0, 67, 0, 58, 59, 0, 62, 57,
	63, 64, 65, 61, 66, 0, 187, 0, 0, 0,
	67, 0, 58, 59, 0, 62, 57, 63, 64, 65,
	61, 66, 0, 180, 0, 0, 0, 67, 0, 58,
	59, 0, 62, 57, 63, 64, 65, 61, 66, 0,
	140, 0, 0, 0, 67, 0, 58, 59, 0, 62,
	57, 63, 64, 65, 61, 66, 0, 105, 0, 0,
	0, 67, 0, 58, 59, 0, 62, 57, 63, 64,
	65, 61, 66, 0, 0, 0, 0, 0, 67, 0,
	58, 59, 0, 62, 0, 0, 0, 0, 61, 0,
	0, 0, 0, 0, 0, 67, 0, 0, 0, 0,
	62,
}
var yyPact = [...]int{

	161, -1000, -1000, 205, -1000, 122, 63, -1000, 248, 247,
	-1000, 199, 205, 121, 117, 115, -1000, -1000, 172, -1000,
	239, 239, 52, 199, -1000, 112, -1000, 17, 111, -1000,
	-1000, 157, 239, -1000, -22, -25, 238, 157, 49, 159,
	-1000, 237, 236, -31, 29, 453, 227, -1000, -1000, 226,
	453, 22, -1000, -7, -1000, -1000, 28, 142, -1000, 218,
	-41, 95, 95, -1000, -1000, 217, 95, 105, 23, -1000,
	19, 436, 195, 95, 95, 95, 95, 95, 95, 95,
	95, 95, 95, 95, 95, 95, 95, -1000, -1000, 95,
	215, 196, 95, 16, 195, -1000, 67, -1000, 298, 81,
	277, 158, 95, 192, 419, -1000, -1000, -1000, -1, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 318, 108, 318, -1000, -36, -37, 318, 95,
	183, -1000, -1000, 95, 180, 95, 453, -33, 107, 102,
	-1000, 91, 176, 95, -1000, -1000, -1000, 97, 82, 256,
	168, 235, 13, -1000, -1000, 95, -1000, 68, 318, -1000,
	95, 453, 95, 95, 106, 88, 80, 58, 12, 214,
	193, -1000, 9, -1000, -1000, -1000, 402, 453, 95, 453,
	-1000, 0, 11, -2, 385, 453, 368, -1000, -17, -1000,
	351, -1000,
}
var yyPgo = [...]int{

	0, 281, 20, 273, 272, 264, 170, 195, 160, 1,
	263, 0, 262, 261, 260, 239, 209, 11, 169, 3,
	133, 2,
}
var yyR1 = [...]int{

	0, 3, 4, 5, 5, 14, 15, 15, 6, 6,
	16, 16, 7, 7, 7, 7, 20, 20, 20, 17,
	17, 17, 18, 18, 8, 8, 1, 1, 1, 1,
	1, 19, 19, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 10, 10, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 2, 2, 2, 2,
	2, 21, 21, 21, 21, 12, 12, 12, 13, 13,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 1, 5, 1, 3, 1, 2,
	1, 3, 0, 10, 10, 1, 0, 6, 8, 0,
	1, 3, 1, 3, 0, 2, 1, 3, 2, 3,
	4, 1, 3, 0, 1, 1, 1, 3, 3, 1,
	5, 7, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 3,
	2, 2, 4, 1, 1, 1, 3, 4, 4, 4,
	1, 1, 3, 3, 5, 8, 10, 12, 6, 10,
}
var yyChk = [...]int{

	-1000, -3, -4, -5, -14, 14, -16, -7, 13, 15,
	5, 37, 39, 4, 4, -15, -6, 5, 4, -7,
	37, 37, 38, 40, 5, -18, -8, 4, -18, 39,
	-6, 38, 40, -1, 4, 42, 32, 38, -20, 16,
	-8, 41, 43, 4, -20, 39, 15, 4, 4, 41,
	39, -19, -9, -11, -12, -13, -2, 4, 17, 18,
	-10, 25, 37, 5, 6, 7, 9, 32, 4, 4,
	-19, 39, 45, 35, 36, 30, 31, 32, 33, 34,
	19, 20, 21, 22, 23, 24, 44, 27, 28, 37,
	41, 42, 26, 4, 45, -11, -2, 4, -11, 4,
	-11, 37, 37, 41, 39, 11, -9, -21, 4, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -17, -11, 4, 6, 5, -11, 37,
	41, -21, 38, 8, 40, 26, 39, -2, -17, 4,
	11, 37, 41, 40, 38, 43, 43, -17, 4, -11,
	4, -11, -19, 38, 38, 37, 38, 4, -11, 38,
	37, 39, 8, 39, 39, -17, 37, -17, -19, -11,
	-11, 11, 10, 38, 38, 38, 39, 39, 39, 39,
	11, -19, -11, -19, 39, 39, 39, 11, -19, 11,
	39, 11,
}
var yyDef = [...]int{

	3, -2, 1, 12, 4, 0, 2, 10, 0, 0,
	15, 0, 12, 0, 0, 0, 6, 8, 0, 11,
	24, 24, 0, 0, 9, 0, 22, 0, 0, 5,
	7, 16, 24, 25, 26, 0, 0, 16, 0, 0,
	23, 0, 0, 28, 0, 33, 0, 29, 27, 0,
	33, 0, 31, 34, 35, 36, 63, 70, 39, 0,
	42, 0, 0, 64, 65, 0, 0, 0, 0, 30,
	0, 33, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 60, 61, 19,
	0, 0, 0, 0, 0, 58, 63, 70, 0, 0,
	0, 0, 19, 0, 33, 13, 32, 43, 71, 45,
	46, 47, 48, 49, 50, 51, 52, 53, 54, 55,
	56, 57, 37, 0, 20, 66, 0, 0, 38, 19,
	0, 44, 59, 0, 0, 0, 33, 0, 0, 0,
	14, 0, 0, 0, 62, 67, 68, 0, 0, 0,
	0, 0, 0, 69, 17, 19, 72, 73, 21, 40,
	19, 33, 0, 0, 33, 0, 0, 0, 0, 0,
	0, 78, 0, 18, 74, 41, 33, 33, 0, 33,
	75, 0, 0, 0, 33, 33, 33, 76, 0, 79,
	33, 77,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 34, 3, 3,
	37, 38, 32, 30, 40, 31, 41, 33, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 39,
	36, 44, 35, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 42, 3, 43, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 45,
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
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:80
		{
			yyVAL.astList = &astList{[]iAstNode{yyDollar[1].iAstNode}}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:81
		{
			yyVAL.astList.Add(yyDollar[3].iAstNode)
		}
	case 33:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line template.y:83
		{
			yyVAL.iAstNode = nil
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:84
		{
			yyVAL.iAstNode = &astWriteValue{yyDollar[1].iAstNode}
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:85
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:86
		{
			yyVAL.iAstNode = yyDollar[1].iAstNode
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:87
		{
			yyVAL.iAstNode = &astAssignment{"=", &astValue{yyDollar[1].string}, yyDollar[3].iAstNode}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:88
		{
			yyVAL.iAstNode = &astAssignment{":=", &astValue{yyDollar[1].string}, yyDollar[3].iAstNode}
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:89
		{
			yyVAL.iAstNode = &astWriteContent{}
		}
	case 40:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:91
		{
			yyVAL.iAstNode = &astProcessTemplate{"", yyDollar[2].string, yyDollar[4].astList}
		}
	case 41:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line template.y:93
		{
			yyVAL.iAstNode = &astProcessTemplate{yyDollar[2].string, yyDollar[4].string, yyDollar[6].astList}
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:94
		{
			yyVAL.iAstNode = &astWriteString{yyDollar[1].iAstNode}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:96
		{
			yyDollar[3].astFilter.value = yyDollar[1].iAstNode
			yyVAL.iAstNode = yyDollar[3].astFilter
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:97
		{
			yyDollar[3].astFilter.value = yyDollar[1].iAstNode
			yyVAL.iAstNode = yyDollar[3].astFilter
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:100
		{
			yyVAL.iAstNode = &astExpr{">", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:101
		{
			yyVAL.iAstNode = &astExpr{"<", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:102
		{
			yyVAL.iAstNode = &astExpr{"+", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:103
		{
			yyVAL.iAstNode = &astExpr{"-", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:104
		{
			yyVAL.iAstNode = &astExpr{"*", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:105
		{
			yyVAL.iAstNode = &astExpr{"/", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:106
		{
			yyVAL.iAstNode = &astExpr{"%", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:107
		{
			yyVAL.iAstNode = &astExpr{"==", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:108
		{
			yyVAL.iAstNode = &astExpr{"!=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:109
		{
			yyVAL.iAstNode = &astExpr{">=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:110
		{
			yyVAL.iAstNode = &astExpr{"<=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:111
		{
			yyVAL.iAstNode = &astExpr{"||", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:112
		{
			yyVAL.iAstNode = &astExpr{"&&", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:113
		{
			yyVAL.iAstNode = &astExpr{"!", nil, yyDollar[2].iAstNode}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:114
		{
			yyVAL.iAstNode = &astParenthesis{yyDollar[2].iAstNode}
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:115
		{
			yyVAL.iAstNode = &astExpr{"++", &astValue{yyDollar[1].string}, nil}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:116
		{
			yyVAL.iAstNode = &astExpr{"--", &astValue{yyDollar[1].string}, nil}
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:118
		{
			yyVAL.iAstNode = &astFunc{yyDollar[1].string, yyDollar[3].astList}
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:119
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:120
		{
			yyVAL.iAstNode = &astString{yyDollar[1].string}
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:121
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:123
		{
			yyVAL.string = yyDollar[1].string + "." + yyDollar[3].string
		}
	case 67:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:124
		{
			yyVAL.string = yyDollar[1].string + "[" + yyDollar[3].string + "]"
		}
	case 68:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:125
		{
			yyVAL.string = yyDollar[1].string + "[" + yyDollar[3].string + "]"
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:126
		{
			yyVAL.string = "*(" + yyDollar[3].string + ")"
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:128
		{
			yyVAL.string = yyDollar[1].string
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:131
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:132
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:133
		{
			yyVAL.astFilter = &astFilter{pkg: yyDollar[1].string, name: yyDollar[3].string}
		}
	case 74:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line template.y:135
		{
			yyVAL.astFilter = &astFilter{pkg: yyDollar[1].string, name: yyDollar[3].string}
		}
	case 75:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line template.y:139
		{
			yyVAL.iAstNode = &astRangeLoop{"_", yyDollar[2].string, yyDollar[4].iAstNode, yyDollar[6].astList}
		}
	case 76:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:141
		{
			yyVAL.iAstNode = &astRangeLoop{yyDollar[2].string, yyDollar[4].string, yyDollar[6].iAstNode, yyDollar[8].astList}
		}
	case 77:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line template.y:143
		{
			yyVAL.iAstNode = &astLoop{&astExpr{":=", &astValue{yyDollar[2].string}, yyDollar[4].iAstNode}, yyDollar[6].iAstNode, yyDollar[8].iAstNode, yyDollar[10].astList}
		}
	case 78:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:145
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, nil}
		}
	case 79:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:147
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, yyDollar[8].astList}
		}
	}
	goto yystack /* stack new state and value */
}
