package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 3, 2, 1, 5}
	missing, repeating := findNumber(arr)
	fmt.Println("Missing: ", missing, " & ", "Repeating: ", repeating)
}

var findNumber = func(arr []int) (int, int) {

	// TC: O(n)
	// SC: O(n)
	var (
		repeating int = -1
		missing   int = -1
	)
	m := map[int]int{}
	for i := 1; i <= len(arr); i++ {
		m[i] = 0
	}
	fmt.Println(m)
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}
	for key := range m {
		if m[key] == 2 {
			repeating = key
		}
		if m[key] == 0 {
			missing = key
		}
	}
	fmt.Println(m)
	return missing, repeating
}

func binarySearch(arr []int, element int) (int, bool) {
	var low int = 0
	var high int = len(arr) - 1
	var mid int
	for low <= high {
		var mid int = (low + high) / 2
		if mid == element {
			return mid, true
		} else if arr[mid] > element {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return mid, false
}
