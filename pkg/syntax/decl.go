package syntax

type Decl interface {
	decl()
}

type VarDecl struct {
	Name string
	Expr Expr
}

func (n *VarDecl) decl() {}

type FuncDecl struct {
	Name string
	Body *BlockStmt
}

func (n *FuncDecl) decl() {}
