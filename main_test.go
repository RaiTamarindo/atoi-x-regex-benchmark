package main

import (
	"regexp"
	"testing"
)

func BenchmarkAtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validateWithAtoi("123456890")
	}
}

func BenchmarkRegex(b *testing.B) {
	r := regexp.MustCompile("[0-9]+")
	for i := 0; i < b.N; i++ {
		validateWithRegex(r, "123456890")
	}
}
