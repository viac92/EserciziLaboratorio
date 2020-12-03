package main

import "fmt"

func main() {
	lista := []int{1, 5, 2, 9, 7, 12, 4}
	fmt.Println(maxRicorsiva(lista))
}

func maxRicorsiva(lista []int) int {
	if len(lista) == 1 {
		return lista[0]
	}
	return max(lista[0], maxRicorsiva(lista[1:]))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}