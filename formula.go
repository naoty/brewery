package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Formula represents Homebrew's formula.
type Formula struct {
	Name      string
	ClassName string
	Desc      string
	Homepage  string
	URL       string
	Sha256    string
}

func newFormula(path string, desc string, homepage string, url string) (Formula, error) {
	filename := filepath.Base(path)
	base := strings.Split(filename, ".")[0]

	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return Formula{}, err
	}

	contents, err := ioutil.ReadFile(absolutePath)
	if err != nil {
		return Formula{}, err
	}

	sum := sha256.Sum256(contents)
	encoded := hex.EncodeToString(sum[:])

	formula := Formula{
		Name:      base,
		ClassName: strings.Title(base),
		Desc:      desc,
		Homepage:  homepage,
		URL:       url,
		Sha256:    encoded,
	}
	return formula, nil
}
