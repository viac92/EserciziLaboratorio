// Vogliamo scrivere una funzione 
// 	giorniInMese(mese int) int
// che, dato come parametro il numero corrispondente a un mese, restituisce il numero di giorni di quel mese (28 per febbraio).

package main

import "fmt"

func giorniMese(mese int) int {
	switch mese {
	case 4, 6, 9, 11:
		return 30
	case 2:
		return 28
	default:
	return 31
	}
}

func main() {
	var giorno, mese, anno int
	fmt.Print("Inserisci un numero nel formato gg-mm-aaaa: ")
	fmt.Scan(&giorno)
	fmt.Scan(&mese)
	fmt.Scan(&anno)

	fmt.Println("Il mese", mese, "ha", giorniMese(mese), "giorni.")
}