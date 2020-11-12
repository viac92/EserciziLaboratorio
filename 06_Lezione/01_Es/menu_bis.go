package main

import "fmt"

func main() {

	const MENU = "Menu del giorno\n" +
		"a. pizza\n" +
		"b. penne al pomodoro\n" +
		"c. cotoletta e patawtine\n" +
		"d. crostata e caffe\n\n" + 
		"f. per terminare l'ordinazione\n"
	fmt.Println(MENU)
	
	var scelta rune
	
	for {
		for {
			fmt.Println("ordinazione?")
			fmt.Scanf("%c", &scelta)
			if scelta == 'f' {
				return
			} else if 'a' <= scelta && scelta <= 'd'  {
				break
			} 
			fmt.Scanf("%c", &scelta) //salto l'a-capo
			fmt.Println("ordinazione non valida")
			fmt.Println()
		}

		var ordinazione string
		switch scelta {
		case 'a':
			ordinazione = "pizza"
		case 'b':
			ordinazione = "penne al pomodoro"
		case 'c':
			ordinazione = "cotoletta e patatine"
		case 'd':
			ordinazione = "crostata e caffe'"
		}
		fmt.Println("hai ordinato", ordinazione)
		fmt.Println()
		fmt.Scanf("%c", &scelta)
	}
}	