package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50}
	var n int
	fmt.Print("Enter the searching element: ")
	fmt.Scan(&n)
	var pos = SearchFunc(arr, n)
	if pos != -1 {
		fmt.Println(pos)
	}
}

var SearchFunc = func(arr []int, n int) int {
	for index, value := range arr {
		if value == n {
			return index
		}
	}
	return -1
}
