package mathlib  // changed from main to mathlib

func Add(a int, b int) int {  // Exported function (capitalized) so it can be used outside the package
	return addInternal(a, b)
}

func addInternal(a int, b int) int {  // Unexported function (lowercase) for internal package use only
	return a + b
}
