package main

import (
	"regexp"
	"strconv"
)

func validateWithAtoi(n string) bool {
	_, err := strconv.Atoi(n)
	return err != nil
}

func validateWithRegex(r *regexp.Regexp, n string) bool {
	return r.MatchString(n)
}

func main() {}
