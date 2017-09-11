package main

// Formula represents Homebrew's formula.
type Formula struct {
	Name      string
	ClassName string
	Desc      string
	Homepage  string
	URL       string
	Sha256    string
}

func newFormula() Formula {
	return Formula{
		Name:      "todo",
		ClassName: "Todo",
		Desc:      "A todo management tool just for myself",
		Homepage:  "https://github.com/naoty/todo",
		URL:       "https://github.com/naoty/todo/releases/1.0.0/todo.tar.gz",
		Sha256:    "050cfb85006b1238f4d2c51f82ad52b212e852ff1292796e2706b9e7ab532c4f",
	}
}
