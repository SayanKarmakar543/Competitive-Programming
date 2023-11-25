package main

import "fmt"

func main() {
	subset("abc", "", 0)
}
func subset(str string, curr string, index int) {
	if len(str) == index {
		fmt.Println(curr)
		return
	}
	subset(str, curr, index+1)
	subset(str, curr+string(str[index]), index+1)
}
