Only main can produce an executable binary; other packages are libraries.

The main function is the entry point for a Go program. Execution starts here

    // infer - compiler automatically will determine the type of the variable.
    // var name = "golang"


    // var price float32 = 50.5
    // var price = 50.5
    // price := 50.5

    ALL SAME TYPE.

    	// shorthand syntax
    // name := "golang"
    in shorthand varialbe should be initialized immediatly.

var x int // OK: type specified, zero value assigned
var x = 10 // OK: type inferred from value
x := 10 // OK: short declaration, type inferred

var x // ERROR: missing type or initialization

const x = "ok"
Constants must be assigned a value at declaration.
The value must be known at compile time.
You can specify the type or let it be inferred.
