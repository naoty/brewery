package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Version is the version of this application. This is updated when this package is compiled.
var Version = ""

func main() {
	args, opts := parseFlag()
	code, message, err := handle(args, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if message != "" {
		fmt.Println(message)
	}
	os.Exit(code)
}

type options struct {
	version bool
}

func parseFlag() ([]string, options) {
	versionShort := flag.Bool("v", false, "print version")
	versionLong := flag.Bool("version", false, "print version")
	flag.Parse()

	return flag.Args(), options{version: *versionShort || *versionLong}
}

func handle(args []string, opts options) (int, string, error) {
	if opts.version {
		return 0, Version, nil
	}

	data, err := Asset("templates/formula.rb")
	if err != nil {
		return 1, "", err
	}

	formula := string(data)
	formula = strings.TrimRight(formula, "\n")

	return 0, formula, nil
}
