package main

import "fmt"

func main() {
	fmt.Println(powFunc(3, 3))
}
func powFunc(n, k int) int {
	if k == 0 {
		return 1
	}
	return n * powFunc(n, k-1)
}
