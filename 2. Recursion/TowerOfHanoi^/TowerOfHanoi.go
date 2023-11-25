package main

import "fmt"

func main() {
	ToH(2, "A", "B", "C")

	// A --> Source
	// B --> Helper
	// C --> Destination
}
func ToH(n int, A string, B string, C string) {
	if n == 1 {
		fmt.Println("Disc 1 is moved from ", A, " to ", C)
		return
	}
	ToH(n-1, A, C, B)
	fmt.Println("Disc", n, "is moved from ", A, " to ", C)
	ToH(n-1, B, A, C)
}
