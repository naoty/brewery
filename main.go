package main

import (
	"fmt"
	"os"
	"strings"
)

// Version is the version of this application. This is updated when this package is compiled.
var Version = ""

func main() {
	code, message, err := handle(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
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

	data, err := Asset("templates/formula.rb")
	if err != nil {
		return 1, "", err
	}

	formula := string(data)
	formula = strings.TrimRight(formula, "\n")

	return 0, formula, nil
}
