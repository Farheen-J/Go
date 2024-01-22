package main

import "fmt"

func main() {
	//In struct you need pointers but in a slice, you do not
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
