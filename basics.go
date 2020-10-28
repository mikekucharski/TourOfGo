// Programs start running in main package.
package main

// Everything defined with capital letter in a package is imported.
// Can import like this: import "fmt"
import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand" // imports files with "package rand"
	"time"
)

// You can also declare variables un the "factored" style.
// Supported types are bool, string, int, byte, float32/64, uintpointer and more.
// Default values are assigned, 0 for numeric. false for boolean. "" for string.
//
// Note, There is no character data type!
// Go is always UTF-8
// chars are represented with `byte`
// `rune` is used for Unicode "code point", it's also an alias for int32.
// See https://blog.golang.org/strings for details.
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// Constants.
// Can also use const Pi = 3.14.
// Can be character, string, boolean, numeric.
// Cannot use := syntax.
const (
	Pi         = 3.14
	DaysOfWeek = 7
	// Shift 1 left 100 times.
	Big = 1 << 100
	// Shift Big right 99 times, leaving 1 << 1, ie value of 2.
	Small = Big >> 99
)

// Functions. Arguments come after variable names. Return type comes after method.
// Can use (x int, y int) or combine matching types with (x,y int)
func add(x, y int) int {
	return x + y
}

// Functions can return more than 1 variable.
func returnTwo() (int, int) {
	return 1, 2
}

// Functions can have named return values. Then you can just "return".
func returnTwoNamed() (x int, y int) {
	x = 1
	y = 2
	return // "naked" return statement. Only use for short methods.
}

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

// Var declares list of variables. Can be at package level or function level.
// var statements can have initializers, even with multiple.
// Note, the type can be omitted if defaults are specified.
var packageVar1, packageVar2 bool = true, true
var myBool, myInt, myFloat, myString = true, 100, 50.2, "hello"

// BasicsMain is th entry point for the basics.
func BasicsMain() {
	// Printf works just like in C.
	// See https://gobyexample.com/string-formatting for formatting.
	fmt.Printf("hello, world!\n")
	fmt.Printf("My favorite number is %d and %f\n", rand.Intn(10), math.Pi)
	fmt.Println("The time is:", time.Now())

	// Format string can also be %T for type, %v for value.
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)

	// var declares list of variables. Can be at package level or function level.
	var functionVar1, functionVar2 bool
	fmt.Println("Package vars:", packageVar1, packageVar2)
	fmt.Println("Package vars with inferred types:", myBool, myInt, myFloat, myString)
	fmt.Println("Function vars:", functionVar1, functionVar2)

	// = is for assignment, := is for declaration + assignment. `var foo int = 10` is the same as `var foo = 10` is the same as `foo := 10`
	// Note, you can't assign with a, b := ... again because they're already defined.
	// But, you CAN use it if even 1 variable is not defined yet. a,c := ... works.
	// Types are inferred automatically. Do not need to declare int, int.
	// := is not available outside a function.
	a, b := returnTwo()
	fmt.Println("Return two:", a, b)

	a1, b1 := returnTwoNamed()
	fmt.Println("Return two named:", a1, b1)

	// Type Conversions must be explicit, unlike C. Removing these casts fails to compile.
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Printf("Casting types: %T, %T, %T\n", i, f, u)

	// If you use type inferrence, like with := operator,  numeric types are inferred based
	//  on precision.
	i2 := 42           // int
	f2 := 3.142        // float64
	g2 := 0.867 + 0.5i // complex128
	fmt.Printf("Inferred types: %T, %T, %T\n", i2, f2, g2)

	// Constants.
	const Truth = true
	fmt.Println("Constants:", Truth, Pi, DaysOfWeek)

	// Numeric constants.
	// An untyped constant takes the type needed by its context.
	// ie. needInt(Small) and needFloat(SMALL) output different things..
	fmt.Println("Types inferred from context:",
		needInt(Small), needFloat(Small), needFloat(Big))
	// gives compiler error "Big overflows int". Max is 64 bit integer.
	// fmt.Println(needInt(Big))
}
