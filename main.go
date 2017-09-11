package main

import (
	"fmt"
	"os"
)

// Version is the version of this application. This is updated when this package is compiled.
var Version = ""

func main() {
	code, message, _ := handle(os.Args[1:])
	if message != "" {
		fmt.Println(message)
	}
	os.Exit(code)
}

func handle(args []string) (int, string, error) {
	if len(args) > 0 {
		switch args[0] {
		case "-v", "--version":
			return 0, Version, nil
		}
	}

	return 0, "", nil
}
