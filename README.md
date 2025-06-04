# Nova

> :warning: In progress...

Nova is a compiled statically typed programming language, aiming to be
a simplified alternative of C++.

The compiler itself is written in Go (will rewrite in Nova once its mature
enough).

Note Nova is only a toy project, so isn't expected to be used...

## Compiler

Build the Nova compiler with `go build -o nova main.go`. Then invoke the compiler
with `nova build <path>`, such as `nova build examples/return.nv`.

The executable will be output to the same path as the input (with the
extension removed) by default, such as `examples/return.nv` will be output
to `examples/return`.

See `nova -h` for details.

## v0.1

This section describes the first version of Nova.

#### Files and Modules

Nova files are defined with a `.nv` extension.

In v0.1 only a single Nova file is supported, with no module support.

#### Data Types

Only integer, boolean and `struct` types are supported in v0.1:
- `u8`
- `i8`
- `u16`
- `i16`
- `u32`
- `i32`
- `u64`
- `i64`
- `bool`
- `struct`

Pointers are also supported, such as `*i32`.

#### Variables

Variables are defined with the `let` keyword:
```
let a: i32 = 5;
let b: bool = false;
let c: u64 = 0xffffff00;
let d: u8 = 0b11001100;
```

v0.1 doesn't support inferring types, so the type must be provided.

#### Functions

Functions are defined with the `fn` keyword:
```
fn add(a: i32, b: i32) -> i32 {
	return a + b;
}
```

#### Conditionals

If-else statements are supported with:
```
if (cond1) {
	// ...
} else if (cond2) {
	// ...
} else {
	// ...
}
```

#### Loops

v0.1 only supports a simple `loop` (equivalent to a while-loop in C):
```
loop (cond) {
	// ...
}
```

Or the condition can be omitted for an infinite loop:
```
loop {
	// ...
}
```

Both `break` and `continue` are supported.

### Structures

v0.1 supports structures with fields and members (including constructors
and destructors):
```
struct MyStruct {
	field1 u32
	field2 u32
};

// Constructor.
fn MyStruct::new(f1 u32, f2: u32) -> MyStruct {
	return MyStruct{
		field1: f1,
		field2: f2,
	};
}

// Destructor.
fn MyStruct::delete() {
	// ...
}

// Method.
fn MyStruct::sum() -> u32 {
	return self.field1 + self.field2;
}
```

### Comments

Nova supports C style single line (`//`) comments.

### Platforms

Nova only supports x86-64 Linux.

#### Examples

See [`examples`](./examples) for examples of the Nova programming language.
