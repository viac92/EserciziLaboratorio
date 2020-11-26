// Anagrammi
// =========

// Scrivere un programma anagrammi.go che legge due stringhe 
// da linea di comando e valuta se le due stringhe sono 
// una l'anagramma dell'altra.

// In particolare il programma è dotato di:

// - una funzione
// 	func isAnagram(s1, s2 string) bool
// che restituisce true se le due stringhe sono una l'anagramma dell'altra,
// false altrimenti

// Domande:
// che caratteristiche ha un anagramma?
// c'è in Go una struttura di dati (array, struct, slice, map) che 
// si presta a rappresentare i dati di una stringa s1 che servono 
// ad individuare se è un anagramma di un'altra stringa s2?
// Se sì, quale e coma la uso? Se no, come imposto la soluzione?

// - una funzione
// 	func main()
// che legge due parole p1 e p2 da linea di comando e stampa uno dei due messaggi:
// p1 e p2 sono anagrammi
// p1 e p2 non sono anagrammi

// In assenza di esattamente due parametri sulla linea di comando il programma stampa:
// input errato

package main

import (
	"fmt"
	"os"
)

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	
	s1Rune := []rune(s1)
	s2Rune := []rune(s2)
	
	s1Map := 
}

func main() {
	str1 := os.Args[1]
	str2 := os.Args[2]


}