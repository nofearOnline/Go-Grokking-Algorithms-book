package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	utils "github.com/nofearOnline/go-algorithms/utils"
)

func testEq[K utils.BSComparable](a, b []K) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestIntMiddleSearch(t *testing.T) {
	rand.Seed(time.Now().Unix())
	arr := rand.Perm(10)
	sortedArr := make([]int, 10)

	copy(sortedArr, arr)
	sort.Ints(sortedArr)

	fmt.Println(arr, sortedArr)

	result := selectionSort(arr)

	fmt.Println(result, arr)

	if !testEq(result, sortedArr) {
		t.Fatalf(`selection sort returned: %v, the expected result was: %v `, result, sortedArr)
	}
}
