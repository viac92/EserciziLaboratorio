//leggi K numeri e stampa la loro somma

package main

import "fmt"

func main() {
	const K int = 5

	var n int 
	var somma int

	for i := 1; i <= K; i++ {
		fmt.Print("Inserisci il ", i, "° numero: ")
		fmt.Scan(&n)

		somma += n
	}

	fmt.Println("La somma di tutti i numeri è", somma)
}

// si può creare un file di inputo da inserire nel programma go buildato 
// nel terminale $ ./programma < fileInput