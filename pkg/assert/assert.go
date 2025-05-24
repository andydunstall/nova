package assert

import "fmt"

func Assert(pred bool, msg string) {
	if !pred {
		panic("assert: " + msg)
	}
}

func Assertf(pred bool, format string, a ...any) {
	if !pred {
		panic("assert: " + fmt.Sprintf(format, a...))
	}
}

func Panicf(format string, a ...any) {
	panic(fmt.Sprintf(format, a...))
}
