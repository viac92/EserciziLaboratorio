/* ### ESERCIZIO 2a (PARI-DISPARI)
Scrivere un programma pari_dispari.go che legge un intero n e, a seconda del valore di n, 
stampa uno dei messaggi "n è pari" oppure "n è dispari". */

package main

import "fmt"

func main() {
	var n int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	if n % 2 == 0 {
		fmt.Println("Il numero è pari.")
	} else {
		fmt.Println("Il numero è dispari.")
	}
}