//leggi K numeri e di ciascuno stampa il doppio

package main

import "fmt"

func main() {
	const K = 5 // "hardcoded", numero fisso di iterazioni

	var n int

	for i := 1; i <= K; i++ {
		fmt.Print("Inserisci un numero: ")
		fmt.Scan(&n)

		fmt.Println(n * 2, " ")
	}
}