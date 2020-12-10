// Dato un testo in input, conta quante volte ciascun carattere di punteggiatura 
// (. , ; : ! ? " ') viene ripetuto.
// Suggerimento: utilizzare (Scanner.Split(bufio.ScanBytes))

// Esempio:
// go run contaPunteggiatura.go < file_uno
// map[":4 ,:4 .:1 ::1 !:1]

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	b := bufio.NewScanner(os.Stdin)
	b.Split(bufio.ScanBytes)
	mappa := make(map[string]int)

	for b.Scan() {
		t := b.Text()
		if strings.IndexAny(t, ".,;:!?\"'") != -1 {
			mappa[t]++
		}
	}

	fmt.Println(mappa)
}