// Scrivere un programma array.go dotato di:

// - una costante DIM per la dimensione di array

// - una funzione main che inizializza un array 
// di dimensione DIM e testa le due seguenti funzioni 
// che lavorano direttamente sull'array stesso

// - una funzione reverse
// che inverte l'ordine dei valori in un array di dimensione DIM,
// mettendo il primo in fondo, il secondo in penultima posizione 
// e così via nell'array stesso

// - una funzione switchFirstLast
// che scambia il primo con l'ultimo dei valori in un array
// di dimensione DIM

package main

import "fmt"

//DIM dimensione array
const DIM int = 10

func main() {
	var ar [DIM]string = [DIM]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "g"}

	ar2 := reverse(ar)
	ar3 := switchFirstLast(ar)

	for i := 0; i < DIM; i++ {
		fmt.Print(ar[i])
	}
	fmt.Println()	

	for i := 0; i < DIM; i++ {
		fmt.Print(ar2[i])
	}
	fmt.Println()

	for i := 0; i < DIM; i++ {
		fmt.Print(ar3[i])
	}
	fmt.Println()
	
}

func reverse(strarr [DIM]string) [DIM]string {
	var arout [DIM]string
	var diff int = DIM
	for i := 0; i < DIM; i++ {
		diff--
		arout[i] = strarr[diff]
	}
	return arout
}

func switchFirstLast(strarr [DIM]string) [DIM]string {
	var arout [DIM]string
	arout = strarr
	arout[0] = strarr[DIM - 1]
	arout[DIM - 1] = strarr[0]

	return arout
}