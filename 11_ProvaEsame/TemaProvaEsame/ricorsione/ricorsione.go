// Scrivere una funzione ricorsiva che, dati un byte c ed una stringa s,
// restituisce la posizione della prima occorrenza di c in s, e -1 se c non
// occorre in s. E un main che legge da linea di comando i dati e testi la funzione

package main

import (
	"bufio"
	"fmt"
	"os"
)

var count = 0
func ricorsiva(c byte, s string) int {
	if len(s) == 1 && c != s[0] {
		return (-count-1) 
	}
	if c == s[0] {
		return 1
	}
	return ricorsiva(c,s[1:]) +1 
}

func main() {
	var inputString, letteraS string
	var lettera byte

	b := bufio.NewScanner(os.Stdin)

	fmt.Print("Inserisci una stringa: ")
	for b.Scan() {
		inputString = b.Text()
		break
	}

	fmt.Print("Inserisci una lettera: ")
	fmt.Scan(&letteraS)

	if len(letteraS) != 1 {
		fmt.Println("Errore - inserire una sola lettera!")
		return
	}
	lettera = letteraS[0]

	fmt.Println(ricorsiva(lettera, inputString))

	fmt.Printf("Stringa: %v, Lettera: %c\n", inputString, lettera)
}
