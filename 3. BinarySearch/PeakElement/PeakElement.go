// peak Element: arr[i-1]<arr[i]>arr[i+1]
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 8, 5, 1}
	// arr1 := []int{1, 5, 1, 2, 1}
	fmt.Println(funcPeakElement(arr))
}

var funcPeakElement = func(arr []int) int {
	low := 1
	high := len(arr) - 2
	if len(arr) == 0 {
		return arr[0]
	}
	if arr[0] > arr[1] {
		return arr[0]
	}
	if arr[len(arr)-1] > arr[len(arr)-2] {
		return arr[len(arr)-1]
	}
	for low <= high {
		mid := (low + high) >> 1
		if arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1] {
			return arr[mid]
		} else if arr[mid-1] < arr[mid] {
			low = mid + 1
		} else if arr[mid] > arr[mid+1] {
			high = mid - 1
		} else {
			high = mid - 1
			// low = mid+1
		}
	}
	return -1
}
