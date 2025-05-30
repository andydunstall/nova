package syntax

type Decl interface {
	decl()
}

type VarDecl struct {
	Name *Ident
	Expr Expr
}

func (n *VarDecl) decl() {}

type FuncDecl struct {
	Name *Ident
	Body *BlockStmt
}

func (n *FuncDecl) decl() {}
