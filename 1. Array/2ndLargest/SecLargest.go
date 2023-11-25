package main

import "fmt"

func main() {
	arr := []int{60, 10, 20, 30, 40, 50, 70}
	fmt.Println(GetSecLargest(arr))
}

var GetSecLargest = func(arr []int) int {
	// TC: O(n)
	// SC: O(n)
	largest_index := GetLargestElement(arr)
	var SecLargestElement int
	if largest_index == 0 {
		SecLargestElement = arr[1]
	} else if largest_index == len(arr)-1 {
		SecLargestElement = arr[len(arr)-2]
	} else {
		SecLargestElement = arr[0]
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] > SecLargestElement && arr[i] < arr[largest_index] {
			SecLargestElement = arr[i]
		}
	}

	return SecLargestElement
}

var GetLargestElement = func(arr []int) int {
	largest_index := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[largest_index] {
			largest_index = i
		}
	}
	return largest_index
}
