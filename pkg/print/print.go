package print

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

// Fprint pretty prints the given item to the writer.
func Fprint(w io.Writer, x any) error {
	return fprint(w, x)
}

// Fprint pretty prints the given item to stdout.
func Print(x any) error {
	return fprint(os.Stdout, x)
}

func fprint(w io.Writer, x any) error {
	p := NewPrinter(w)
	p.print(reflect.ValueOf(x))
	fmt.Fprintf(w, "\n")
	return nil
}

type Printer struct {
	output io.Writer
	indent int
	last   byte
	line   int
	ptrmap map[any]int // *T -> line number
}

func NewPrinter(output io.Writer) *Printer {
	return &Printer{
		output: output,
		line:   1,
		last:   '\n',
		ptrmap: make(map[any]int),
	}
}

var indent = []byte(".  ")

func (p *Printer) Write(data []byte) (n int, err error) {
	var m int
	for i, b := range data {
		// invariant: data[0:n] has been written
		if b == '\n' {
			m, err = p.output.Write(data[n : i+1])
			n += m
			if err != nil {
				return
			}
			p.line++
		} else if p.last == '\n' {
			_, err = fmt.Fprintf(p.output, "%6d  ", p.line)
			if err != nil {
				return
			}
			for j := p.indent; j > 0; j-- {
				_, err = p.output.Write(indent)
				if err != nil {
					return
				}
			}
		}
		p.last = b
	}
	if len(data) > n {
		m, err = p.output.Write(data[n:])
		n += m
	}
	return
}

type localError struct {
	err error
}

// printf is a convenience wrapper that takes care of print errors.
func (p *Printer) printf(format string, args ...any) {
	if _, err := fmt.Fprintf(p, format, args...); err != nil {
		panic(localError{err})
	}
}

func (p *Printer) print(x reflect.Value) {
	if !isNil(x) {
		p.printf("nil")
		return
	}

	switch x.Kind() {
	case reflect.Interface:
		p.print(x.Elem())

	case reflect.Map:
		p.printf("%s (len = %d) {", x.Type(), x.Len())
		if x.Len() > 0 {
			p.indent++
			p.printf("\n")
			for _, key := range x.MapKeys() {
				p.print(key)
				p.printf(": ")
				p.print(x.MapIndex(key))
				p.printf("\n")
			}
			p.indent--
		}
		p.printf("}")

	case reflect.Pointer:
		p.printf("*")

		ptr := x.Interface()
		if line, exists := p.ptrmap[ptr]; exists {
			p.printf("(obj @ %d)", line)
		} else {
			p.ptrmap[ptr] = p.line
			p.print(x.Elem())
		}

	case reflect.Array:
		p.printf("%s {", x.Type())
		if x.Len() > 0 {
			p.indent++
			p.printf("\n")
			for i, n := 0, x.Len(); i < n; i++ {
				p.printf("%d: ", i)
				p.print(x.Index(i))
				p.printf("\n")
			}
			p.indent--
		}
		p.printf("}")

	case reflect.Slice:
		if s, ok := x.Interface().([]byte); ok {
			p.printf("%#q", s)
			return
		}
		p.printf("%s (len = %d) {", x.Type(), x.Len())
		if x.Len() > 0 {
			p.indent++
			p.printf("\n")
			for i, n := 0, x.Len(); i < n; i++ {
				p.printf("%d: ", i)
				p.print(x.Index(i))
				p.printf("\n")
			}
			p.indent--
		}
		p.printf("}")

	case reflect.Struct:
		t := x.Type()
		p.printf("%s {", t)
		p.indent++
		first := true
		for i, n := 0, t.NumField(); i < n; i++ {
			// exclude non-exported fields because their
			// values cannot be accessed via reflection
			if name := t.Field(i).Name; IsExported(name) {
				value := x.Field(i)
				if first {
					p.printf("\n")
					first = false
				}
				p.printf("%s: ", name)
				p.print(value)
				p.printf("\n")
			}
		}
		p.indent--
		p.printf("}")

	default:
		v := x.Interface()
		switch v := v.(type) {
		case string:
			// print strings in quotes
			p.printf("%q", v)
			return
		}
		// default
		p.printf("%v", v)

	}
}

func IsExported(name string) bool {
	return name[0] >= 'A' && name[0] <= 'Z'
}

func isNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return !v.IsNil()
	}
	return true
}
