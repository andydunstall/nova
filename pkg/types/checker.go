package types

import (
	"fmt"

	"github.com/andydunstall/nova/pkg/assert"
	"github.com/andydunstall/nova/pkg/syntax"
)

// Check type checks the given file and returns the type info.
func Check(file *syntax.File) (*Info, error) {
	checker := newChecker()
	if err := checker.checkFile(file); err != nil {
		return nil, err
	}
	return checker.info, nil
}

type checker struct {
	info *Info
}

func newChecker() *checker {
	return &checker{
		info: newInfo(),
	}
}

func (c *checker) checkFile(file *syntax.File) error {
	for _, decl := range file.Decls {
		if err := c.checkDecl(decl); err != nil {
			return err
		}
	}
	return nil
}

// Statements.

func (c *checker) checkStmt(stmt syntax.Stmt) error {
	switch stmt := stmt.(type) {
	case *syntax.DeclStmt:
		return c.checkDecl(stmt.Decl)
	case *syntax.ReturnStmt:
		return c.checkReturnStmt(stmt)
	case *syntax.ExprStmt:
		// TODO(andydunstall)
		return nil
	case *syntax.BlockStmt:
		return c.checkBlockStmt(stmt)
	default:
		assert.Panicf("unsupported stmt type: %#v", stmt)
		return nil // Unreachable.
	}
}

func (c *checker) checkReturnStmt(stmt *syntax.ReturnStmt) error {
	return nil
}

func (c *checker) checkBlockStmt(stmt *syntax.BlockStmt) error {
	for _, stmt := range stmt.List {
		if err := c.checkStmt(stmt); err != nil {
			return err
		}
	}

	return nil
}

// Declarations.

func (c *checker) checkDecl(decl syntax.Decl) error {
	switch decl := decl.(type) {
	case *syntax.VarDecl:
		return c.checkVarDec(decl)
	case *syntax.FuncDecl:
		return c.checkFuncDec(decl)
	default:
		assert.Panicf("unsupported decl type: %#v", decl)
		return nil // Unreachable.
	}
}

func (c *checker) checkVarDec(decl *syntax.VarDecl) error {
	// TODO(andydunstall): Check conflicts.

	p, ok := primatives[decl.Type]
	if !ok {
		return fmt.Errorf("unknown type: %s", decl.Type)
	}

	c.info.Defs[decl.Name] = &Object{
		Name: decl.Name.Name,
		Type: p,
	}

	return nil
}

func (c *checker) checkFuncDec(decl *syntax.FuncDecl) error {
	// TODO(andydunstall): Check conflicts.

	// TODO(andydunstall)

	var params []*Object
	for _, param := range decl.Params {
		p, ok := primatives[param.Type]
		if !ok {
			return fmt.Errorf("unknown type: %s", param.Type)
		}

		o := &Object{
			Name: param.Name.Name,
			Type: p,
		}
		c.info.Defs[param.Name] = o

		params = append(params, o)
	}

	var ret Type
	if decl.ReturnType != "" {
		var ok bool
		ret, ok = primatives[decl.ReturnType]
		if !ok {
			return fmt.Errorf("unknown type: %s", decl.ReturnType)
		}
	}

	fn := &Func{
		Params: params,
		Return: ret,
	}
	c.info.Defs[decl.Name] = &Object{
		Name: decl.Name.Name,
		Type: fn,
	}

	if err := c.checkBlockStmt(decl.Body); err != nil {
		return err
	}
	return nil
}
