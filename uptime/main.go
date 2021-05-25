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
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go func(l string) {
			time.Sleep(time.Second * 5)
			checkLink(l, c)
		}(link)
	}

	for {
		go checkLink(<-c, c)
	}
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }
}

func checkLink(link string, c chan string) {
	// time.Sleep(time.Second * 5)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!!")
		c <- link
		return
	}
	fmt.Println(link, "is up :)")
	c <- link
}
