// Scrivere una funzione ricorsiva che, avendo in input una lista di interi, dia in output TRUE se 10 è un elemento della lista, 
// FALSE altrimenti.

package main

import (
	"fmt"
)

func trovaDieci(n []int) bool {
	if n[0] == 10 {
		return true
	}
	if len(n) == 1 {
		return false
	}	
	return trovaDieci(n[1:])
}

func main() {
	lista := []int{1, 2, 3, 9, 20, 40, 10}
	fmt.Println(trovaDieci(lista))
}