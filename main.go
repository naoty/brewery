package main

import (
	"fmt"
	"os"
)

// Version is the version of this application. This is updated when this package is compiled.
var Version = ""

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-v", "--version":
			fmt.Println(Version)
		}
	}

	os.Exit(0)
}
