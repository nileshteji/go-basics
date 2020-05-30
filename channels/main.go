package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Asynce bitch
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://apple.com",
		"https://stackoverflow.com",
	}
	// this chan is an abbrevation for channel
	cha := make(chan string)

	for _, value := range links {
		// this automatically creates a new routine which will run in the background of the main thread
		go checkLink(value, cha)
	}
	for {

		time.Sleep(5 * time.Second)
		go checkLink(<-cha, cha)

	}

}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)

	if err != nil {
		// most probaly we are using this on Ui thread ?
		// channels are used to communicate between routines
		fmt.Println(link + " is down")

		c <- link

	} else {
		fmt.Println(link + " is up")
		c <- link

	}

}
