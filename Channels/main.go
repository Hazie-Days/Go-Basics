package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.com",
		"https://amazon.com",
	}
	// intitialising a channel of type string
	c := make(chan string)

	// Accessing every request using a loop
	for _, link := range links {
		go checkLink(link, c)
	}
	// run through the loop whenever the channels emits a value
	for l := range c {
		// function literal
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

// checking Child Routine

func checkLink(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
