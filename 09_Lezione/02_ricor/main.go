package main

import "fmt"

func f(n int) int {
	if n%2 == 1 {
		return 0
	}
	fmt.Println(n)
	return 1 + f(n/2)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(f(n))
}