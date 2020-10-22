/* ### ESERCIZIO 1a

Scrivere un programma Go 'voto_valido.go' che legge un numero intero.
Se il numero non è compreso tra 0 e 30, stampa “voto non valido”, altrimenti non stampa niente. */

package main

import "fmt"

func main() {
	var v int
	fmt.Print("Inserisci un voto da 0 a 30: ")
	fmt.Scan(&v)
	if v >= 0 && v <= 30 {
		fmt.Println("Voto valido.", v)
	} else {
		fmt.Println(v, "non è consentito.")
	}
}