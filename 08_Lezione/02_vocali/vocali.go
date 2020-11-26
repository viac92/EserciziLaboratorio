// Vocali
// ======
// #mappe - conteggio

// Scrivere un programma vocali.go che analizza un testo e conta
// le occorrenze delle vocali (sia minuscole che maiuscole, 
// ma non le accentate) nel testo e stampa, ma solo per le vocali
// presenti nel testo, il numero di volte che le vocali stesse sono 
// presenti nel testo.

// In particolare il programma è dotato di:

// - una funzione
// 	func contaVocali(s string, vocali map[rune]int)
// per contare le occorrenze delle diverse vocali
// (sia minuscole che maiuscole - vedi es sotto)
// in tutte le stringhe che le vengono passate.
// La funzione, data una stringa s e una mappa vocali, aggiorna 
// opportunamente la mappa vocali aggiungendo 
// eventuali vocali non ancora presenti / incrementandone i valori.
// Per individuare le vocali e aggiornare la variabile vocali
// usate uno switch con un solo case.

package main

import (
	"fmt"
	"bufio"
	"os"
)

func contaVocali(s string, vocali map[rune]int) {
	for _, k := range s {
		switch k {
		case 'a','A','e','E','i','I','o','O','u','U':
			vocali[k]++
		}
	}
}

func main() {
	var frase string
	//var vocaliPresenti []string
	vocali := map[rune]int{}

	b := bufio.NewScanner(os.Stdin)

	fmt.Print("Inserisci una stringa: ")
	b.Scan()
	frase = b.Text()

	contaVocali(frase, vocali)
	
	// for k, v := range vocali {
	// 	vocaliPresenti = append(vocaliPresenti, string(k))
	// 	fmt.Printf("%c ---> %d\n", k, v)
	// }

	for _,r := range "aeiouAEIOU" {
		_, ok := vocali[r]
		if ok {
		  fmt.Printf("%c ---> %d\n", r, vocali[r]) 
		}
	}	
}