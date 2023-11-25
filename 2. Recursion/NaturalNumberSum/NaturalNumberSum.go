package main

import "fmt"

func main() {
	fmt.Println(NaturalNumberSumFunc(4))
}

func NaturalNumberSumFunc(n int) int {
	if n == 1 {
		return 1
	}
	return n + NaturalNumberSumFunc(n-1)
}
