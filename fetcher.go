package main

import (
	"fmt"
	"sync"
)

// Fetcher is an example of an HTTP web page crawler.
type Fetcher interface {
	// Fetch returns the body of URL and a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// FakeFetcher is Fetcher that returns canned results.
type FakeFetcher map[string]*FakeResult

// FakeResult is the data returned from fetching a single HTTP page.
type FakeResult struct {
	body string
	urls []string
}

// Fetch returns the body of URL and a slice of URLs found on that page.
func (f FakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// FakeFetcherImpl is a populated FakeFetcher.
var FakeFetcherImpl = FakeFetcher{
	"https://golang.org/": &FakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &FakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &FakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &FakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

// SafeCache is a map of url to HTTP body, safe for concurrent use.
type SafeCache struct {
	// Map of url to body.
	cache map[string]string
	mux   sync.Mutex
}

// Add safely to cache.
func (safeCache *SafeCache) Add(url string, body string) {
	safeCache.mux.Lock()
	defer safeCache.mux.Unlock()
	safeCache.cache[url] = body
}

// Get from cache.
func (safeCache *SafeCache) Get(url string) (body string, ok bool) {
	safeCache.mux.Lock()
	defer safeCache.mux.Unlock()
	body, ok = safeCache.cache[url]
	return
}
