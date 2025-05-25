package lex

import (
	"github.com/andydunstall/nova/pkg/assert"
)

// Token defines the lexical token types in Nova.
type Token int

const (
	EOF Token = iota

	// Identifiers and literals.
	literal_beg
	IDENT // foo
	INT   // 12345
	BOOL  // true
	literal_end

	// Operators.
	operator_beg
	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // %

	ADD_ASSIGN // +=
	SUB_ASSIGN // -=
	MUL_ASSIGN // *=
	QUO_ASSIGN // /=
	REM_ASSIGN // %=

	AND // &
	OR  // |
	XOR // ^
	SHL // <<
	SHR // >>

	INC // ++
	DEC // --

	LAND // &&
	LOR  // ||

	EQL    // ==
	NEQ    // !=
	LSS    // <
	GTR    // >
	ASSIGN // =
	NOT    // !
	LEQ    // <=
	GEQ    // >=

	LPAREN    // (
	LBRACE    // {
	RPAREN    // )
	RBRACE    // }
	COLON     // :
	SEMICOLON // ;
	COMMA     // ,
	ARROW     // ->
	TILDE     // ~
	operator_end

	// Keywords.
	keyword_beg
	FN
	RETURN

	LET
	MUT

	IF
	ELSE

	LOOP
	CONTINUE
	BREAK
	keyword_end
)

var tokens = [...]string{
	EOF: "EOF",

	IDENT: "IDENT",
	INT:   "INT",
	BOOL:  "BOOL",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	QUO: "/",
	REM: "%",

	ADD_ASSIGN: "+=",
	SUB_ASSIGN: "-=",
	MUL_ASSIGN: "*=",
	QUO_ASSIGN: "/=",
	REM_ASSIGN: "%=",

	AND: "&",
	OR:  "^",
	XOR: "^",
	SHL: "<<",
	SHR: ">>",

	INC: "++",
	DEC: "--",

	LAND: "&&",
	LOR:  "||",

	EQL:    "==",
	NEQ:    "!=",
	LSS:    "<",
	GTR:    ">",
	ASSIGN: "=",
	NOT:    "!",
	LEQ:    "<=",
	GEQ:    ">=",

	LPAREN:    "(",
	LBRACE:    "{",
	RPAREN:    ")",
	RBRACE:    "}",
	COLON:     ":",
	SEMICOLON: ";",
	COMMA:     ",",
	ARROW:     "->",
	TILDE:     "-~",

	FN:     "fn",
	RETURN: "return",

	LET: "let",
	MUT: "mut",

	IF:   "if",
	ELSE: "else",

	LOOP:     "loop",
	CONTINUE: "continue",
	BREAK:    "break",
}

func (tok Token) String() string {
	if 0 <= tok && tok < Token(len(tokens)) {
		return tokens[tok]
	}
	assert.Panicf("unknown token: %d", tok)
	return ""
}

var keywords map[string]Token

func init() {
	keywords = make(map[string]Token, keyword_end-(keyword_beg+1))
	for i := keyword_beg + 1; i < keyword_end; i++ {
		keywords[tokens[i]] = i
	}
}

// Lookup maps an identifier to its keyword token or [IDENT] (if not a
// keyword).
func Lookup(ident string) Token {
	if tok, is_keyword := keywords[ident]; is_keyword {
		return tok
	}
	return IDENT
}

func (tok Token) IsLiteral() bool {
	return literal_beg < tok && tok < literal_end
}

func (tok Token) IsOperator() bool {
	return (operator_beg < tok && tok < operator_end) || tok == TILDE
}

func (tok Token) IsKeyword() bool {
	return keyword_beg < tok && tok < keyword_end
}
