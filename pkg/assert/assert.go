package assert

import (
	"fmt"
	"runtime"
)

func Panicf(format string, a ...any) {
	// Include the location as it could be buried in the middle
	// of the panicking stack.
	prefix := ""
	if _, file, line, ok := runtime.Caller(1); ok {
		prefix = fmt.Sprintf("%s:%d: ", file, line)
	}
	panic(prefix + fmt.Sprintf(format, a...))
}

func Assert(pred bool, msg string) {
	assertf(pred, msg)
}

func Assertf(pred bool, format string, a ...any) {
	assertf(pred, format, a...)
}

func assertf(pred bool, format string, a ...any) {
	if !pred {
		// Include the assertion location as it could be buried in the middle
		// of the panicking stack.
		//
		// Use a depth of 2, since we want the caller or Assert/Assertf, not
		// the caller of assertf.
		prefix := ""
		if _, file, line, ok := runtime.Caller(2); ok {
			prefix = fmt.Sprintf("%s:%d: ", file, line)
		}

		panic("assert: " + fmt.Sprintf(prefix+format, a...))
	}
}
