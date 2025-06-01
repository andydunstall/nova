package syntax

import (
	"fmt"

	"github.com/andydunstall/nova/pkg/assert"
	"github.com/andydunstall/nova/pkg/lex"
)

// Parse parses the token stream into an AST.
func Parse(scanner *lex.Scanner) (*File, error) {
	p := newParser(scanner)
	f := p.parseFile()
	return f, nil
}

type parser struct {
	tok lex.Token
	lit string

	scanner *lex.Scanner

	line   int
	indent int
	debug  bool
}

func newParser(scanner *lex.Scanner) *parser {
	return &parser{
		scanner: scanner,
		line:    1,
		debug:   true,
	}
}

func (p *parser) parseFile() *File {
	// TODO(andydunstall): Handle position and errors.
	p.tok, p.lit, _, _ = p.scanner.Scan()

	if p.debug {
		defer un(trace(p, "File"))
	}

	var decls []Decl
	for p.tok != lex.EOF {
		decls = append(decls, p.parseDecl())
	}

	return &File{
		Decls: decls,
	}
}

// Expressions.

func (p *parser) parseExpr(minPrec int) Expr {
	if p.debug {
		defer un(trace(p, "Expr"))
	}

	l := p.parseFactor()
	for {
		prec := p.precedence(p.tok)
		if prec <= minPrec {
			break
		}

		if p.tok == lex.ASSIGN {
			l = p.parseAssignExpr(l, prec)
		} else {
			l = p.parseBinaryExpr(l, prec)
		}
	}

	return l
}

func (p *parser) parseAssignExpr(l Expr, prec int) *AssignExpr {
	if p.debug {
		defer un(trace(p, "AssignExpr"))
	}

	p.expect(lex.ASSIGN)
	return &AssignExpr{
		L: l,
		R: p.parseExpr(prec),
	}
}

func (p *parser) parseBinaryExpr(l Expr, prec int) *BinaryExpr {
	if p.debug {
		defer un(trace(p, "BinaryExpr"))
	}

	op := p.tok
	p.next()

	return &BinaryExpr{
		Op: op,
		L:  l,
		R:  p.parseExpr(prec + 1),
	}
}

func (p *parser) parseCallExpr(name *Ident) *CallExpr {
	if p.debug {
		defer un(trace(p, "CallExpr"))
	}

	var args []Expr

	p.expect(lex.LPAREN)
	for p.tok != lex.RPAREN {
		args = append(args, p.parseExpr(0))

		if p.tok != lex.RPAREN {
			p.expect(lex.COMMA)
		}
	}
	p.expect(lex.RPAREN)

	return &CallExpr{
		Func: name,
		Args: args,
	}
}

func (p *parser) parseFactor() Expr {
	if p.debug {
		defer un(trace(p, "Factor"))
	}

	switch p.tok {
	case lex.INT:
		f := &BasicLitExpr{
			Kind:  p.tok,
			Value: p.lit,
		}
		p.next()
		return f
	case lex.SUB, lex.TILDE, lex.NOT:
		op := p.tok
		p.next()
		expr := p.parseExpr(0)
		return &UnaryExpr{
			Op:   op,
			Expr: expr,
		}
	case lex.LPAREN:
		p.next()
		expr := p.parseExpr(0)
		p.expect(lex.RPAREN)
		return expr
	case lex.IDENT:
		name := p.parseIdent()
		if p.tok == lex.LPAREN {
			return p.parseCallExpr(name)
		} else {
			return &VarExpr{
				Name: name,
			}
		}
	default:
		panic("unknown: " + p.tok.String())
	}
}

// Statements.

func (p *parser) parseStmt() (s Stmt) {
	if p.debug {
		defer un(trace(p, "Stmt"))
	}

	switch p.tok {
	case lex.LBRACE:
		s = p.parseBlockStmt()
	case lex.RETURN:
		s = p.parseReturnStmt()
	case lex.LET:
		s = p.parseDeclStmt()
	case lex.IF:
		s = p.parseIfStmt()
	case lex.LOOP:
		s = p.parseLoopStmt()
	case lex.BREAK:
		s = p.parseBreakStmt()
	case lex.CONTINUE:
		s = p.parseContinueStmt()
	default:
		s = p.parseExprStmt()
	}
	return
}

func (p *parser) parseBlockStmt() *BlockStmt {
	if p.debug {
		defer un(trace(p, "BlockStmt"))
	}

	p.expect(lex.LBRACE)
	var list []Stmt
	for p.tok != lex.RBRACE && p.tok != lex.EOF {
		list = append(list, p.parseStmt())
	}
	p.expect(lex.RBRACE)
	return &BlockStmt{
		List: list,
	}
}

func (p *parser) parseReturnStmt() *ReturnStmt {
	if p.debug {
		defer un(trace(p, "ReturnStmt"))
	}

	p.expect(lex.RETURN)

	expr := p.parseExpr(0)
	p.expect(lex.SEMICOLON)
	return &ReturnStmt{
		Result: expr,
	}
}

func (p *parser) parseExprStmt() *ExprStmt {
	if p.debug {
		defer un(trace(p, "ExprStmt"))
	}

	expr := p.parseExpr(0)
	p.expect(lex.SEMICOLON)
	return &ExprStmt{
		E: expr,
	}
}

func (p *parser) parseDeclStmt() *DeclStmt {
	if p.debug {
		defer un(trace(p, "DeclStmt"))
	}

	return &DeclStmt{
		Decl: p.parseDecl(),
	}
}

func (p *parser) parseIfStmt() *IfStmt {
	if p.debug {
		defer un(trace(p, "IfStmt"))
	}

	p.expect(lex.IF)
	p.expect(lex.LPAREN)
	cond := p.parseExpr(0)
	p.expect(lex.RPAREN)
	thenStmt := p.parseStmt()

	var elseStmt Stmt
	if p.tok == lex.ELSE {
		p.next()
		elseStmt = p.parseStmt()
	}

	return &IfStmt{
		Cond: cond,
		Then: thenStmt,
		Else: elseStmt,
	}
}

func (p *parser) parseLoopStmt() *LoopStmt {
	if p.debug {
		defer un(trace(p, "LoopStmt"))
	}

	p.expect(lex.LOOP)
	p.expect(lex.LPAREN)
	cond := p.parseExpr(0)
	p.expect(lex.RPAREN)
	body := p.parseBlockStmt()
	return &LoopStmt{
		Cond: cond,
		Body: body,
	}
}

func (p *parser) parseBreakStmt() *BreakStmt {
	if p.debug {
		defer un(trace(p, "BreakStmt"))
	}

	p.expect(lex.BREAK)
	p.expect(lex.SEMICOLON)

	return &BreakStmt{}
}

func (p *parser) parseContinueStmt() *ContinueStmt {
	if p.debug {
		defer un(trace(p, "ContinueStmt"))
	}

	p.expect(lex.CONTINUE)
	p.expect(lex.SEMICOLON)

	return &ContinueStmt{}
}

// Declaration.

func (p *parser) parseDecl() Decl {
	if p.debug {
		defer un(trace(p, "Decl"))
	}

	switch p.tok {
	case lex.FN:
		return p.parseFuncDecl()
	case lex.LET:
		return p.parseVarDecl()
	default:
		panic("unsupported decl")
	}
}

func (p *parser) parseFuncDecl() *FuncDecl {
	if p.debug {
		defer un(trace(p, "FuncDecl"))
	}

	var funcDecl FuncDecl

	p.expect(lex.FN)
	funcDecl.Name = p.parseIdent()

	p.expect(lex.LPAREN)

	for p.tok != lex.RPAREN {
		var param FuncParam

		param.Name = p.parseIdent()

		// Parse type.
		p.expect(lex.COLON)
		typ := p.parseIdent()
		param.Type = typ.Name

		funcDecl.Params = append(funcDecl.Params, param)

		if p.tok != lex.RPAREN {
			p.expect(lex.COMMA)
		}
	}
	p.expect(lex.RPAREN)

	if p.tok == lex.ARROW {
		p.next()

		typ := p.parseIdent()
		funcDecl.ReturnType = typ.Name
	}

	funcDecl.Body = p.parseBlockStmt()
	return &funcDecl
}

func (p *parser) parseVarDecl() *VarDecl {
	if p.debug {
		defer un(trace(p, "VarDecl"))
	}

	p.expect(lex.LET)
	name := p.parseIdent()

	// Parse type.
	p.expect(lex.COLON)
	typ := p.parseIdent()

	p.expect(lex.ASSIGN)
	expr := p.parseExpr(0)
	p.expect(lex.SEMICOLON)

	return &VarDecl{
		Name: name,
		Expr: expr,
		Type: typ.Name,
	}
}

func (p *parser) parseIdent() *Ident {
	name := p.lit
	p.expect(lex.IDENT)
	return &Ident{
		Name: name,
	}
}

func (p *parser) expect(tok lex.Token) {
	if p.tok != tok {
		assert.Panicf("unexpected token: %s; wanted: %s", p.tok, tok)
		return // Unreachable.
	}
	p.next()
}

func (p *parser) next() {
	if p.debug {
		s := p.tok.String()
		switch {
		case p.tok.IsLiteral():
			p.printTrace(s, p.lit)
		case p.tok.IsOperator(), p.tok.IsKeyword():
			p.printTrace("\"" + s + "\"")
		default:
			p.printTrace(s)
		}
	}

	// TODO(andydunstall): Handle position and errors.
	p.tok, p.lit, _, _ = p.scanner.Scan()
}

func (p *parser) precedence(tok lex.Token) int {
	switch tok {
	case lex.MUL, lex.QUO, lex.REM:
		return 50
	case lex.ADD, lex.SUB:
		return 45
	case lex.LSS, lex.LEQ, lex.GTR, lex.GEQ:
		return 35
	case lex.EQL, lex.NEQ:
		return 30
	case lex.LAND:
		return 10
	case lex.LOR:
		return 5
	case lex.ASSIGN:
		return 1
	default:
		return -1
	}
}

func (p *parser) printTrace(a ...any) {
	fmt.Printf("%6d  ", p.line)

	const dots = ". . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . "
	const n = len(dots)
	i := 2 * p.indent
	for i > n {
		fmt.Print(dots)
		i -= n
	}
	// i <= n
	fmt.Print(dots[0:i])
	fmt.Println(a...)

	p.line++
}

func trace(p *parser, msg string) *parser {
	p.printTrace(msg, "(")
	p.indent++
	return p
}

// Usage pattern: defer un(trace(p, "..."))
func un(p *parser) {
	p.indent--
	p.printTrace(")")
}
