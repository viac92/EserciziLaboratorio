// Scrivere una funzione ricorsiva f(n) che restituisce la successione di Collatz generata a partire da n>0, cioè la sequenza di numeri generati secondo la funzione f:
// f(1) =   1
// f(n) =   n/2         se n è pari
// 	     n*3+1     se n è dispari

package main

import "fmt"

func successioneCollatz(n int) *[]int {
	
}

func main() {
	n := 1
	fmt.Println(successioneCollatz(n))
}