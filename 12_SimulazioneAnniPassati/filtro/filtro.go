package main

import (
	"strconv"
	"fmt"
	"os"
)

// DrawPoint restituisce una stringa formata da k spazi bianchi
// seguiti dal carattere c
func DrawPoint(c byte, k int) string {
	var point string
	for i := 0; i < k; i++ {
		point += " "
	}
	point += string(c)

	return point
}

// DrawSegment restituisce una stringa formata da k spazi bianchi
// seguiti da l caratteri c
func DrawSegment(c byte, k,l int) string {
	var segment string 
	for i := 0; i < k; i++ {
		segment += " "
	}
	for i := 0; i < l; i++ {
		segment += string(c)
	}
	return segment
}

// leggere come parametri da linea di comando (in quest'ordine)
// due numeri interi l e n e un carattere c (byte),
// e, se l e n sono > 0, produce su standard output una scala di n gradini
// di lunghezza e altezza l disegnati usando il carattere c (vedi esempi sotto),
// altrimenti non fa niente.
func main() {
	l,_ := strconv.Atoi(os.Args[1])
	n,_ := strconv.Atoi(os.Args[2])
	cS := os.Args[3]

	c := cS[0]

	if l > 0 && n > 0 {
		space := 0
		for i := 0; i < n; i++ {
			space = (l-1)*i
			fmt.Println(DrawSegment(c,space,l))
			for j := 2; j <= l; j++ {
				fmt.Println(DrawPoint(c,space+l-1))
			} 
		}
	}
}