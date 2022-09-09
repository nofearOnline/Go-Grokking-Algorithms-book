package main

import (
	"flag"
	"fmt"
	"strings"
)

func primeGenerator(index int) (prime int) {
	prime = 1

	for index > 0 {
		prime += 1

		primeFlag := true
		for dividerTest := 2; dividerTest < prime; dividerTest += 1 {
			if prime%dividerTest == 0 {
				primeFlag = false
			}
		}
		if primeFlag {
			index -= 1
		}
	}

	return prime
}

func main() {
	results := []int{}
	flag.Parse()

	flagCount := flag.NArg()

	for i := 0; i < flagCount; i += 1 {
		result := 0

		word := flag.Arg(i)

		word = strings.ToLower(word)

		for _, l := range word {
			result += primeGenerator(int(l) - 97)
		}

		results = append(results, result%10)
	}

	fmt.Printf("%+v", results)
}
