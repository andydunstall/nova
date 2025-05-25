package syntax

type Stmt interface {
	stmt()
}

type BlockStmt struct {
	List []Stmt
}

func (n *BlockStmt) stmt() {}

type ReturnStmt struct {
	Result Expr
}

func (n *ReturnStmt) stmt() {}

type ExprStmt struct {
	E Expr
}

func (n *ExprStmt) stmt() {}

type DeclStmt struct {
	Decl Decl
}

func (n *DeclStmt) stmt() {}

type IfStmt struct {
	Cond Expr
	Then Stmt
	Else Stmt
}

func (n *IfStmt) stmt() {}

type LoopStmt struct {
	Cond Expr
	Body *BlockStmt

	Label string
}

func (n *LoopStmt) stmt() {}

type BreakStmt struct {
	Label string
}

func (n *BreakStmt) stmt() {}

type ContinueStmt struct {
	Label string
}

func (n *ContinueStmt) stmt() {}
