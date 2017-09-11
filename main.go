package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
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

	text := string(data)
	text = strings.TrimRight(text, "\n")

	template := template.New("formula.rb")
	template, err = template.Parse(text)
	if err != nil {
		return 1, "", err
	}

	buf := bytes.NewBufferString("")
	formula := newFormula()
	err = template.Execute(buf, formula)
	if err != nil {
		return 1, "", err
	}

	return 0, buf.String(), nil
}
