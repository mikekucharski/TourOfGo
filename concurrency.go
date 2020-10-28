package main

import (
	"./tree"
	"fmt"
	"sync"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func genValues(c chan int) {
	for i := 0; i < cap(c); i++ {
		c <- i
	}
	close(c)
}

func channelSelect(dataChan, quitChan chan int) {
	x, y := 0, 1
	for { // infinite loop
		// Note, this is a `select` not a switch. Select blocks until ONE of the
		// cases is ready, if multiple are, it picks at random.
		select {
		case dataChan <- x: // Can write to dataChan (buffer is not full)
			x, y = y, x+y
		case <-quitChan: // Can read from quitChan (buffer is not empty)
			fmt.Println("quit")
			return
			// Note, there is a default in select too. It runs when no other case it true,
			// So you can use it to test if a channel is ready to be read or written to.
			// ie. if it would block.
			//
			// case <read from channel>
			// default:
			//   // receiving from channel would block.
		}
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// SameInOrderTraversal determines whether the trees t1 and t2 contain the same values.
// However, it does so by explicitly reading 10 times from the channel.
// This only works because tree.Trees always have 10 values, by example.
func SameInOrderTraversal(t1, t2 *tree.Tree) bool {
	ch, ch2 := make(chan int, 1), make(chan int, 1)
	go Walk(t1, ch)
	go Walk(t2, ch2)
	var first, second int
	for i := 0; i < 10; i++ {
		first, second = <-ch, <-ch2
		if first != second {
			return false
		}
	}
	return true
}

// WalkAndClose is a wrapper method for Walk which
// also closes the stream after walking.
func WalkAndClose(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// SameInOrderTraversalGeneric is the same as SameInOrderTraversal but works in generic way.
// It reads from the channels until close,
func SameInOrderTraversalGeneric(t1, t2 *tree.Tree) bool {
	ch, ch2 := make(chan int, 1), make(chan int, 1)
	go WalkAndClose(t1, ch)
	go WalkAndClose(t2, ch2)
	var first, second int
	var ok, ok2 bool
	for {
		first, ok = <-ch
		second, ok2 = <-ch2
		// One of the channels closed before the other.
		if (ok && !ok2) || (!ok && ok2) {
			return false
		}
		// Both channels are closed. Compare the last element.
		if !ok {
			return first == second
		}
		// Check next two elements.
		if first != second {
			return false
		}
	}
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	m   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (counter *SafeCounter) Inc(key string) {
	// Lock so only one goroutine at a time can access the map c.v.
	counter.mux.Lock()
	counter.m[key]++
	counter.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (counter *SafeCounter) Value(key string) int {
	// Lock so only one goroutine at a time can access the map c.v.
	counter.mux.Lock()
	// Remember, defer executes immediately after a method returns. Good to not
	// forget to unlock at the end.
	defer counter.mux.Unlock()
	return counter.m[key]
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func crawl(url string, depth int, fetcher Fetcher, safeCache *SafeCache) {
	if depth <= 0 {
		return
	}
	// Don't fetch same URL twice!
	if cachedBody, ok := safeCache.Get(url); ok {
		fmt.Printf("Cache Hit: %s %q\n", url, cachedBody)
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("200 OK: %s %q\n", url, body)
	safeCache.Add(url, body)
	for _, u := range urls {
		// Spawn goroutine to recursively fetch url.
		go crawl(u, depth-1, fetcher, safeCache)
	}
	return
}

// ConcurrencyMain entry point for concurrency.
func ConcurrencyMain() {
	s := []int{7, 2, 8, -9, 4, 0}
	// Make an empty channel for go routines to write to. Channels can be
	// buffered, and make lets you set the buffer size (e.g. 100 here).
	// Blocks happen on write when buffer is full
	// Blocks happen on read when buffer is empty.
	//
	// If you overflow the buffer and nobody reads from it, the program goes into
	// deadlock.
	c := make(chan int, 100)
	// Spawn two go reoutines, each runs in their own thread.
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// Reading from a channel blocks.
	x, y := <-c, <-c // receive from c
	fmt.Println("First routine:", x, "Second routine:", y, "sum:", x+y)

	// You can close a channel with `close(c)`` and test if a channel is closed on read
	// `val, ok := <-ch`. You can also use `for i := range c {...}` to loop until close.
	// Note: Only the SENDER should close a channel. Sending on closed channel causes panic.
	// Note: Channels are NOT files / os resources, so you don't need to close them. Closing
	// is only necessary to tell receiver there are no more values coming.
	c2 := make(chan int, 3)
	go genValues(c2)
	for gened := range c2 {
		fmt.Println(gened)
	}

	dataChan := make(chan int)
	quitChan := make(chan int)
	go func() {
		// Read 10 times from first chan. Since channel size is 1, will block.
		for i := 0; i < 5; i++ {
			fmt.Println("Consumed:", <-dataChan)
		}
		// Only after all 10 reads from dataChan are done.
		quitChan <- 0
	}()
	channelSelect(dataChan, quitChan)

	// true
	fmt.Println(SameInOrderTraversal(tree.New(1), tree.New(1)))
	// false
	fmt.Println(SameInOrderTraversal(tree.New(1), tree.New(2)))
	fmt.Println(SameInOrderTraversal(tree.New(100), tree.New(200)))

	// true
	fmt.Println(SameInOrderTraversalGeneric(tree.New(1), tree.New(1)))
	// false
	fmt.Println(SameInOrderTraversalGeneric(tree.New(1), tree.New(2)))
	fmt.Println(SameInOrderTraversalGeneric(tree.New(100), tree.New(200)))

	// Mutexes.
	safeCounter := SafeCounter{m: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go safeCounter.Inc("somekey")
	}
	// Delay to give enough time for the goroutines to finish.
	time.Sleep(time.Second)
	fmt.Println(safeCounter.Value("somekey")) // 100

	// Note, the time.Sleep is a poor way to wait for all goroutines to complete.
	// Use WaitGroups instead for real impl: https://gobyexample.com/waitgroups
	safeCache := SafeCache{cache: make(map[string]string)}
	crawl("https://golang.org/", 4, FakeFetcherImpl, &safeCache)
	time.Sleep(time.Second)
}
