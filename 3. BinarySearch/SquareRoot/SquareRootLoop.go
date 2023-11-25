package main

import "fmt"

func main() {
	fmt.Println(SqRoot(18))
}
func SqRoot(n int) int {
	// for i := 1; i*i < n; i++ {
	// 	if i*i == n {
	// 		return i
	// 	} else if i*i < n && (i+1)*(i+1) > n {
	// 		return i
	// 	}
	// }

	// TC: T(sqrt(n))
	var ans int
	for i := 1; i*i <= n; i++ {
		ans = i
	}
	return ans
}
