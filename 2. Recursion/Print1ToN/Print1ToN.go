package main

import "fmt"

func main() {
	fun(5)
}

func fun(n int) {
	if n == 0 {
		return
	}
	fun(n - 1) // non-tail recurtion
	fmt.Print(n, " ")
}
