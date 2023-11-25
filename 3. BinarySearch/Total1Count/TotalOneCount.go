// Given a sorted non-increasing binary array
package main

import "fmt"

func main() {
	arr := []int{1, 1, 1, 0, 0, 0}
	fmt.Println(countOnes(arr))

	// TC = O(log n)
	// SC = O(1)

}
func countOnes(arr []int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		var mid int = (low + high) / 2
		// check mid index holds 0
		if arr[mid] == 0 {
			// []int{0}  -----> this condition comes first o/w code will be crushed (Array index error)
			// []int{1, 1, 0, 0, 0}
			if mid == 0 || arr[mid-1] == 0 {
				return mid
			} else { // []int{0, 0, 0, 0, 0} or []int{1, 1 , 0, 0, 0, 0, 0, 0, 0} so on...
				high = mid - 1
			}
		} else { // check mid index doesn't holds 0
			// []int{1}
			// []int{1, 1, 1, 0, 0}
			if mid == len(arr)-1 || arr[mid+1] == 0 {
				return mid + 1
			} else { // []int{1, 1, 1, 1, 1} or []int{1, 1, 1, 1, 1, 0} so on...
				low = mid + 1
			}
		}
	}
	return -1 // if the array is empty
}
