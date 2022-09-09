package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestIntMiddleSearch(t *testing.T) {
	rand.Seed(time.Now().Unix())

	arr := rand.Perm(10)
	sort.Ints(arr)

	result := search(arr, arr[3])

	if result != 3 {
		t.Fatalf(`search("<string in place 3>") = %v, instead of 3 `, result)
	}
}

func TestStringLastElementSearch(t *testing.T) {
	arr := []string{"aba", "bcc", "cdd", "gg", "hihi", "igebe", "jcxcw", "kdfs"}

	i := search(arr, "kdfs")

	if i != 7 {
		t.Fatalf(`search("kdfs") = %v, instead of 7 `, i)
	}
}
