package main

import "fmt"

func main() {
	num := 12345
	fmt.Println(fun(num, 0))
	fun1(num)
}
func fun1(n int) {

	if n == 0 {
		return
	}
	fmt.Print(n % 10)
	fun1(n / 10)
}
func fun(n int, rev int) int {
	if n == 0 {
		return rev
	}
	return fun(n/10, (rev*10)+(n%10))
}

// Hypothesis/Intution:-
// fun(12, 0) return 21
// 	|			^
// 	|			|
// fun(1, 2) return 21
// 	|			^
// 	|			|
// fun(0, 21) return 21
