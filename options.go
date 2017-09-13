package main

import "flag"

type options struct {
	desc     string
	url      string
	homepage string
	version  bool
}

func parseFlag() ([]string, options) {
	descShort := flag.String("d", "", "description")
	descLong := flag.String("desc", "", "description")
	urlShort := flag.String("u", "", "url")
	urlLong := flag.String("url", "", "url")
	homepageShort := flag.String("H", "", "homepage")
	homepageLong := flag.String("homepage", "", "homepage")
	versionShort := flag.Bool("v", false, "print version")
	versionLong := flag.Bool("version", false, "print version")
	flag.Parse()

	var desc string
	if *descShort == "" {
		desc = *descLong
	} else {
		desc = *descShort
	}

	var url string
	if *urlShort == "" {
		url = *urlLong
	} else {
		url = *urlShort
	}

	var homepage string
	if *homepageShort == "" {
		homepage = *homepageLong
	} else {
		homepage = *homepageShort
	}

	return flag.Args(), options{
		desc:     desc,
		url:      url,
		homepage: homepage,
		version:  *versionShort || *versionLong,
	}
}
