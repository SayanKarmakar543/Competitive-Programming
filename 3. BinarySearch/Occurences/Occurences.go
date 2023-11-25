package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 2, 2, 2, 3}
	x := 3
	fmt.Println(findOccurence(arr, x))
}

func findOccurence(arr []int, x int) int {
	m := map[int]int{}
	for i := 1; i <= len(arr); i++ {
		m[i] = 0
	}
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}
	val, ok := m[x]
	// If the key exists
	if ok {
		return val
	}
	return 0
}
