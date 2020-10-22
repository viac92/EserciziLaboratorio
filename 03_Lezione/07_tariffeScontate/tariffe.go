/*  ESERCIZIO 3c (TARIFFE SCONTATE)
Scrivere un programma tariffe.go che chiede età (int) e se studente (bool) e stampa il costo del 
biglietto d'ingresso secondo la seguente tabella:
[0 -9) anni: gratis
[9 -18) anni: 5
[18 - 26) anni: 7.5
[26 - 65) anni: 10
>= 65 anni: 7.5
studenti >= 18: 5 */

package main

import "fmt"

func main() {
	var e int
	var s bool

	fmt.Print("Inserisci la tua età: ")
	fmt.Scan(&e)
	fmt.Print("Inserisci t se sei studente o f se non studi: ")
	fmt.Scan(&s)

	if e >= 0 && e < 9 {
		fmt.Println("gratis")
	} else if e >= 9 && e < 18 {
		fmt.Println("7.5")
	} else if s == true {
		fmt.Println("5")
	} else if e >= 18 && e < 65 {
		fmt.Println("10")
	} else if e >=  65 {
		fmt.Println("7.5")
	}
}