package main

import (
	"bytes"
	"errors"
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

func handle(args []string, opts options) (int, string, error) {
	if opts.version {
		return 0, Version, nil
	}

	if len(args) == 0 {
		err := errors.New("USAGE: brewery [options] <path>")
		return 1, "", err
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
	formula := newFormula(args[0], opts.desc, opts.homepage, opts.url)
	err = template.Execute(buf, formula)
	if err != nil {
		return 1, "", err
	}

	return 0, buf.String(), nil
}
