package main

import (
	"fmt"
	"runtime"
	"time"
)

func loops() {
	// Note, no parenthesis.
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// There is no while, just for. The init and "post" are optional,
	// so you can simulate while like below.
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// To loop forever, just use "for {}"
}

func conditionals() {
	x := 1
	// Note, no parenthesis.
	if x > 1 {
		fmt.Println("abcd")
	}

	// You can initialize variables like "for" and use them in else.
	if y := 1; y > 1 {
		fmt.Println("xyz")
	} else {
		fmt.Println(y)
	}

	// You can switch-case over any type, and there are automatic "breaks" applied in each case.
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	// Cases don't have to be constants.
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// switch with no condition is like "switch true"
	// then each case statement is compared to true.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func defers() {
	// Defers executing this line until the surrounding function is over.
	// If you defer multiple times, they are pushed onto a stack. Once a
	// function finished, it's defers are popped & run one by one.
	defer fmt.Println("world")
	fmt.Println("hello")
}

// FlowControlMain is the entry point for flow control.
func FlowControlMain() {
	loops()
	conditionals()
	defers()
}
