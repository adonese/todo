package main

import (
	"log"
	"regexp"
)

func capture(s string) string {
	r, err := regexp.Compile(`^(?P<to>to\s+)(?P<want>\w+.*)$`)
	if err != nil {
		panic(err)
	}
	log.Print(r.SubexpNames())

	res := r.FindStringSubmatch(s)
	if res == nil {
		return ""
	}
	return res[len(res)-1]
}

func captureAt(s string) string {
	r, err := regexp.Compile(`^(?P<to>at\s+)(?P<want>\w+.*)$`)
	if err != nil {
		panic(err)
	}
	log.Print(r.SubexpNames())

	res := r.FindStringSubmatch(s)
	if res == nil {
		return ""
	}
	return res[len(res)-1]
}
