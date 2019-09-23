package main

import (
	"regexp"
	"testing"
)

var workloads []string
var worksize int

func setup(success bool) {
	if success {
		workloads = []string{
			"1",
			"12",
			"123",
			"1234",
			"12345",
			"123456",
			"1234567",
			"12345678",
			"123456789",
			"1234567890",
		}
	} else {
		workloads = []string{
			"123456789o",
			"abcdefghij",
			"ABCDEFGHIJ",
			"          ",
			"!@#$%*()[]",
		}
	}
	worksize = len(workloads)
}

func BenchmarkSuccessAtoi(b *testing.B) {
	setup(true)
	for i := 0; i < b.N; i++ {
		validateWithAtoi(workloads[i%worksize])
	}
}

func BenchmarkSuccessRegex(b *testing.B) {
	setup(true)
	r := regexp.MustCompile("[0-9]+")
	for i := 0; i < b.N; i++ {
		validateWithRegex(r, workloads[i%worksize])
	}
}

func BenchmarkFailAtoi(b *testing.B) {
	setup(false)
	for i := 0; i < b.N; i++ {
		validateWithAtoi(workloads[i%worksize])
	}
}

func BenchmarkFailRegex(b *testing.B) {
	setup(false)
	r := regexp.MustCompile("[0-9]+")
	for i := 0; i < b.N; i++ {
		validateWithRegex(r, workloads[i%worksize])
	}
}
