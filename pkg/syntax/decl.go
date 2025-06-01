package syntax

type Decl interface {
	decl()
}

type VarDecl struct {
	Name *Ident
	Expr Expr
	Type string
}

func (n *VarDecl) decl() {}

type FuncParam struct {
	Name *Ident
	Type string
}

type FuncDecl struct {
	Name *Ident
	Body *BlockStmt

	Params     []FuncParam
	ReturnType string
}

func (n *FuncDecl) decl() {}
