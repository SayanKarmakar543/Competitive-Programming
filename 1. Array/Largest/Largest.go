package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50}
	var ans = FindLargestElement(arr)
	fmt.Println(ans)
}

var FindLargestElement = func(arr []int) int {
	var max = arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}
