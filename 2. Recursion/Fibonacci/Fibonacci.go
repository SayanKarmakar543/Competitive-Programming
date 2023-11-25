package main

import "fmt"

func main() {
	fmt.Println(fun(4))
}
func fun(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fun(n-1) + fun(n-2)
}
