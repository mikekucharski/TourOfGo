package main

import (
	"fmt"
	"io"
	"math"
	"strings"
)

// Vertex - vertex
type Vertex struct {
	X, Y float64
}

// There are no classes in Go. However, you can attach functions to Types.
// This attaches abs() function to the Vertex type. The *receiver* is of
// type Vertex named v.
// You can also write like normal and pass v in.
// func abs(v Vertex) float64 {...}.
func (v Vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Declare an alias for a new type.
type customFloat float64

// You can even declare a function with a receiver,
// IFF the type is declared in the same package as you.
func (cf customFloat) abs() customFloat {
	if cf < 0 {
		return cf * -1
	}
	return cf
}

// You can also declare a method on *pointer receivers*. A method
// which has receives of "v Vertex" recieves a COPY (pass by value) of v.
// Pointer receivers are more common than value ones because usually
// you want to modify the receiver.
func (v *Vertex) scale(factor float64) {
	v.X = v.X * factor
	v.Y = v.Y * factor
}

func methods() {
	v := Vertex{3, 4}
	fmt.Println(v.abs())

	var cf customFloat = -3.2
	fmt.Println(cf.abs())
	// You can also make new custom type this way.
	cf2 := customFloat(-3.2)
	fmt.Println(cf2.abs())

	// Receivers.
	// Scale by a factor of 2.
	// NOTE. If you use regular methods (not receivers) then you can pass by
	// value or reference to them. ie. you can define scale(v Vertex)
	// or scale(v *Vertex) and you call them with v or &v respectively.
	//
	// For pointer receivers, you can call on a value OR on a reference.
	// v.scale(2) is syntax sugar for (&v).scale(2) because scale is defined
	// with POINTER receiver.
	v.scale(2)
	fmt.Println(v)

	// Method indirection on pointers works in the opposite way as well.
	var vert Vertex
	fmt.Println(vert.abs()) // OK
	p := &vert
	// This is syntax sugar for (*p).abs() because abs() is defined with
	// VALUE receiver.
	fmt.Println(p.abs()) // OK

	// In general, all methods on a given  type should use EITHER value
	// OR pointer receivers, and not mix both. Pointer receivers have 2 main
	// advantages, you can modify the recevier and the data is not copied
	// on each call, which can be memory consuming even you don't need to
	// modify the value.
}

// Abser - abser
type Abser interface {
	abs() float64
}

// Animal - animal
type Animal interface {
	numLegs() int
}

// Dog - dog
type Dog struct {
	name string
}

func (dog Dog) numLegs() int {
	return 4
}

// Attach the same numLegs to type "Pointer of customFloat" and it works.
func (cf *customFloat) numLegs() int {
	if cf == nil {
		return -100
	}
	return 100
}

// Takes in an empty interface.
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func interfaces() {
	var abser Abser
	// Vertex does not need to declare that it implements abser. It is
	// implied based on the methods in the implementation.
	myVertex := Vertex{1.2, 2.4}
	abser = myVertex // allowed because myVertex is an Abser
	fmt.Println(abser.abs())

	// You can attach methods to the TYPE or the POINTER to the type
	// and this is important when determining if a type implements
	// the interface.
	var animal Animal
	animal = Dog{"Barney"}
	fmt.Println(animal.numLegs())

	// Now, Pointers to 'customFloat' implement interface Animal.
	// Assigning just customFloat type to animal does NOT work.
	var custFloat customFloat = 5.6
	animal = &custFloat
	fmt.Println(animal.numLegs())

	// This shows interface values should be thought of as a tuple between
	// a value and a concrete type, like this: (value, type). When you call
	// a method on an interface, the method of the underlying type is called.
	//
	// The underlying / concrete type can be nil. Calling the method does not
	// throw NPE. Instead, you can define what to do with nil types.
	var uninitializedFloatPointer *customFloat
	animal = uninitializedFloatPointer
	// animal reference holds nil for the value, but it is not nil itself.
	// The tuple would be (nil, *customFloat). Printing it returns the
	// value (nil). But calling a method on it still calls the method on the
	// underlying type!
	fmt.Println(animal)           // nil
	fmt.Println(animal.numLegs()) // -100

	// If you don't assign a type to an interface, the tuple is (nil, nil)
	var nilAnimal Animal
	fmt.Println(nilAnimal) // nil
	// fmt.Println(nilAnimal.numLegs()) // throws runtime exception.

	// There are also empty interfaces. Empty interfaces look like interface{}
	// and can hold the value of any type.
	var anything interface{}
	anything = 42
	describe(anything)
	anything = "hello"
	describe(anything)
}

type person struct {
	name string
}

func (p person) String() string {
	return "Person has awesome name: " + p.name
}

func moreInterfaces() {
	// Type assertions provide access to an interfaces value's concrete type.
	var i interface{}
	i = Dog{"sparky"}
	// This asserts that you want to assign the values of i to t, but ONLY if
	// its a Dog type. If its not an instance of Dog, panic.
	t := i.(Dog)
	describe(t)

	// Similar to maps, this interface type assertions can return 2 values!
	// The first is the value (or default value), second is boolean if it
	// really is that type or not.
	val, ok := i.(Dog)
	fmt.Println(val, ok)

	// i does not hold a string type, so ok is false and val is the default value.
	// Note, we cannot reassign val, because compile time checks and sees that we
	// actually stored Dog type in `i` last.
	val2, ok := i.(string)
	fmt.Println(val2, ok)

	// panic. Basically exception is thrown because we "get" from the interface
	// without checking the internal type or capturing in the 2 argument return value.
	// f := i.(float64)
	// fmt.Println(f)

	// Type swtich can switch case the TYPE of interface, not value. This is like calling
	// i.(XXX) where XXX is replaced with the case statment type.
	i = Dog{"sparky"}
	switch v := i.(type) {
	// Note, we put a Dog in v, but this triggers first because Dog is Animal.
	case Animal:
		fmt.Println("it was Animal!")
	case Dog:
		fmt.Println("it was Dog!")
	case customFloat:
		fmt.Println("it was customFloat!")
	case *customFloat:
		// Pointer type
		fmt.Println("it was *customAnimal!")
	default:
		// v is assigned the same type and value as i, but it was not captured in this switch.
		fmt.Printf("Not sure, it was %T\n", v)
	}

	// Stringer interface has one `String() string` method, declared in `fmt` package.
	// It is basically like a toString().
	// Notice how we defined String() method on person, and fmt is calling that method to print the value.
	fmt.Println(person{"mike"})
}

// ErrNegativeSqrt - Custom type for the error.
type ErrNegativeSqrt float64

// Make the type an Error by implementing the method.
func (e ErrNegativeSqrt) Error() string {
	// Note, we need to cast e to float64 here or stack overflow happens.
	return fmt.Sprint("cannot Sqrt negative number:", float64(e))
}

// Now you can return the ErrNegativeSqrt in place of interface `error`.
func safeSqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func errors() {
	// There is built-in interface `type error interface { Error() string }`
	// The APIs of functions can return error as the second argument, for the caller to test against.
	fmt.Println(safeSqrt(-2))
	fmt.Println(safeSqrt(2))
}

func readers() {
	// io.Reader interface looks like this: `func (T) Read(b []byte) (n int, err error)`
	// There are many implementations in Go, to read from files, networks, etc.
	// Read() populates the given byte slice with data, and returns the number of bytes populated.
	// it returns io.EOF from the error if the stream ends.

	// String reader impl.
	r := strings.NewReader("Hello, Reader!")

	// Make slice of 8 bytes.
	b := make([]byte, 8)
	for {
		n, err := r.Read(b) // will read 8 bytes at a time.
		fmt.Printf("n = %v  |   err = %v  |  b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n]) // %q is used to double quote strings.
		if err == io.EOF {
			break
		}
	}
}

// MethodsMain is the entry point for Methods.
func MethodsMain() {
	methods()
	interfaces()
	moreInterfaces()
	errors()
	readers()
}
