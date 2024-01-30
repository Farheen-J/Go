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

	//make channel of type string
	c := make(chan string)

	for _, link := range links {
		//go keyword ensures each link will have a new go routine
		go checkLink(link, c)
	}

	//We need loop because main routine needs to print 5 go routine messages
	// for i := 0; i < len(links); i++ {
	// 	//Receiving value from channel
	// 	//fmt.Println(<-c)
	// }

	// Infinite for loop
	// for {
	// 	//status checker
	// 	go checkLink(<-c, c)
	// }

	//Alternate syntax for above: for each value l returned by channel
	for l := range c {
		/*
			Pause for 5 seconds
			If we block this in main routine, this is a blocking call. Main routine will wait for message for 5 seconds
			Not appropriate here
			It is not appropriate in checkLink either because that questions its purpose of instance checks
			time.Sleep(5 * time.Second)
			instead use function literal.
			Also, use pass by value functionality. Pass l as the argument
		*/
		//status checker
		go func(link string) {
			time.Sleep(5 * time.Second)
			//Here, l in child routine == l in main routine, hence same link unless passed as an argument
			checkLink(link, c)
		}(l) //Passing l as a parameter
	}
}

func checkLink(link string, c chan string) {
	_, error := http.Get(link) //Blocking call
	if error != nil {
		fmt.Println(link, " might be down.")

		//Sending value to channel
		//c <- "It might be down."
		//We send link in the channel for status checker
		c <- link
		return
	}

	fmt.Println(link, " is up!")
	//Sending value to channel
	c <- link
}
