package types

import (
	"github.com/andydunstall/nova/pkg/syntax"
)

type Info struct {
	Defs map[*syntax.Ident]*Object
}

func newInfo() *Info {
	return &Info{
		Defs: make(map[*syntax.Ident]*Object),
	}
}
