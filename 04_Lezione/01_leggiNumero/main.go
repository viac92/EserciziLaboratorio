// leggi un numero n (da stdin) e stampa n asterischi (con o senza newline?)

package main

import "fmt"

func main() {
	// dichiarazione variabili varie di utilizzo
	var n int

	// richiesta all'utente (o accettazione parametri e/o lettura da file)
	fmt.Print("Immetti un numero: ") // prompting
	fmt.Scan(&n) // lettura da stdin

	// iterazione
	for 
		i := 1;	i <= n;	i++	{	
		fmt.Print("*")
	}
	fmt.Println()
}