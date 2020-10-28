package main

import "fmt"
import "strings"

type vertex struct {
	x int
	y int
}

func pointers() {
	// Pointers work similar to C, but no pointer arithmatic.
	// Default value is nil.
	i := 42
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i
}

func structs() {
	// Structs
	v := vertex{1, 2}
	fmt.Println("Vertex struct is:", v)
	fmt.Println("Vertex x is:", v.x)

	// Pointers to Structs. You can derefernce normally or
	// dereference and access using just dot.
	structPointer := &v
	fmt.Println("Vertex x is:", (*structPointer).x)
	fmt.Println("Vertex x is:", structPointer.x)

	// Create struct literal by referencing field.
	var (
		v1        = vertex{1, 2}  // has type Vertex
		v2        = vertex{x: 1}  // Y:0 is implicit
		v3        = vertex{}      // X:0 and Y:0
		v1Pointer = &vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, v1Pointer, v2, v3)
}

func arrays() {
	// Array of 2 strings. Arrays cannot be resized. The length is part of the type.
	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "World"
	fmt.Println(arr[0], arr[1]) // indexing works as expected.
	fmt.Println(arr)

	// Can initialize values like this:
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	// Slices. Dynamically sized, flexible view into the array.
	// Slices do NOT store any data. They just describe parts of an underlying array.
	var slice []int = primes[2:4]
	fmt.Println(slice)

	// Changing a slice changes it's underlying array. If multiple slices point to
	// an array, they all will see the changes.
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println("Original list:", names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println("Overlapping slices:", a, b)
	b[0] = "XXX"
	fmt.Println("Modified slices:", a, b)
	fmt.Println("Modified underlying list:", names)

	// Slices are just arrays without length.
	// This is an array of 3 booleans.
	fmt.Println([3]bool{true, true, false})
	// This is a slice to an underlying array of 3 booleans.
	fmt.Println([]bool{true, true, false})
	// This is a slice of structs.
	structSlice := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}
	fmt.Println("StructSlice is:", structSlice)
	// Note, the default lower bound is 0, upper is length of array.
	// So, these are equivalent for array of length 10:
	// a[:], a[0:10], a[0:], a[:10]
	//
	// Slices have capacity cap(a) and length len(a). Capacity is length
	// of underlying array *starting from the first element of the slice*.
	// length is the number of elements in the slice.
	//
	// Slices can be resized, ie. re-sliced. a := arr[:2], a = arr[:4]
	//
	// Zero vallue of slice is nil. var nilSlice []int

	// Make. Make creates a zero-ed array and returns a slice to that array.
	// By default, it will make length == capacity unless you explicitly provide cap.
	zeroedArraySlice := make([]int, 5) // len(a)=5
	printSlice(zeroedArraySlice)

	b2 := make([]int, 0, 5) // len(b)=0, cap(b)=5
	b2 = b2[:cap(b2)]       // len(b)=5, cap(b)=5  Adjust the slice len to max cap.
	b2 = b2[1:]             // len(b)=4, cap(b)=4
	printSlice(b2)
}

func advancedSlices() {
	// You can have slices of slices! This is a 2D array of strings.
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// Modify board.
	board[0][0] = "X"
	board[1][1] = "O"
	// Print board.
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// You can append to slices! If the underlying array is not big enough,
	// it will allocate a new one.
	var s []int // nil slice.
	printSlice(s)
	s = append(s, 0) // append works on nil slices.
	printSlice(s)
	// The slice grows as needed. Note, during expansion the cap might become
	// larger than the len. Similar to table doubling concept.
	s = append(s, 1, 2, 3, 4) // len = 5 but cap = 6 after resize.
	printSlice(s)
}

func rangeLoop() {
	// Similar to python, you can range over a slice or map.
	// On each iteration of the loop, you get the index
	// and a *copy* of the element at that index.
	var pow = []int{1, 2, 4, 8, 16}
	for i, v := range pow {
		fmt.Printf("index=%d and value=%d\n", i, v)
	}

	// zeroed slice of 5
	pow = make([]int, 5)
	// you can avoid capturing the value of range by omitting it.
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	// You can avoid capturing the index of range by using '_'. Note, the same
	//  goes for ignoring the value but there is convenience way to do that.
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

// SlicesExercise from the go tour. They have a graphics drawing tool at
// golang.org/x/tour/pic.
func SlicesExercise(dx, dy int) [][]uint8 {
	// Allocate the top level slice.
	mySlice := make([][]uint8, dx)
	// Allocate the inner slices.
	for i := range mySlice {
		mySlice[i] = make([]uint8, dy)
	}
	// Assign value to each cell. Can use cool functions of x and y.
	// (x+y)/2, x*y, and x^y.
	for x := range mySlice {
		for y := range mySlice[x] {
			mySlice[x][y] = uint8((x + y) / 2)
		}
	}
	return mySlice
}

func printSlicesExercise(fn func(dx, dy int) [][]uint8) {
	fmt.Println(fn(5, 5))
}

// Declare a map of string -> vertex.
var m map[string]vertex

// Zero value of map is nil. nil has no keys and none can be added.
// Make returns a map of the given type.
func maps() {
	// initialize the map with make.
	m = make(map[string]vertex)
	m["p1"] = vertex{1, 2}
	fmt.Println(m["p1"])
	// returns the default value, vertex {0, 0}
	fmt.Println(m["missing"])

	// Create a map literal.
	var mapLiteral = map[string]vertex{
		"first": vertex{
			1, 2,
		},
		"second": vertex{
			3, 4,
		},
	}
	fmt.Println(mapLiteral)
	// Can also crate map literal like this:
	var mapLiteral2 = map[string]vertex{
		"first":  {1, 2},
		"second": {3, 4},
	}
	fmt.Println(mapLiteral2)

	m := make(map[string]int)
	// Set.
	m["foo"] = 1
	// Get.
	elem := m["foo"]
	fmt.Println(elem)
	// Check key exists by getting & assigning to two values.
	// Second one it boolean if exists.
	// If present is false, elem will be the default value.
	elem, present := m["foo"]
	fmt.Println(elem, present)
	// Delete.
	delete(m, "foo")
	fmt.Println(m["foo"])
}

// Maps exercise. Count the number of words in a string.
// Note, strings.Fields splits the string by space.
func wordCount(s string) map[string]int {
	frequency := make(map[string]int)
	var words []string = strings.Fields(s)
	for _, word := range words {
		frequency[word] = frequency[word] + 1
	}
	return frequency
}

// Take in a function which accepts int and return int, and just return it.
func passThroughFunc(input func(int) int) func(int) int {
	return input
}

func functionValues() {
	squareFunc := func(x int) int {
		return x * x
	}
	// Give the function to the passThroughFunc, which just returns
	// the same input func, then call the return func.
	fmt.Println(passThroughFunc(squareFunc)(5)) // prints 25
}

// A closure i sa function value that references variables from outside its body.
// Adder returns a closure. Sum is "bound" to the returned function, and can be
// updated, modified, and returned from the returned function.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// pos and neg are each holding references to the returned closures.
// Each closure has it's own sum variable "bound" to it.
func callClosure() {
	pos, neg := adder(), adder()
	for i := 0; i < 5; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

// TypesMain is the entry point for types.
func TypesMain() {
	pointers()
	structs()
	arrays()
	slices()
	advancedSlices()
	rangeLoop()
	printSlicesExercise(SlicesExercise)
	maps()
	fmt.Println(wordCount("This This is is a string"))
	functionValues()
	callClosure()
}
