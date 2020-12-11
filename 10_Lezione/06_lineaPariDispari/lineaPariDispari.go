// Estrai da un testo le righe che contengono un numero pari di rune e quelle che
// contengono un numero dispari e stampale (prima tutte le pari e poi tutte le dispari)

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var righePari, righeDispari []string

	b := bufio.NewScanner(os.Stdin)
	for b.Scan() {
		t := b.Text()
		r := []rune(t)

		if len(r) % 2 == 0 {
			righePari = append(righePari, t)
		} else {
			righeDispari = append(righeDispari, t)
		}
	}

	fmt.Printf("Linee Pari: %v\nLinee Dispari: %v\n", righePari, righeDispari)
}