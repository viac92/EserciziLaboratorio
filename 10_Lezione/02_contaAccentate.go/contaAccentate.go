// Dato un testo in input, conta quante lettere accentate ci sono (à, è, é, ì, ò, ù).
// Suggerimento: utilizzare (Scanner.Split(bufio.ScanRunes))

// Esempio:
// go run contaAccentate.go < file_uno
// Accenti trovati:  6

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	b := bufio.NewScanner(os.Stdin)
	b.Split(bufio.ScanRunes)
	var contaAccentate int

	for b.Scan() {
		t := b.Text()

		if strings.IndexAny(t,"àèìòùé") != -1 {
			contaAccentate++
		}
	}
	fmt.Println(contaAccentate)
}