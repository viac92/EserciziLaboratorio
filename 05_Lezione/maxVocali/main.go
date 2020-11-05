// calcolare il massimo numero di vocali in una serie di stringhe (parole?)

// connessione a atrent.it:11778


package main

import (
	"fmt"
	"io"
	"strings"
)

func main(){
	var parola string
	var max int
	for {
		_, err := fmt.Scan(&parola)
		if err == io.EOF {
			break
		}

		fmt.Print(parola, ": ")
		count := 0
		for _, char := range parola {
			// fmt.Println(string(char))
			if strings.IndexRune("aeiouAEIOU", char) != -1 {
				count++
			}
		}
		fmt.Println(count, " vocali")
		if count > max {
			max = count
		}		
	}
	fmt.Println("Il massimo numero di vocali è stato: ", max)
}