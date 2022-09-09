package main

import "fmt"

func recursionSum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	return arr[0] + recursionSum(arr[1:])
}

func main() {
	fmt.Println(recursionSum([]int{1, 5, 67, -39, 9}))
}
