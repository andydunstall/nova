package lex

import "fmt"

const (
	eof = 0xff // end of file
)

// Scanner parses a Nova file and outputs a stream of token literals.
//
// Comments are discarded.
type Scanner struct {
	src []byte

	ch     byte // current character
	offset int  // character offset

	pos Position
}

func NewScanner(src []byte) *Scanner {
	ch := byte(eof)
	if len(src) > 0 {
		ch = src[0]
	}
	return &Scanner{
		src:    src,
		ch:     ch,
		offset: 0,
		pos: Position{
			Line:   1,
			Column: 1,
		},
	}
}

func (s *Scanner) Scan() (tok Token, lit string, pos Position, err error) {
	s.skipWhitespace()

	pos = s.pos

	switch ch := s.ch; {
	case isLetter(ch):
		lit = s.scanIdentifier()
		tok = Lookup(lit)
	case isDecimal(ch):
		lit = s.scanNumber()
		tok = INT
	default:
		s.next()
		switch ch {
		case '+':
			if s.ch == '+' {
				tok = INC
				s.next()
			} else if s.ch == '=' {
				tok = ADD_ASSIGN
				s.next()
			} else {
				tok = ADD
			}
		case '-':
			if s.ch == '-' {
				tok = DEC
				s.next()
			} else if s.ch == '=' {
				tok = SUB_ASSIGN
				s.next()
			} else if s.ch == '>' {
				tok = ARROW
				s.next()
			} else {
				tok = SUB
			}
		case '*':
			if s.ch == '=' {
				tok = MUL_ASSIGN
				s.next()
			} else {
				tok = MUL
			}
		case '/':
			if s.ch == '=' {
				tok = QUO_ASSIGN
				s.next()
			} else {
				tok = QUO
			}
		case '%':
			if s.ch == '=' {
				tok = REM_ASSIGN
				s.next()
			} else {
				tok = REM
			}
		case '&':
			if s.ch == '&' {
				tok = LAND
				s.next()
			} else {
				tok = AND
			}
		case '|':
			if s.ch == '|' {
				tok = LOR
				s.next()
			} else {
				tok = OR
			}
		case '^':
			tok = XOR
		case '=':
			if s.ch == '=' {
				tok = EQL
				s.next()
			} else {
				tok = ASSIGN
			}
		case '!':
			if s.ch == '=' {
				tok = NEQ
				s.next()
			} else {
				tok = NOT
			}
		case '<':
			if s.ch == '<' {
				tok = SHL
				s.next()
			} else if s.ch == '=' {
				tok = LEQ
				s.next()
			} else {
				tok = LSS
			}
		case '>':
			if s.ch == '>' {
				tok = SHR
				s.next()
			} else if s.ch == '=' {
				tok = GEQ
				s.next()
			} else {
				tok = GTR
			}
		case '(':
			tok = LPAREN
		case '{':
			tok = LBRACE
		case ')':
			tok = RPAREN
		case '}':
			tok = RBRACE
		case ':':
			tok = COLON
		case ';':
			tok = SEMICOLON
		case ',':
			tok = COMMA
		case '~':
			tok = TILDE
		case eof:
			tok = EOF
		default:
			err = fmt.Errorf("unexpected character: %s", ch)
		}
	}

	return
}

func (s *Scanner) scanIdentifier() string {
	for i, b := range s.src[s.offset:] {
		if 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' || b == '_' || '0' <= b && b <= '9' {
			continue
		}

		ident := s.src[s.offset : s.offset+i]
		s.offset += i
		s.ch = s.src[s.offset]
		return string(ident)
	}

	// Will never get here as we know there is one character to scan.
	panic("eof")
}

func (s *Scanner) scanNumber() string {
	for i, b := range s.src[s.offset:] {
		if '0' <= b && b <= '9' {
			continue
		}

		ident := s.src[s.offset : s.offset+i]
		s.offset += i
		s.ch = s.src[s.offset]
		return string(ident)
	}

	// Will never get here as we know there is one character to scan.
	panic("eof")
}

func (s *Scanner) skipWhitespace() {
	for {
		if s.ch == ' ' || s.ch == '\t' || s.ch == '\n' || s.ch == '\r' {
			// Whitespace.
			s.next()
		} else if s.offset < len(s.src)-1 && s.ch == '/' && s.src[s.offset+1] == '/' {
			// Comment. Skip to next line.
			for s.ch != '\n' {
				s.next()
			}
		} else {
			break
		}
	}
}

func (s *Scanner) next() {
	if s.ch == '\n' {
		s.pos.Line++
		s.pos.Column = 0
	}

	s.pos.Column++
	if s.offset < len(s.src)-1 {
		s.offset++
		s.ch = s.src[s.offset]
	} else {
		s.offset = len(s.src)
		s.ch = eof
	}
}

func isLetter(ch byte) bool {
	return 'a' <= lower(ch) && lower(ch) <= 'z' || ch == '_'
}

func isDecimal(ch byte) bool { return '0' <= ch && ch <= '9' }

func lower(ch byte) byte { return ('a' - 'A') | ch }
