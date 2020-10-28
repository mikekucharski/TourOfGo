// Programs start running in main package.
package main

import "fmt"

func printBanner(s string) {
	fmt.Printf("================= BEGIN %s =================\n", s)
}

func main() {
	printBanner("Basics")
	BasicsMain()
	printBanner("Flow Control")
	FlowControlMain()
	printBanner("Types")
	TypesMain()
	printBanner("Methods")
	MethodsMain()
	printBanner("Concurrency")
	ConcurrencyMain()
	printBanner("Concurrency2")
	TalkingGophers()
}
