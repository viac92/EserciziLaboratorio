// Scrivere una funzione ricorsiva f(n) che restituisce la somma di una lista di interi
package main

import (
	"fmt"
)

func sommaRicorsiva(n []int) int {
	if len(n) == 1 {
		return n[0]
	}
	return n[0] + sommaRicorsiva(n[1:]) 
}


func main() {
	//lista := []int{1}
	//lista := []int{1, 2}
	lista := []int{4, 5 , 6, 5, 10, 20}
	fmt.Println(sommaRicorsiva(lista))
}