package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}

	for _, url := range urls {
		checkURL(url)
	}
}

func checkURL(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down!")
		return
	}

	fmt.Println(url, "is up!")
}
