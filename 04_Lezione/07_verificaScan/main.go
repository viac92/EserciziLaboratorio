package main

import "fmt"

func main() {
	var n int
	//var numero int
	//var errore error
	fmt.Print("inserisci intero: ")
	numero,errore := fmt.Scan(&n)
	fmt.Println(numero)
	fmt.Println(errore)
}