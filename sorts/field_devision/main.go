package main

import (
	"fmt"

	"github.com/nofearOnline/go-algorithms/utils/gcd"
)

type Size struct {
	x int
	y int
}

func biggestSquare(size Size) int {
	return gcd.GCD(size.x, size.y)
}

func main() {
	s := Size{1680, 640}

	fmt.Println(biggestSquare(s))
}
