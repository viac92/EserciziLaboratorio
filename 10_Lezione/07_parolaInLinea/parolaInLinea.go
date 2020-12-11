// Estrai da un testo ogni riga che contiene una certa parola (simil grep) e stampala

// Esempio:
// //parola : ti
// go run parolaInLinea.go < file_uno 
// [tienti col corno, e con quel ti disfoga  quand’ira o altra passion ti tocca!"]

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	b := bufio.NewScanner(os.Stdin)
	var rigaConTarget []string

	for b.Scan() {
		t := b.Text()
		target := "ti"
		
		if strings.Contains(t, target) {
			rigaConTarget = append(rigaConTarget, t)
		} 
	}

	for _,s := range rigaConTarget{
		fmt.Println(s)
	}
}