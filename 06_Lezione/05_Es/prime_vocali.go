// Si vuole scrivere un programma che legga una sequenza di 5 stringhe e stampi, per ogni stringa, 
// la sua prima vocale (per semplicità assumiamo che le stringhe contengano solo lettere minuscole). 
// Nel caso in cui una stringa non contenga vocali il programma stampa “la parola non contiene vocali”.
//
// Scrivi il programma prime_vocali.go completando il seguente codice. Caricalo poi su upload.
// Nota che il for per la ricerca delle vocali è annidato nel for che scorre la sequenza di stringhe
//
// const K = 5
// var parola string
//
// for i:=0; i<K; i++ { 
//    fmt.Scan(&parola)
//    for _, c := range parola {
//       ....
//    } 
// }
// ...

package main

import (
		"fmt"
		"strings"
)

func main() {
	const K = 5
	var parola string

	for i := 0; i < K; i++ {
		fmt.Print("Inserisci la ", i + 1, " parola: ")
		fmt.Scan(&parola)
		
		for i, c := range parola {
			
			if strings.IndexRune("aeiouAEIOU", c) != -1 {
				fmt.Println(string(c), "prima vocale.")
				break
			}

			if i == len(parola) - 1 {
				fmt.Println("La parola non ha vocali.")
				break
			}
		}
		fmt.Println()
	}
}