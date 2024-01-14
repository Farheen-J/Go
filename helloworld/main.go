/*
Package = collection of common source code files
Two types of packages
1. Executable: When compiled gives an executable file that we can run - go build
2. Resuable: Code dependences/helpers/libraries
Name of the package determines which type we are using.
If the name is main = executable. If its anything else, its anything else, its reusable
An executable package should always have a function named
*/
package main

/*
Package fmt is standard library used to print out information
Other Go packages: debug, main, math, encoding, crypto, fmt, io
Third party packages: golang.org/pkg
*/
import "fmt"

/*
Functions
*/
func main() {
	fmt.Println("Hello World")
}

/*
Commands:
	go build
	go run
	go fmt
	go install
	go get
	go test
*/
