package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50}
	ReverseFunc(arr)
	fmt.Println(arr)
}

var ReverseFunc = func(arr []int) {
	var (
		low  int = 0
		high int = len(arr) - 1
	)
	for low < high {
		swap(&arr[low], &arr[high])
		low++
		high--
	}
}
var swap = func(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
