// Scrivere una funzione ricorsiva che, data una slice di rune s ed un carattere c, restituisca il numero di occorrenze di c in s.

package main

import (
	"fmt"
)

func trovaRuna(s []rune, c rune) int {
	var count int
	if len(s) == 1 && s[0] == c {
		return count + 1
	}
	if len(s) == 1 && s[0] != c {
		return count
	}
	
	if s[0] == c {
		count++
	}

	return count + trovaRuna(s[1:], c)
}

func main() {
	lista := []rune{'c', 'a', 'c','x', 'c', 'c'}
	fmt.Println(lista)
	fmt.Println(trovaRuna(lista, 'x'))
}