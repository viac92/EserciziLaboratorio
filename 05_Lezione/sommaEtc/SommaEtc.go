// creare un programma go chiamato SommaEtc.go che chieda all'utente una serie di numeri (float) e calcoli somma, media, min, max

package main

import (
		"fmt"
		"io"
)

func main() {
	var f, somma, i, min, max float64
	fmt.Print("Inserisci dei numeri calcolo la somma, media e min, max: ")

	for {
		_, err := fmt.Scan(&f) 
		if err == io.EOF {
			break
		}
		if f > max {
			max = f
		}
		if f < min || i == 0 {
			min = f
		}
		somma += f
		i++
	}
	fmt.Println("La somma è:", somma)
	fmt.Println("La media è:", somma/i)
	fmt.Println("Il numero più grande inserito è:", max)
	fmt.Println("Il numero più piccolo è:", min)
}
