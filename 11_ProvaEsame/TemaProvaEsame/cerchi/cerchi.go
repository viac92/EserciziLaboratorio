package main

import (
	"math"
	"strconv"
	"strings"
	"bufio"
	"fmt"
	"os"
)

//Cerchio è una struttura per gestire cerchi geometrici
type Cerchio struct {
	nome string
	x,y,raggio float64
}

func newCircle(descr string) Cerchio {
	var cerchioOut Cerchio
	parametri := strings.Split(descr, " ")
	cerchioOut.nome = parametri[0]
	cerchioOut.x,_  = strconv.ParseFloat(parametri[1],64)
	cerchioOut.y,_  = strconv.ParseFloat(parametri[2],64)
	cerchioOut.raggio,_  = strconv.ParseFloat(parametri[3],64)
	return cerchioOut
}

func distanzaPunti(x1, y1, x2, y2 float64) float64 {
	distanza :=math.Sqrt(math.Pow((x2 - x1),2) + math.Pow((y2 -y1),2))
	return distanza
}

func tocca(cerc1, cerc2 Cerchio) bool {
	var cerchioMaggiore, cerchioMinore Cerchio

	distanza := distanzaPunti(cerc1.x, cerc1.y, cerc2.x, cerc2.y)
	
	if maggiore(cerc1, cerc2) {
		cerchioMaggiore = cerc1
		cerchioMinore = cerc2
	} else {
		cerchioMaggiore = cerc2
		cerchioMinore = cerc1
	}
	
	if distanza == cerc1.raggio + cerc2.raggio || distanza + cerchioMinore.raggio == cerchioMaggiore.raggio {
		return true
	}
	return false
}

func maggiore(cerc1, cerc2 Cerchio) bool {
	if cerc1.raggio > cerc2.raggio {
		return true
	}
	return false
}

func main()  {
	var inserimenti []string
	var cerchi []Cerchio
	
	b := bufio.NewScanner(os.Stdin)
	for b.Scan() {
		inserimenti = append(inserimenti, b.Text())
	}

	for _,righe := range inserimenti {
		cerchi = append(cerchi, newCircle(righe))
	}

	fmt.Println(cerchi[0], "non tangente, maggiore")
	for i := 1; i < len(cerchi); i++ {
		
		var isTangente string
		if tocca(cerchi[i], cerchi[i-1]) {
			isTangente = "tangente,"
		} else {
			isTangente = "non tangente,"
		}

		var isMaggiore string
		if maggiore(cerchi[i], cerchi[i-1]) {
			isMaggiore = "maggiore"
		} else {
			isMaggiore = "minore o uguale"
		}
		fmt.Printf("%v %v %v\n", cerchi[i], isTangente, isMaggiore)	
	}
}