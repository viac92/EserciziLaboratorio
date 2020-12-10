// Dato un testo in input, estrai e stampa prima le righe pari e poi le dispari. 
// La prima riga è la riga uno.

package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	b := bufio.NewScanner(os.Stdin)
	var count int
	var sliceRighePari, sliceRigheDispari []string

	for b.Scan() {
		t := b.Text()
		count++
		if count % 2 == 0 {
			sliceRighePari = append(sliceRighePari, t)
		}
		if count % 2 == 1 {
			sliceRigheDispari = append(sliceRigheDispari, t)
		}
	}

	for _,str := range sliceRighePari {
		fmt.Println(str)
	}
	for _,str := range sliceRigheDispari {
		fmt.Println(str)
	}
}