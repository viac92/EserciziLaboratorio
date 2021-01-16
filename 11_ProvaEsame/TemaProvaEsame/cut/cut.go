package main

import (
	"strconv"
	"bufio"
	"os"
	"fmt"
)

func main()  {
	nomeFile := os.Args[3]
	taglioI,_ := strconv.Atoi(os.Args[1])
	taglioJ,_ := strconv.Atoi(os.Args[2])

	file,err := os.Open(nomeFile) 
	if err != nil {
		fmt.Println("Errore")
		return
	}
	defer file.Close()

	b := bufio.NewScanner(file)

	var cut string
	for b.Scan() {
		t := b.Text()
		if taglioI <= len(t) &&  taglioJ <= len(t) {
			cut = t[taglioI-1:taglioJ]
		} else if taglioJ >= len(t) && taglioI <= len(t) {
			cut = t[taglioI-1:]
		} else {
			cut = ""
		}
		fmt.Println(cut)
	}
}