package main

import "fmt"

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}

	for _, url := range urls {
		fmt.Println(url)
	}
}
