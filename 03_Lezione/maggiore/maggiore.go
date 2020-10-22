/* ### ESERCIZIO 1b (MAGGIORE) 

Scrivere un programma Go maggiore.go che legge due interi e stampa il maggiore fra i due. */

package main

import "fmt"

func main() {
	var n, m int
	fmt.Println("Inserire due numeri: ")
	fmt.Print("Inserisci il primo numero: ")
	fmt.Scan(&n)
	fmt.Print("Inserisci il secondo numero: ")
	fmt.Scan(&m)

	if n > m {
		fmt.Println(n)
	} else  {
		fmt.Println(m)
	}
}
