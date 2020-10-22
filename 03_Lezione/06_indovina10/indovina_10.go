/* ### ESERCIZIO 3a (INDOVINA 10)
Scrivere un programma indovina_10.go che fissa un numero intero tra 1 e 10 da indovinare, 
legge un intero da standard input e
- se il numero in input è fuori dall’intervallo 1-10, stampa “Valore non valido”;
- se il numero è quello fissato, stampa “Hai indovinato!”; 
- altrimenti stampa “Non hai indovinato”. */

package main

import ("fmt"
		"math/rand"
		"time"
)

func main() {
	var n int
	v := rand.NewSource(time.Now().UnixNano())
	r := rand.New(v)

	c := r.Intn(10)
	var vinto bool

	for !vinto {
		fmt.Print("Inserisci un numero da 1 a 10 prova ad indovinare: ")
		fmt.Scan(&n)	
		if n >= 1 && n <= 10 {
			if n == c {
				fmt.Println("Hai indovinato!")
				vinto = true 
			} else {
				fmt.Println("Hai perso riprova.")
			}
		} else {
			fmt.Println("Valore non valido.")
		}
	}
}