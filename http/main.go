package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	/*
		Condense using Writer Interface
	*/
	//Make empty byte slice with 99999 elements
	//bs := make([]byte, 99999)

	//Read function converts response body to byte slice
	//resp.Body.Read(bs)

	//Print the html code
	//fmt.Println(string(bs))

	//Condensed form of above lines: os.Stdout: Dst and resp.Body: source, source has to be something that implements Reader interface
	//io.Copy(os.Stdout, resp.Body)

	//logwriter implements writer
	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

/*
	This function enables logWriter to implement Writer interfaces

You need to be careful with the implementation of the function
*/
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("Just wrote %v lines", len(bs))
	return len(bs), nil
}
