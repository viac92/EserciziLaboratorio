// Si vuole scrivere un programma che legga una sequenza di 10 interi positivi 
// e stampi il massimo intero letto e quante volte il massimo compare nella sequenza.
// Completa il programma seguente inserendo, nell’ordine corretto, le istruzioni che trovi più sotto. Chiama il programma conta_massimi.go e caricalo su upload.
// Nota che la soluzione, invece di trovare prima il massimo e poi contare quante volte compare, utilizza un solo cicli for, 
// fondendo quindi le soluzioni ai due sottoproblemi di trovare il massimo e di contare quante volte compare.

// --- codice da completare  con le istruzioni che trovi qui sotto ----
// package main
// import "fmt"
// func main() {
//    var numero int 
//    fmt.Scan(&numero) 
//    max := numero
//    ...
//    fmt.Println(max)
//    fmt.Println(count) 
// }

// --- istruzioni da inserire per completare il programma qui sopra ---
// }
// if numero == max {
// for i := 1; i<10; i++ { 
// count ++
// max = numero
// count = 1
// count := 1 
// fmt.Scan(&numero) 
// } else if numero > max{ 
// }

package main

import "fmt"

func main() {
	var numero int 
	max := numero
	count := 1
	
	for i := 1; i < 10; i++ {
		fmt.Scan(&numero) 
		
		if numero == max {
			count++
		} else if numero > max {
			max = numero
			count = 1
		}
	}
	fmt.Println(count) 
	fmt.Println(max)
}
	
