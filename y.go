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

//line template.y:143

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 76
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 496

var yyAct = [...]int{

	49, 90, 103, 148, 47, 31, 86, 87, 119, 83,
	84, 141, 140, 39, 46, 129, 125, 99, 85, 92,
	126, 100, 86, 87, 48, 76, 77, 78, 79, 80,
	81, 38, 33, 131, 182, 71, 72, 73, 74, 75,
	69, 70, 32, 178, 88, 151, 130, 64, 176, 171,
	68, 83, 84, 167, 166, 138, 138, 168, 91, 94,
	85, 52, 52, 96, 86, 87, 153, 82, 138, 158,
	105, 106, 107, 108, 109, 110, 111, 112, 113, 114,
	115, 116, 117, 118, 98, 67, 120, 52, 149, 124,
	138, 42, 102, 127, 139, 34, 138, 29, 41, 26,
	120, 53, 59, 60, 61, 28, 62, 29, 135, 20,
	36, 21, 12, 154, 54, 55, 150, 133, 52, 137,
	97, 19, 57, 102, 18, 11, 120, 43, 63, 93,
	144, 5, 146, 58, 142, 23, 7, 147, 156, 152,
	53, 59, 60, 61, 22, 62, 165, 164, 27, 17,
	16, 120, 52, 54, 55, 120, 63, 162, 163, 159,
	161, 57, 145, 160, 25, 37, 143, 63, 35, 10,
	136, 174, 58, 104, 173, 52, 175, 8, 52, 9,
	123, 122, 180, 102, 121, 95, 89, 66, 52, 52,
	65, 52, 45, 102, 44, 40, 52, 52, 52, 24,
	14, 102, 52, 102, 13, 6, 15, 102, 76, 77,
	78, 79, 80, 81, 4, 51, 50, 56, 71, 72,
	73, 74, 75, 69, 70, 3, 2, 177, 76, 77,
	78, 79, 80, 81, 1, 30, 0, 0, 71, 72,
	73, 74, 75, 69, 70, 0, 0, 170, 76, 77,
	78, 79, 80, 81, 0, 0, 0, 0, 71, 72,
	73, 74, 75, 69, 70, 0, 0, 169, 76, 77,
	78, 79, 80, 81, 0, 0, 0, 0, 71, 72,
	73, 74, 75, 69, 70, 0, 0, 157, 76, 77,
	78, 79, 80, 81, 0, 0, 0, 0, 71, 72,
	73, 74, 75, 69, 70, 0, 0, 155, 76, 77,
	78, 79, 80, 81, 0, 0, 0, 0, 71, 72,
	73, 74, 75, 69, 70, 0, 0, 132, 76, 77,
	78, 79, 80, 81, 0, 0, 0, 0, 71, 72,
	73, 74, 75, 69, 70, 0, 128, 76, 77, 78,
	79, 80, 81, 0, 0, 0, 0, 71, 72, 73,
	74, 75, 69, 70, 53, 59, 60, 61, 0, 62,
	0, 183, 0, 0, 0, 0, 0, 54, 55, 53,
	59, 60, 61, 0, 62, 57, 181, 0, 0, 0,
	0, 63, 54, 55, 0, 0, 58, 0, 0, 0,
	57, 0, 53, 59, 60, 61, 63, 62, 0, 179,
	0, 58, 0, 0, 0, 54, 55, 53, 59, 60,
	61, 0, 62, 57, 172, 0, 0, 0, 0, 63,
	54, 55, 0, 0, 58, 0, 0, 0, 57, 0,
	53, 59, 60, 61, 63, 62, 0, 134, 0, 58,
	0, 0, 0, 54, 55, 53, 59, 60, 61, 0,
	62, 57, 101, 93, 59, 60, 0, 63, 54, 55,
	0, 0, 58, 0, 0, 0, 57, 0, 0, 0,
	0, 0, 63, 0, 57, 0, 0, 58, 0, 0,
	63, 0, 0, 0, 0, 58,
}
var yyPact = [...]int{

	117, -1000, -1000, 164, -1000, 89, 74, -1000, 200, 196,
	-1000, 145, 164, 88, 85, 72, -1000, -1000, 195, 195,
	61, 143, 68, -1000, 1, 58, -1000, -1000, 94, 195,
	-1000, -9, -29, 191, 60, 53, 112, -1000, 190, 188,
	-26, 97, 97, 186, -1000, -1000, 183, 47, -1000, 6,
	-1000, -1000, 24, 18, -1000, 182, -43, 459, 459, -1000,
	-1000, 181, 459, 84, 46, -19, -1000, 451, 169, 459,
	459, 459, 459, 459, 459, 459, 459, 459, 459, 459,
	459, 459, 459, -1000, -1000, 459, 180, 175, 459, -20,
	169, -1000, -18, -1000, 309, 7, 289, 125, 436, 459,
	166, -1000, -1000, -1000, 83, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 328, 57,
	328, -1000, -30, -31, 328, 459, 162, -1000, -1000, 459,
	158, 459, 97, -34, -1000, 51, 80, 8, 459, -1000,
	-1000, -1000, 29, 77, 269, 130, 249, 31, -1000, -1000,
	459, -1000, 328, -1000, 459, 97, 459, 459, 136, 17,
	16, 19, 229, 209, -1000, 11, -1000, -1000, 413, 97,
	459, 97, -1000, 10, 189, 5, 398, 97, 375, -1000,
	-4, -1000, 360, -1000,
}
var yyPgo = [...]int{

	0, 235, 19, 234, 226, 225, 136, 135, 24, 217,
	0, 216, 215, 214, 206, 205, 8, 144, 4, 168,
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
	20, 11, 11, 11, 12, 12,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 1, 5, 1, 3, 1, 3,
	0, 10, 9, 1, 0, 6, 8, 0, 1, 3,
	1, 3, 0, 2, 1, 3, 2, 3, 4, 1,
	3, 0, 1, 1, 1, 3, 3, 1, 5, 7,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 3, 2, 2,
	4, 1, 1, 1, 3, 4, 4, 4, 1, 1,
	3, 8, 10, 12, 6, 10,
}
var yyChk = [...]int{

	-1000, -3, -4, -5, -13, 14, -15, -6, 13, 15,
	5, 36, 38, 4, 4, -14, 5, -6, 36, 36,
	37, 39, -17, -7, 4, -17, 38, 5, 37, 39,
	-1, 4, 41, 31, 37, -19, 16, -7, 40, 42,
	4, 38, 38, 15, 4, 4, 40, -18, -8, -10,
	-11, -12, -2, 4, 17, 18, -9, 25, 36, 5,
	6, 7, 9, 31, -18, 4, 4, 38, 44, 34,
	35, 29, 30, 31, 32, 33, 19, 20, 21, 22,
	23, 24, 43, 27, 28, 36, 40, 41, 26, 4,
	44, -10, -2, 4, -10, 4, -10, 36, 38, 36,
	40, 11, -8, -20, 4, -10, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, -10, -10, -10, -16,
	-10, 4, 6, 5, -10, 36, 40, -20, 37, 8,
	39, 26, 38, -2, 11, -16, 4, 36, 39, 37,
	42, 42, -16, 4, -10, 4, -10, -18, 37, 37,
	36, 37, -10, 37, 36, 38, 8, 38, 38, -16,
	-16, -18, -10, -10, 11, 10, 37, 37, 38, 38,
	38, 38, 11, -18, -10, -18, 38, 38, 38, 11,
	-18, 11, 38, 11,
}
var yyDef = [...]int{

	3, -2, 1, 10, 4, 0, 2, 8, 0, 0,
	13, 0, 10, 0, 0, 0, 6, 9, 22, 22,
	0, 0, 0, 20, 0, 0, 5, 7, 14, 22,
	23, 24, 0, 0, 0, 0, 0, 21, 0, 0,
	26, 31, 31, 0, 27, 25, 0, 0, 29, 32,
	33, 34, 61, 68, 37, 0, 40, 0, 0, 62,
	63, 0, 0, 0, 0, 0, 28, 31, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 58, 59, 17, 0, 0, 0, 0,
	0, 56, 61, 68, 0, 0, 0, 0, 31, 17,
	0, 12, 30, 41, 69, 43, 44, 45, 46, 47,
	48, 49, 50, 51, 52, 53, 54, 55, 35, 0,
	18, 64, 0, 0, 36, 17, 0, 42, 57, 0,
	0, 0, 31, 0, 11, 0, 0, 0, 0, 60,
	65, 66, 0, 0, 0, 0, 0, 0, 67, 15,
	17, 70, 19, 38, 17, 31, 0, 0, 31, 0,
	0, 0, 0, 0, 74, 0, 16, 39, 31, 31,
	0, 31, 71, 0, 0, 0, 31, 31, 31, 72,
	0, 75, 31, 73,
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
		yyDollar = yyS[yypt-9 : yypt+1]
		//line template.y:51
		{
			yyVAL.iAstNode = &astWrapper{yyDollar[2].string, yyDollar[4].astList, yyDollar[7].astList}
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
		yyDollar = yyS[yypt-8 : yypt+1]
		//line template.y:133
		{
			yyVAL.iAstNode = &astRangeLoop{"_", yyDollar[2].string, yyDollar[4].iAstNode, yyDollar[6].astList}
		}
	case 72:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:135
		{
			yyVAL.iAstNode = &astRangeLoop{yyDollar[2].string, yyDollar[4].string, yyDollar[6].iAstNode, yyDollar[8].astList}
		}
	case 73:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line template.y:137
		{
			yyVAL.iAstNode = &astLoop{&astExpr{":=", &astValue{yyDollar[2].string}, yyDollar[4].iAstNode}, yyDollar[6].iAstNode, yyDollar[8].iAstNode, yyDollar[10].astList}
		}
	case 74:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line template.y:139
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, nil}
		}
	case 75:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line template.y:141
		{
			yyVAL.iAstNode = &astCondition{yyDollar[2].iAstNode, yyDollar[4].astList, yyDollar[8].astList}
		}
	}
	goto yystack /* stack new state and value */
}
