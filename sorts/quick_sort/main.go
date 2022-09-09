package main

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"time"

	"net/http"
	_ "net/http/pprof"

	"github.com/nofearOnline/go-algorithms/utils"
)

func quickSort[C utils.BSComparable](arr []C) []C {
	if len(arr) < 2 {
		return arr
	}
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return []C{arr[1], arr[0]}
		}
		return arr
	} else {
		pivot, arr := arr[0], arr[1:]
		var left []C
		var right []C
		for _, num := range arr {
			if num > pivot {
				right = append(right, num)
			} else {
				left = append(left, num)
			}
		}
		return append(append(quickSort(left), pivot), quickSort(right)...)
	}
}

func quickSortRandomPivot[C utils.BSComparable](arr []C) []C {
	if len(arr) < 2 {
		return arr
	}
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return []C{arr[1], arr[0]}
		}
		return arr
	} else {
		arrWithoutPivot := make([]C, 0)
		pivotIndex := rand.Intn(len(arr))
		pivot, arrWithoutPivot := arr[pivotIndex], append(append(arrWithoutPivot, arr[:pivotIndex]...), arr[pivotIndex:]...)
		var left []C
		var right []C
		for _, num := range arrWithoutPivot {
			if num > pivot {
				right = append(right, num)
			} else {
				left = append(left, num)
			}
		}
		return append(append(quickSort(left), pivot), quickSort(right)...)
	}
}

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	rand.Seed(time.Now().Unix())
	arr := rand.Perm(10)
	sortedArr := make([]int, 10)

	// arrWithoutPivot := make([]int, 0)
	// pivotIndex := rand.Intn(len(arr))
	// pivot, arrWithoutPivot := arr[pivotIndex], append(append(arrWithoutPivot, arr[:pivotIndex]...), arr[pivotIndex+1:]...)

	// fmt.Printf("'%+v', '%+v', '%+v', '%+v', \n", arr, pivotIndex, pivot, arrWithoutPivot)

	// using go sort for testing
	copy(sortedArr, arr)
	sort.Ints(sortedArr)

	fmt.Println("before sorts: ", arr, sortedArr)

	result := quickSort(arr)
	randomPivotResult := quickSortRandomPivot(arr)

	fmt.Println(arr, result, randomPivotResult)

	time.Sleep(300 * time.Second)
}
