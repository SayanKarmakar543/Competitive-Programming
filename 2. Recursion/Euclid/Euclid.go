// gcd(a, b) = gcd (b, a%b)
// gcd(a, 0) = a

package main

import "fmt"

func main() {
	fmt.Println(gcd(14, 10))
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
