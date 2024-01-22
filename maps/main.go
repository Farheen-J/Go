package main

import "fmt"

func main() {
	/*
		Ways to declare maps
		1. variable := map[key datatype] datatype of values {}
		2. var variable map[key datatype] datatype of values
		3. variable := make(map[string]string)
	*/
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	//var colors map[string]string
	//colors := make(map[string]string)

	//Add
	colors["white"] = "#ffffff"
	//We cannot have colors.white in maps like structs

	//Delete a key
	//delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color + " : " + hex)
	}
}
