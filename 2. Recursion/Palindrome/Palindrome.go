package main

import "fmt"

func main() {
	str := "abcba"
	fmt.Println(isPalindrom(str, 0, len(str)-1))
}
func isPalindrom(str string, low int, high int) bool {
	if low >= high {
		return true
	}
	return (str[low] == str[high] && isPalindrom(str, low+1, high-1))
}
