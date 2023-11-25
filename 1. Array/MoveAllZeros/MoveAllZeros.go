package main

import "fmt"

func main() {
	arr := []int{0, 10, 20, 30, 40, 50}
	MoveAllZerosFunc(arr)
	fmt.Println(arr)
}

var MoveAllZerosFunc = func(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			for j := i + 1; j < len(arr); j++ {
				if arr[j] != 0 {
					swap(&arr[i], &arr[j])
					break
				}
			}
		}
	}
}

var swap = func(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
