package syntax

import "github.com/andydunstall/nova/pkg/lex"

type Expr interface {
	expr()
}

type UnaryExpr struct {
	Op   lex.Token
	Expr Expr
}

func (n *UnaryExpr) expr() {}

type BinaryExpr struct {
	Op lex.Token
	L  Expr
	R  Expr
}

func (n *BinaryExpr) expr() {}

type VarExpr struct {
	Name *Ident
}

func (n *VarExpr) expr() {}

type AssignExpr struct {
	L Expr
	R Expr
}

func (n *AssignExpr) expr() {}

type CallExpr struct {
	Func *Ident
	Args []Expr
}

func (n *CallExpr) expr() {}

type BasicLitExpr struct {
	Kind  lex.Token
	Value string
}

func (n *BasicLitExpr) expr() {}

type Ident struct {
	Name string
}

func (n *Ident) expr() {}
