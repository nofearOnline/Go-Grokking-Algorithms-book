package main

import (
	utils "github.com/nofearOnline/go-algorithms/utils"
)

func selectionSort[K utils.BSComparable](arr []K) []K {
	//var emptyValue

	sortedArr := make([]K, len(arr))
	for i := range arr {
		min := arr[0]
		minI := 0
		for j := range arr {
			if arr[j] < min {
				min = arr[j]
				minI = j
			}
		}
		sortedArr[i] = min
		copy(arr[minI:], arr[minI+1:])
		arr = arr[:len(arr)-1]
	}

	return sortedArr
}
