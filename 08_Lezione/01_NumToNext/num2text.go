// Scrivere un programma num2text.go per convertire un numero intero non negativo nella sequenza delle parole corrispondenti alle sue cifre.package main.
// Il programma legge un intero non negativi da standard input, per ogni nuova (non incontrata finora) cifra del numero chiede il nome corrispondente (e alimenta un dizzionario), e infine
// stampa la sequenza delle parole corrispondenti alle sue cifre.
// Ad esempio, per il numero 203, il programma stampa due - zero - tre.

package main

import (
	"fmt"
)

func main() {
	var n int
	var nome, parola string
	
	m := map[int]string{}

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	for n != 0 {
		cifra := n % 10

		_,ok := m[cifra]
		if !ok {
			fmt.Printf("parola per %d ", cifra)
			fmt.Scan(&nome)
			m[cifra] = nome
		}
		
		parola = m[cifra] + " - " + parola
		n /= 10
	}

	fmt.Println(parola[:len(parola)-3])

}