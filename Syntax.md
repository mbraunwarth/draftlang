# Draft Syntax

## Symbols

Introduce a new variable with the `let` keyword. The type can be either inferred
or explicitly annotaded, separated by the name with a colon.

```
// explicit type annotation
let NAME: TYPE = VALUE

// type inferred by compiler
let NAME = VALUE
```

A symbols scope is delimited by its block. Blocks may have access symbols introduced
in themself and to those introduced in nested blocks.

## Functions

A function is defined with the `fn` keyword followed by its name. An left arrow (`->`) 
separates the arguments from the return type, then follows the body enclosed by curly 
braces. If the function does not return anything (TODO this will be covered later) the
arrow and the return type can be omitted.

```
fn NAME (ARGS) -> RET_TYPE {...}

// example - omit return type
fn printHello() { println("Hello, World!") }
```


### Lambda Functions

A lambda function is defined with the `fn` keyword followed by its arguments, since
lambdas are anonymous functions and don't need a name. An left arrow (`->`) separates
the arguments from the return type, then follows the body enclosed by curly braces.
If the function does not return anything (TODO this will be covered later) the
arrow and the return type can be omitted.

`fn (ARGS) -> RET_TYPE {...}`

Although lambdas are nameless, they can be bound to a symbol and invoked later in the
symbols scope.

### Closures

## Keywords

TODO
