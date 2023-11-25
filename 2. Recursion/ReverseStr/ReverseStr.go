package main

import "fmt"

func main() {
	str := "Sayan"
	fun(str, len(str)-1)
}
func fun(str string, n int) {
	if n < 0 {
		return
	}
	fmt.Printf("%c", str[n])
	fun(str, n-1)
}
