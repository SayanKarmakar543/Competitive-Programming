package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 20, 20, 30, 30, 30}
	brr := RemDup(arr)
	fmt.Println(brr)
}
func RemDup(arr []int) []int {
	temp := []int{} // slice declearation
	temp = append(temp, arr[0])
	for i := 1; i < len(arr); i++ {
		flag := true
		for j := i - 1; j >= 0; j-- {
			if arr[i] == arr[j] {
				flag = false
				break
			}
		}
		if flag {
			temp = append(temp, arr[i])
		}
	}
	return temp
}
