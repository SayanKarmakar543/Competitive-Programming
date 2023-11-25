package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50}
	rotateArr(arr, 2)
	fmt.Println(arr)
}
func rotateArr(arr []int, k int) {
	reverseArr(arr, 0, k-1)
	reverseArr(arr, k, len(arr)-1)
	reverseArr(arr, 0, len(arr)-1)
}

var reverseArr = func(arr []int, low int, high int) {
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
