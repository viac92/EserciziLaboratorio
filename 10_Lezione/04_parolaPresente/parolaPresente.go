// Verifica se una data parola (già codificata all'interno del codice) è presente 
// (come parola o anche come sottostringa di una parola) nel testo fornito in input. 
// Nel caso in cui lo sia, stampa in che posizione si trova la parola (o la sua 
// sovrastringa), altrimenti stampa "not found". La prima parola si trova in posizione 1.
// Suggerimento: (Scanner.Split(bufio.ScanWords)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	b := bufio.NewScanner(os.Stdin)
	b.Split(bufio.ScanWords)
	var count int

	for b.Scan() {
		t := b.Text()
		count++

		if strings.Index(t, "duca") != -1 {
			fmt.Printf("found at: %d\n", count)
			return
		}
	}
	fmt.Println("not found")

}