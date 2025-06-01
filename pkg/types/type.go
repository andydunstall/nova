package types

type Type interface {
	String() string

	typeImpl()
}

type Primative int

const (
	Invalid Primative = iota

	Bool
	U8
	I8
	U16
	I16
	U32
	I32
	U64
	I64
)

func (t Primative) String() string {
	return primativeStrs[t]
}

func (t Primative) typeImpl() {}

var primativeStrs = [...]string{
	Invalid: "invalid",

	Bool: "bool",
	U8:   "u8",
	I8:   "i8",
	U16:  "u16",
	I16:  "i16",
	U32:  "u32",
	I32:  "i32",
	U64:  "u64",
	I64:  "i64",
}

var primatives map[string]Primative

func init() {
	primatives = make(map[string]Primative, len(primativeStrs))
	for i := 0; i != len(primativeStrs); i++ {
		if Primative(i) == Invalid {
			continue
		}
		primatives[primativeStrs[i]] = Primative(i)
	}
}

type Func struct {
	Params []*Object
	Return Type
}

func (t Func) String() string {
	s := "func("
	// TODO(andydunstall): Params
	s += ")"
	if t.Return != nil {
		s += " -> "
		s += t.Return.String()
	}
	return s
}

func (t Func) typeImpl() {}
