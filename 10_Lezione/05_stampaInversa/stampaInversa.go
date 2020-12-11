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
	"reflect"
)

func main() {
	b := bufio.NewScanner(os.Stdin)
	b.Split(bufio.ScanWords)

	var wordsSlice []string 

	for b.Scan() {
		t := b.Text()
		wordsSlice = append(wordsSlice, t)
	}

	swapSlice := reflect.Swapper(wordsSlice)
	
	for i := 0; i < len(wordsSlice)/2; i++ {
		swapSlice(i, len(wordsSlice)-1-i)
	}

	fmt.Println(wordsSlice)
}