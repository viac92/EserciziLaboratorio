// Scrivi una funzione operazioni(n1, n2 int) (int, int, int) che accetta due interi e restituisce somma, prodotto e differenza.
// Scrivine una versione con variabili di ritorno nominate e una senza.
//
// Scrivi un main per invocare e testare la funzione. Il programma legge da standard input due int.

package main

import "fmt"

func operazioni(n1, n2 int) (somma, prodotto, differenza int) {
	return n1 + n2, n1 * n2, n1 - n2
}

func main() {
	var n1, n2 int
	fmt.Print("Inserisci il primo numero: ")
	fmt.Scan(&n1)
	fmt.Print("Inserisci il secondo numero: ")
	fmt.Scan(&n2)

	somma, prodotto, differenza := operazioni(n1, n2)

	fmt.Println("Somma:", somma, "\nProdotto:", prodotto, "\nDifferenza:", differenza)
}