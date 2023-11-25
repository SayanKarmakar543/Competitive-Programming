package main

import "fmt"

func main() {
	fmt.Println(fun(1234))
}
func fun(n int) int {
	if n == 0 {
		return n % 10
	}
	return n%10 + fun(n/10)
}
