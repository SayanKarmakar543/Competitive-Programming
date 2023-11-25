package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 20, 30, 40, 50}
	var status bool = CheckSortedArrayFunc(arr)
	fmt.Println(status)
}

var CheckSortedArrayFunc = func(arr []int) bool {
	var status bool = true
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			status = false
			break
		}
	}
	return status
}
