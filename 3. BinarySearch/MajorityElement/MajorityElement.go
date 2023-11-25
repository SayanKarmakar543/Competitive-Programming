package main

import "fmt"

func main() {

	// TC: O(n^2)
	// SC: O(n)

	arr := []int{3, 1, 2, 2, 3, 3, 2, 2, 1, 1, 1, 1, 1}
	fmt.Println(findMaorityElement(arr))
}
func findMaorityElement(arr []int) int {
	m := map[int]int{}
	for i := 0; i < len(arr); i++ {
		count := 0
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				count++
			}
		}
		m[arr[i]] = count
	}
	return findMaxElement(m)
}
func findMaxElement(m map[int]int) int {

	var largest int = 0
	var ans int
	for key, value := range m {
		if value > largest {
			ans = key
			largest = value
		}
	}
	return ans
}
