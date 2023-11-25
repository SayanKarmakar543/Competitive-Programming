// Formula: nCr = (n-1)Cr + (n-1)C(r-1)
package main

import "fmt"

func main() {
	fmt.Println(Com(3, 1))
	fmt.Println(Com(3, 3))
	fmt.Println(Com(3, 0))
	fmt.Println(Com(3, 2))
}
func Com(n, r int) int {
	if r == n || r == 0 {
		return 1
	}
	return Com(n-1, r) + Com(n-1, r-1)
}
