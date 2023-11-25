package main

import "fmt"

func main() {
	fmt.Println(fun(1234, 0))
}
func fun(n int, count int) int {
	if n == 0 {
		return count
	}
	return fun(n/10, count+1)
}

// Hypothesis/Intution:-
// fun(12, 0) return 2
// 	|			^
// 	|			|
// fun(1, 1) return 2
// 	|			^
// 	|			|
// fun(0, 2) return 2
