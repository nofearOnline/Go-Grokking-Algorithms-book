package main

import (
	"fmt"

	utils "github.com/nofearOnline/go-algorithms/utils"
)

type Sorted interface {
	search() int
}

func search[K utils.BSComparable](s []K, target K) int {
	min := 0
	max := len(s) - 1

	for max >= min {
		mid := (min + max) / 2
		if s[mid] == target {
			return mid
		}
		if s[mid] < target {
			min = mid + 1
		} else {
			max = mid
		}
	}
	return -1
}

func recursiveSearch[K utils.BSComparable](s []K, target K) int {
	arrSize := len(s)
	mid := arrSize / 2
	println(mid)

	if arrSize == 0 {
		return -1
	}
	if target == s[mid] {
		return mid
	}
	if target > s[mid] {
		return mid + recursiveSearch(s[mid+1:], target) + 1
	} else {
		fmt.Println(mid, s[:mid], target)
		return recursiveSearch(s[:mid], target)
	}
}

func main() {
	arr := []string{"aba", "bcc", "cdd", "gg", "hihi", "igebe", "jcxcw", "kdfs"}

	fmt.Println(arr[:3])

	res := search(arr, "gg")
	reqRes := recursiveSearch(arr, "gg")

	fmt.Println(res, reqRes)
}
