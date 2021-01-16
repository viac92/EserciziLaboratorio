package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// cancella per ogni numero n >0 (intero) presente in lista,
// cancella il numero stesso e gli n elementi successivi della lista
// (o quelli che ci sono per arrivare alla fine della lista)
// e restituisce la nuova lista così prodotta;
func cancella(lista []string) []string {
	var listaOut []string
	for i := 0; i < len(lista); i++ {
		if !strings.ContainsAny(lista[i], "1234567890") {
			listaOut = append(listaOut, lista[i])
		}

		if strings.ContainsAny(lista[i], "1234567890") {
			num, _ := strconv.Atoi(lista[i])
			i += num
		}
	}
	return listaOut
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Fornire 1 nome di file")
		return
	}
	file := os.Args[1]

	contenutoFile, err := os.Open(file)
	if err != nil {
		fmt.Println("File non accessibile")
		return
	}
	defer contenutoFile.Close()

	b := bufio.NewScanner(contenutoFile)

	var riga []string
	for b.Scan() {
		riga = append(riga, b.Text())
	}

	// compatto le righe lette in un unica riga
	var stringaUnicaRiga string
	for i := 0; i < len(riga); i++ {
		stringaUnicaRiga += riga[i]
		if i != len(riga)-1 {
			stringaUnicaRiga += " "
		}
	}

	// creo una slice di di stringhe
	lista := strings.Split(stringaUnicaRiga, " ")

	// stampo
	fmt.Println(lista)
	fmt.Println(cancella(lista))
}
