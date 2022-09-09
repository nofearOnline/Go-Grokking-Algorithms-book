package main

import "fmt"

func recursionMax(arr []int) int {
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return arr[0]
		}
		return arr[1]
	}
	tempArr := make([]int, 0)
	return recursionMax(append(tempArr, arr[0], recursionMax(arr[1:])))

}

func main() {
	fmt.Println(recursionMax([]int{1, 5, 67, -39, 9, 89, -99}))
}
