package main

import "fmt"

/*
To all types whom it may concern, if you have a function called getGreeting()
associated with it then that datatype is a member of interface bot
*/
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// Since we are not using eb value inside the function, we just provide the datadtype englishBot
func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// No function overloading
// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }
