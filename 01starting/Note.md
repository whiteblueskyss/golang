# Go Basics: Notes from 01starting

## Executable Binaries and Entry Point

- **Only `main` can produce an executable binary; other packages are libraries.**
- The `main` function is the entry point for a Go program. Execution starts here.

---

## Variable Declaration

### Type Inference

```go
// infer - compiler automatically will determine the type of the variable.
var name = "golang"
```

### Explicit Type

```go
var price float32 = 50.5
```

### Type Inference with Default Type

```go
var price = 50.5
```

### Short Variable Declaration

```go
price := 50.5
```

- **All above declare the same type of variable (float) if not explicitly specified.**
- In shorthand variable declaration (`:=`), the variable must be initialized immediately.

---

## Valid and Invalid Variable Declarations

```go
var x int      // OK: type specified, zero value assigned
var x = 10     // OK: type inferred from value
x := 10        // OK: short declaration, type inferred

var x          // ERROR: missing type or initialization
```

---

## Constants

```go
const x = "ok"
```

- Constants must be assigned a value at declaration.
- The value must be known at compile time.
- You can specify the type or let it be inferred.
