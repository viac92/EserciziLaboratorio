// Scrivere una funzione ricorsiva che, data una slice di rune s ed un carattere c, restituisca il numero di occorrenze di c in s.

package main

import (
	"fmt"
)

func trovaRuna(sliR []rune, r rune) int {
	var count int
	if sliR[0] == r {
		count++
	}
	if len(sliR) == 0 {
		return count
	}
	if len(sliR) == 1 {
		return count
	}	
	fmt.Println(sliR[:1])
	return count + trovaRuna(sliR[:1], r)
}

func main() {
	lista := []rune{'a','b','c','c'}

	fmt.Println(trovaRuna(lista, 'c'))
}