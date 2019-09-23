package main

import (
	"regexp"
	"strconv"
)

func validateWithAtoi(n string) bool {
	_, err := strconv.Atoi(n)
	return err != nil
}

func validateWithRegex(n string) bool {
	exp := regexp.MustCompile("[0-9]+")
	return exp.MatchString(n)
}

func main() {}
