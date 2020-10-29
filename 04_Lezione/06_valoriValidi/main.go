//chiedi h (ore) e m (minuti) fino a ottenere valori validi (due numeri)

package main

import "fmt"

func main() {
	var h, m int
	
	fmt.Print("Inserisci l'ora attuale (nel formato hh:nn) : ")
	
	for {

		fmt.Scan(&h)
		fmt.Scan(&m)

		if h > 0 && h < 24 && m > 0 && m < 60 {
			break
		} 

		fmt.Println("Inserimenot non valido.")
		fmt.Println("Inserisci nuovamente l'orario")
	} 

	fmt.Print("L'ora attuale è: ", h, ":", m, "\n")
}