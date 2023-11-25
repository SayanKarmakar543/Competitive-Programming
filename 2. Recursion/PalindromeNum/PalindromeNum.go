package main

import "fmt"

func main() {
	num := 12321
	if fun(num, 0) == num {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("not Palindrome")
	}
}
func fun(n int, rev int) int {
	if n == 0 {
		return rev
	}
	rev = rev*10 + n%10
	return fun(n/10, rev)
}
