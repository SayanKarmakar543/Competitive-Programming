package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50}
	x := 20
	fmt.Println(BinarySearch(arr, 0, len(arr)-1, x))
}
func BinarySearch(arr []int, low int, high int, x int) int {
	for low <= high {
		var mid int = (low + high) / 2
		if arr[mid] == x {
			return mid
		} else if arr[mid] > x {
			high = mid - 1
		} else if arr[mid] < x {
			low = mid + 1
		}
	}
	return -1
}
