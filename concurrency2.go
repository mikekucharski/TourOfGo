package main

import (
	"fmt"
	"math/rand"
	"time"
)

func talkToGeneratedChannel(gopher, msg string) chan string {
	c := make(chan string)
	// This creates a closure, because it references c variable outside of it.
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Message %d from %s says: %s\n", i, gopher, msg)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func talkToChannel(gopher, msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("Message %d from %s says: %s\n", i, gopher, msg)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

// TalkingGophers is an example from the slides on concurrency.
// https://talks.golang.org/2012/concurrency.slide
func TalkingGophers() {
	// talk to shared channel.
	c := make(chan string)
	go talkToChannel("emily", "bar", c)
	go talkToChannel("john", "foo", c)
	for i := 0; i < 5; i++ {
		fmt.Print(<-c)
	}

	// talk in lock step
	var sam, steve chan string
	sam = talkToGeneratedChannel("sam", "one")
	steve = talkToGeneratedChannel("steve", "two")
	for i := 0; i < 5; i++ {
		fmt.Print(<-sam)
		fmt.Print(<-steve)
	}

	// Fan-in approach

}
