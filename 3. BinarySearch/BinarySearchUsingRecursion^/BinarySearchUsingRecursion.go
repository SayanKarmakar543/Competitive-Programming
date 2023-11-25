package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50}
	x := 20
	fmt.Println(BinarySearch(arr, 0, len(arr)-1, x))
}
func BinarySearch(arr []int, low int, high int, x int) int {
	if low > high {
		return -1
	}
	var mid int = (low + high) / 2
	if arr[mid] == x {
		return mid
	} else if arr[mid] > x {
		return BinarySearch(arr, low, mid-1, x)
	} else if arr[mid] < x {
		return BinarySearch(arr, mid+1, high, x)
	}
	return -1
}
