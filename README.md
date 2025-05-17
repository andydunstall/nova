# Nova

> :warning: In progress...

Nova is a compiled statically typed a programming language. It aims to be
a simplified combination of C and Rust.

Note Nova is only a toy project, so isn't expected to be used...

## Compiler

The Nova compiler, `novac`, can be build with Bazel:
```
bazel build //novac
```


## v0.1

This section describes the first version of Nova.

#### Files and Modules

Nova files are defined with a `.nv` extension.

In v0.1 only a single Nova file is supported, with no module support.

#### Data Types

Only integer and boolean types are supported in v0.1:
- `u8`
- `i8`
- `u16`
- `i16`
- `u32`
- `i32`
- `u64`
- `i64`
- `bool`

#### Variables

Variables are defined with the `let` keyword:
```
let a: i32 = 5;
let b: bool = false;
let mut c: u64 = 0xffffff00;
let mut d: u8 = 0b11001100;
```

Similar to Rust, variables are constant by default, though can be mutable
by setting the `mut` keyword.

The type can be omitted if it can be inferred from the initializer:
```
let a = 5;
let b = false;
```

#### Functions

Functions are defined with the `fn` keyword:
```
fn add(a: mut i32, b: mut i32) -> i32 {
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

Note unlike Rust, if-else statements are not expressions.

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

### Comments

Nova supports C style single line (`//`) comments.

#### Examples

See [`examples`](./examples) for examples of the Nova programming language.
