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

//line template.y:154

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 74
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 519

var yyAct = [...]int{

	49, 117, 82, 83, 145, 47, 31, 85, 86, 46,
	138, 84, 137, 39, 126, 85, 86, 90, 81, 123,
	38, 148, 48, 124, 75, 76, 77, 78, 79, 80,
	179, 175, 128, 33, 70, 71, 72, 73, 74, 68,
	69, 82, 83, 32, 164, 127, 135, 97, 63, 67,
	84, 98, 151, 173, 85, 86, 168, 89, 92, 52,
	52, 163, 94, 135, 150, 146, 135, 135, 165, 103,
	104, 105, 106, 107, 108, 109, 110, 111, 112, 113,
	114, 115, 116, 155, 52, 118, 96, 66, 122, 100,
	75, 76, 77, 78, 79, 80, 42, 41, 118, 132,
	70, 71, 72, 73, 74, 68, 69, 26, 12, 174,
	91, 58, 59, 130, 52, 136, 34, 135, 29, 100,
	28, 20, 29, 21, 118, 139, 147, 141, 134, 143,
	95, 56, 19, 18, 11, 144, 149, 62, 91, 23,
	87, 22, 57, 36, 43, 5, 7, 52, 118, 156,
	153, 27, 118, 157, 159, 160, 10, 16, 158, 17,
	142, 25, 121, 120, 8, 62, 9, 140, 171, 37,
	52, 133, 170, 52, 172, 119, 102, 93, 100, 88,
	177, 65, 64, 52, 52, 45, 52, 44, 100, 40,
	24, 52, 52, 52, 14, 13, 100, 52, 100, 101,
	35, 6, 100, 75, 76, 77, 78, 79, 80, 15,
	4, 51, 50, 70, 71, 72, 73, 74, 68, 69,
	3, 2, 167, 75, 76, 77, 78, 79, 80, 1,
	30, 0, 0, 70, 71, 72, 73, 74, 68, 69,
	0, 0, 166, 75, 76, 77, 78, 79, 80, 0,
	0, 0, 0, 70, 71, 72, 73, 74, 68, 69,
	0, 0, 154, 75, 76, 77, 78, 79, 80, 0,
	0, 0, 0, 70, 71, 72, 73, 74, 68, 69,
	0, 0, 152, 75, 76, 77, 78, 79, 80, 0,
	0, 0, 0, 70, 71, 72, 73, 74, 68, 69,
	0, 0, 129, 75, 76, 77, 78, 79, 80, 0,
	0, 0, 0, 70, 71, 72, 73, 74, 68, 69,
	0, 125, 53, 58, 59, 60, 0, 61, 162, 161,
	0, 0, 0, 0, 0, 54, 55, 0, 0, 0,
	0, 0, 0, 56, 0, 0, 0, 0, 0, 62,
	0, 0, 0, 0, 57, 75, 76, 77, 78, 79,
	80, 0, 0, 0, 0, 70, 71, 72, 73, 74,
	68, 69, 53, 58, 59, 60, 0, 61, 0, 180,
	0, 0, 0, 0, 0, 54, 55, 53, 58, 59,
	60, 0, 61, 56, 178, 0, 0, 0, 0, 62,
	54, 55, 0, 0, 57, 0, 0, 0, 56, 0,
	53, 58, 59, 60, 62, 61, 0, 176, 0, 57,
	0, 0, 0, 54, 55, 53, 58, 59, 60, 0,
	61, 56, 169, 0, 0, 0, 0, 62, 54, 55,
	0, 0, 57, 0, 0, 0, 56, 0, 53, 58,
	59, 60, 62, 61, 0, 131, 0, 57, 0, 0,
	0, 54, 55, 53, 58, 59, 60, 0, 61, 56,
	99, 0, 0, 0, 0, 62, 54, 55, 0, 0,
	57, 0, 0, 0, 56, 0, 53, 58, 59, 60,
	62, 61, 0, 0, 0, 57, 0, 0, 0, 54,
	55, 0, 0, 0, 0, 0, 0, 56, 0, 0,
	0, 0, 0, 62, 0, 0, 0, 0, 57,
}
var yyPact = [...]int{

	131, -1000, -1000, 151, -1000, 98, 70, -1000, 191, 190,
	-1000, 152, 151, 97, 96, 84, -1000, -1000, 186, 186,
	69, 146, 83, -1000, 2, 79, -1000, -1000, 127, 186,
	-1000, -20, -29, 185, 59, 58, 129, -1000, 183, 181,
	-31, 482, 482, 178, -1000, -1000, 177, 49, -1000, 5,
	-1000, -1000, -25, 114, -1000, 175, 106, 106, -1000, -1000,
	173, 106, 94, 48, 11, -1000, 459, 172, 106, 106,
	106, 106, 106, 106, 106, 106, 106, 106, 106, 106,
	106, 106, -1000, -1000, 106, 171, 157, 106, -17, -1000,
	14, -1000, 284, 6, 264, 134, 444, 106, 167, -1000,
	-1000, -1000, 92, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 336, 78, 336, -1000,
	-30, -32, 336, 106, 163, -1000, 106, 156, 106, 482,
	-33, -1000, 28, 90, -16, 106, -1000, -1000, -1000, 27,
	16, 244, 142, 224, 45, -1000, -1000, 106, -1000, 336,
	-1000, 106, 482, 106, 106, 318, 24, 7, 30, 204,
	184, -1000, 18, -1000, -1000, 421, 482, 106, 482, -1000,
	15, 71, -7, 406, 482, 383, -1000, -8, -1000, 368,
	-1000,
}
var yyPgo = [...]int{

	0, 230, 17, 229, 221, 220, 146, 139, 22, 0,
	212, 211, 210, 209, 201, 1, 141, 5, 200, 199,
}
var yyR1 = [...]int{

	0, 3, 4, 5, 5, 12, 13, 13, 14, 14,
	6, 6, 6, 6, 18, 18, 18, 15, 15, 15,
	16, 16, 7, 7, 1, 1, 1, 1, 1, 17,
	17, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 2, 2, 2, 2, 2, 19, 19, 10,
	10, 10, 11, 11,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 1, 5, 1, 3, 1, 3,
	0, 10, 9, 1, 0, 6, 8, 0, 1, 3,
	1, 3, 0, 2, 1, 3, 2, 3, 4, 1,
	3, 0, 1, 1, 1, 3, 3, 1, 5, 7,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 2, 3, 2, 2, 4, 1,
	1, 1, 3, 4, 4, 4, 1, 1, 3, 8,
	10, 12, 6, 10,
}
var yyChk = [...]int{

	-1000, -3, -4, -5, -12, 14, -14, -6, 13, 15,
	5, 36, 38, 4, 4, -13, 5, -6, 36, 36,
	37, 39, -16, -7, 4, -16, 38, 5, 37, 39,
	-1, 4, 41, 31, 37, -18, 16, -7, 40, 42,
	4, 38, 38, 15, 4, 4, 40, -17, -8, -9,
	-10, -11, -2, 4, 17, 18, 25, 36, 5, 6,
	7, 9, 31, -17, 4, 4, 38, 44, 34, 35,
	29, 30, 31, 32, 33, 19, 20, 21, 22, 23,
	24, 43, 27, 28, 36, 40, 41, 26, 4, -9,
	-2, 4, -9, 4, -9, 36, 38, 36, 40, 11,
	-8, -19, 4, -9, -9, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -15, -9, 4,
	6, 5, -9, 36, 40, 37, 8, 39, 26, 38,
	-2, 11, -15, 4, 36, 39, 37, 42, 42, -15,
	4, -9, 4, -9, -17, 37, 37, 36, 37, -9,
	37, 36, 38, 8, 38, 38, -15, -15, -17, -9,
	-9, 11, 10, 37, 37, 38, 38, 38, 38, 11,
	-17, -9, -17, 38, 38, 38, 11, -17, 11, 38,
	11,
}
var yyDef = [...]int{

	3, -2, 1, 10, 4, 0, 2, 8, 0, 0,
	13, 0, 10, 0, 0, 0, 6, 9, 22, 22,
	0, 0, 0, 20, 0, 0, 5, 7, 14, 22,
	23, 24, 0, 0, 0, 0, 0, 21, 0, 0,
	26, 31, 31, 0, 27, 25, 0, 0, 29, 32,
	33, 34, 59, 66, 37, 0, 0, 0, 60, 61,
	0, 0, 0, 0, 0, 28, 31, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 56, 57, 17, 0, 0, 0, 0, 54,
	59, 66, 0, 0, 0, 0, 31, 17, 0, 12,
	30, 40, 67, 41, 42, 43, 44, 45, 46, 47,
	48, 49, 50, 51, 52, 53, 35, 0, 18, 62,
	0, 0, 36, 17, 0, 55, 0, 0, 0, 31,
	0, 11, 0, 0, 0, 0, 58, 63, 64, 0,
	0, 0, 0, 0, 0, 65, 15, 17, 68, 19,
	38, 17, 31, 0, 0, 31, 0, 0, 0, 0,
	0, 72, 0, 16, 39, 31, 31, 0, 31, 69,
	0, 0, 0, 31, 31, 31, 70, 0, 73, 31,
	71,
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
			yyVAL.iAstNode = &astProcessTemplate{"", yyDollar[2].string, yyDollar[4].astList}
		}
	case 39:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line template.y:104
		{
			yyVAL.iAstNode = &astProcessTemplate{yyDollar[2].string, yyDollar[4].string, yyDollar[6].astList}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:105
		{
			yyDollar[3].astFilter.value = yyDollar[1].iAstNode
			yyVAL.iAstNode = &astWriteString{yyDollar[3].astFilter}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:108
		{
			yyVAL.iAstNode = &astExpr{">", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:109
		{
			yyVAL.iAstNode = &astExpr{"<", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:110
		{
			yyVAL.iAstNode = &astExpr{"+", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:111
		{
			yyVAL.iAstNode = &astExpr{"-", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:112
		{
			yyVAL.iAstNode = &astExpr{"*", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:113
		{
			yyVAL.iAstNode = &astExpr{"/", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:114
		{
			yyVAL.iAstNode = &astExpr{"%", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:115
		{
			yyVAL.iAstNode = &astExpr{"==", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:116
		{
			yyVAL.iAstNode = &astExpr{"!=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:117
		{
			yyVAL.iAstNode = &astExpr{">=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:118
		{
			yyVAL.iAstNode = &astExpr{"<=", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:119
		{
			yyVAL.iAstNode = &astExpr{"||", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:120
		{
			yyVAL.iAstNode = &astExpr{"&&", yyDollar[1].iAstNode, yyDollar[3].iAstNode}
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:121
		{
			yyVAL.iAstNode = &astExpr{"!", nil, yyDollar[2].iAstNode}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:122
		{
			yyVAL.iAstNode = &astParenthesis{yyDollar[2].iAstNode}
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:123
		{
			yyVAL.iAstNode = &astExpr{"++", &astValue{yyDollar[1].string}, nil}
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line template.y:124
		{
			yyVAL.iAstNode = &astExpr{"--", &astValue{yyDollar[1].string}, nil}
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:126
		{
			yyVAL.iAstNode = &astFunc{yyDollar[1].string, yyDollar[3].astList}
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:127
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:128
		{
			yyVAL.iAstNode = &astString{yyDollar[1].string}
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:129
		{
			yyVAL.iAstNode = &astValue{yyDollar[1].string}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:131
		{
			yyVAL.string = yyDollar[1].string + "." + yyDollar[3].string
		}
	case 63:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:132
		{
			yyVAL.string = yyDollar[1].string + "[" + yyDollar[3].string + "]"
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:133
		{
			yyVAL.string = yyDollar[1].string + "[" + yyDollar[3].string + "]"
		}
	case 65:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line template.y:134
		{
			yyVAL.string = "*(" + yyDollar[3].string + ")"
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:136
		{
			yyVAL.string = yyDollar[1].string
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line template.y:139
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line template.y:140
		{
			yyVAL.astFilter = &astFilter{name: yyDollar[1].string}
		}
	case 69:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line template.y:144
		{
			yyVAL.iAstNode = &astRangeLoop{"_", yyDollar[2].string, yyDollar[4].iAstNode, yyDollar[6].astList}
		}
	case 70:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:146
		{
			yyVAL.iAstNode = &astRangeLoop{yyDollar[2].string, yyDollar[4].string, yyDollar[6].iAstNode, yyDollar[8].astList}
		}
	case 71:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line template.y:148
		{
			yyVAL.iAstNode = &astLoop{&astExpr{":=", &astValue{yyDollar[2].string}, yyDollar[4].iAstNode}, yyDollar[6].iAstNode, yyDollar[8].iAstNode, yyDollar[10].astList}
		}
	case 72:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:150
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, nil}
		}
	case 73:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:152
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, yyDollar[8].astList}
		}
	}
	goto yystack /* stack new state and value */
}
