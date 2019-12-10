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

	c := make(chan string)

	for _, url := range urls {
		go checkURL(url, c)
	}

	fmt.Println(<-c)
}

func checkURL(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprint(url, " might be down!")
		return
	}

	c <- fmt.Sprint(url, " is up!")
}
