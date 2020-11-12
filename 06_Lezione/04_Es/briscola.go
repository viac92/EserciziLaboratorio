// Scrivi una funzione 
// punti(s string) int 
// che accetta una stringa di 3 caratteri che rappresenta una mano di tre carte e restituisce il punteggio complessivo relativo per il gioco della briscola. 
// Ad esempio per la mano "53J" restituisce 12 (10 della carta 3 + 2 del fante). 
// Se la stringa contiene un simbolo che non corrisponde a nessuna carta, la funzione restituisce -1.
//
//
// Punti a briscola:	
// Asso (A): 11	
// 3: 10	
// Re (K): 4	
// Donna (Q): 3	
// Fante (J): 2
// 7, 6, 5, 4, 2: 0

package main

import "fmt"

func punti(s string) int {
	var punti int
	for _, char := range s {

		switch char {
		case 'A':
			punti += 11
		case '3':
			punti += 10
		case 'K':
			punti += 4
		case 'Q':
			punti += 3
		case 'J': 
			punti += 2
		}
	}
	return punti
}

func main() {
	var mano string
	
	fmt.Print("Inserisci una mano di briscola: ")
	
	for {
		fmt.Scan(&mano)
		if len(mano) == 3 {
			break
		}
		fmt.Print("Inserisci una stringa valida: ")
	}

	fmt.Print("mano ", mano,": ", punti(mano), " punti\n")
}