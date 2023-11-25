package main

import "fmt"

func main() {
	arr := []int{1, 4, 2, 10, 23, 3, 1, 0, 20}
	k := 4
	fmt.Println(SlidingWindowFunc(arr, k))
}

func SlidingWindowFunc(arr []int, k int) int {
	brr := []int{}
	for i := 0; i <= len(arr)-k; i++ {
		sum := 0
		for j := i; j < i+k; j++ {
			sum += arr[j]
		}
		brr = append(brr, sum)
	}
	fmt.Println(brr)
	return GetLargestElementFunc(brr)
}
func GetLargestElementFunc(arr []int) int {
	largest := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
	}
	return largest
}
