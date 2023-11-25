package main

import "fmt"

func main() {
	fmt.Println(fun(5))
}
func fun(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * fun(n-1)
}
