// Dato un testo in input, stamparlo invertendo l'ordine delle parole 
// (stampando l'ultima per prima, la penultima per seconda, ecc) 
// Suggerimento: utilizza (myScanner.Split(bufio.ScanWords))

// Esempio:
// go run stampaInversa.go < file_due 
// tavolo sul è gatto il  

package  main

import  (
	"bufio"
	"fmt"
	"os"
)

func swap(s []string) []string0 {
	if len(s) == 1 {
		return s[0]
	}
}

func main() {
	b := bufio.NewScanner(os.Stdin)
	b.Split(bufio.ScanWords)

	var wordsSlice []string 

	for b.Scan() {
		t := b.Text()
		wordsSlice = append(wordsSlice, t)
	}
}