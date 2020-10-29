/* Scrivere un programma Go condizioni.go che testa, una a una, le seguenti condizioni.
Implementarne una alla volta, testarla su almeno due input (uno che la verifica
e uno che la falsifica), e solo poi procedere alla successiva.
Il programma legge un valore (del tipo indicato) da tastiera e stampa “true” o “false” a seconda che la condozione sia verificata o no.

1 a int uguale a 10
2 b int diverso da 10
3 c int diverso da 10 e da 20
4 d int diverso da 10 o da 20
5 e int maggiore o uguale a 10
6 f int compreso tra 10 e 20, nell’intervallo [10,20]
7 g int compreso tra 10 e 20, nell’intervallo (10,20]
8 h int compreso tra 10 e 20, nell’intervallo [10,20)
9 i int esterno all’intervallo [10,20]
10 j int tra 10 e 20 (nell’intervallo [10,20]) o tra 30 e 40 ([30,40])
11 k int multiplo di 4 ma non di 100
12 l int dispari e compreso tra 0 e 100 ([0,100])
13 x float64 “vicino” a 10 (entro 10^-6)
*/

package main

import "fmt"

func main() {

	// 1a int uguale a 10
	var n int 
	var verifica bool
	fmt.Println("Condizione int uguale a 10")
	fmt.Print("Inserisci il numero: ")
	fmt.Scan(&n)

	if n == 10 {
		verifica = true
	}
	fmt.Println(verifica)

	// 2b int diverso da 10
	verifica = false
	fmt.Println("Condizione int diverso da 10")
	fmt.Print("Inserisci il numero: ")
	fmt.Scan(&n)
	
	if n != 10 {
		verifica = true
	}
	fmt.Println(verifica)

	// 3c int diverso da 10 e da 20
	verifica = false
	fmt.Println("Condizione diverso da 10 e da 20")
	fmt.Print("Inserisci il numero: ")
	fmt.Scan(&n)

	if n != 10 && n != 20{
		verifica = true
	}
	fmt.Println(verifica)

	// 4d int diverso da 10 o da 20
	verifica = false
	fmt.Println("Condizione diverso da 10 o da 20")
	fmt.Print("Inserisci il numero: ")
	fmt.Scan(&n)

	if n != 10 || n != 20{
		verifica = true
	}
	fmt.Println(verifica)

	// 5e int maggiore o uguale a 10
	verifica = false
	fmt.Println("Condizione maggiore o uguale a 10")
	fmt.Print("Inserisci il numero: ")
	fmt.Scan(&n)

	if n >= 10 {
		verifica = true
	}
	fmt.Println(verifica)

	// 6f int compreso tra 10 e 20, nell’intervallo [10,20]
	verifica = false
	fmt.Println("Condizione compreso tra 10 e 20, nell'intervallo [10, 20]")
	fmt.Print("Inserisci il numero: ")
	fmt.Scan(&n)

	if n >= 10 && n <= 20 {
		verifica = true
	}
	fmt.Println(verifica)

	// 7g int compreso tra 10 e 20, nell’intervallo (10,20]
	verifica = false
	fmt.Println("Condizione compreso tra 10 e 20, nell'intervallo (10, 20)")
	fmt.Print("Inserisci il numero: ")
	fmt.Scan(&n)

	if n > 10 && n < 20 {
		verifica = true
	}
	fmt.Println(verifica)

	// 8h int compreso tra 10 e 20, nell’intervallo [10,20)
	verifica = false
	fmt.Println("Condizione compreso tra 10 e 20, nell'intervallo [10, 20)")
	fmt.Print("Inserisci il numero: ")
	fmt.Scan(&n)

	if n >= 10 && n < 20 {
		verifica = true
	}
	fmt.Println(verifica)
	
	// 9g int esterno all’intervallo [10,20]
	verifica = false
	fmt.Println("Condizione esterno all’intervallo [10,20]")
	fmt.Print("Inserisci il numero: ")
	fmt.Scan(&n)

	if n <= 10 || n >= 20 {
		verifica = true
	}
	fmt.Println(verifica)

	// 10h int tra 10 e 20 (nell’intervallo [10,20]) o tra 30 e 40 ([30,40])
	verifica = false
	fmt.Println("Condizione esterno all’intervallo [10,20]")
	fmt.Print("Inserisci il numero: ")
	fmt.Scan(&n)

	if (n >= 10 && n <= 20) || (n >= 30 && n <= 40) {
		verifica = true
	}
	fmt.Println(verifica)

	// 11k int multiplo di 4 ma non di 100	
	verifica = false
	fmt.Println("Condizione multiplo di 4 ma non di 100")
	fmt.Print("Inserisci il numero: ")
	fmt.Scan(&n)

	if n % 4 == 0 && n % 100 != 0 {
		verifica = true
	}
	fmt.Println(verifica)

	// 12l int dispari e compreso tra 0 e 100 ([0,100])
	verifica = false
	fmt.Println("Condizione dispari e compreso tra 0 e 100 ([0,100]")
	fmt.Print("Inserisci il numero: ")
	fmt.Scan(&n)

	if n % 2 == 0 && n >= 0 && n <= 100 {
		verifica = true
	}
	fmt.Println(verifica)

	// 13x float64 “vicino” a 10 (entro 10^-6)
	verifica = false
	var f float64
	fmt.Println("Condizione float64 “vicino” a 10 (entro 10^-6)")
	fmt.Print("Inserisci il numero: ")
	fmt.Scan(&f)

	if f >= 1E-6 {
		verifica = true
	}
	fmt.Println(verifica)

}


