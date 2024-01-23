package main

import (
	"fmt"
	"io"
	"os"
)

type fileReader struct {
}

func main() {
	//Cmd arguments
	fmt.Println(os.Args)
	content, error := os.Open(os.Args[1])

	if error != nil {
		fmt.Printf("Unable to open the file: %v", error)
	}

	//Old format
	// bs := make([]byte, 999999)
	// content.Read(bs)
	// fmt.Println(string(bs))

	//New format
	io.Copy(os.Stdout, content)
}
