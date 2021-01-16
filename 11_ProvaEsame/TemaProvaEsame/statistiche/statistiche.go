package main

import (
	"sort"
	"strconv"
	"os"
	"bufio"
	"fmt"
)

func moda(serie []float64) float64 {
	sort.Float64s(serie)
	ricercaModa := make(map[float64]int)
	var sliceMode []float64

	for _,numero := range serie {
		ricercaModa[numero]++
	}

	ricorrenza := 0
	for _,value := range ricercaModa {
		if value > ricorrenza {
			ricorrenza = value
		}
	}
	for key,value := range ricercaModa {
		if value == ricorrenza {
			sliceMode = append(sliceMode, key)
		}
	}

	sort.Float64s(sliceMode)
	return sliceMode[0]
}

func mediana(serie []float64) float64 {
	sort.Float64s(serie)
	var ricercaMediana float64
	if len(serie) % 2 == 1 {
		ricercaMediana = serie[len(serie)/2]
	} else {
		ricercaMediana = (serie[(len(serie)/2)-1] + serie[len(serie)/2])/2
	}
	return ricercaMediana
}

func main()  {
	var serieNumeri []float64
	b := bufio.NewScanner(os.Stdin)
	
	for b.Scan() {
		t := b.Text()
		f,_ := strconv.ParseFloat(t,64)
		serieNumeri = append(serieNumeri, f)
	}
	
	fmt.Println("moda:", moda(serieNumeri))
	fmt.Println("mediana:", mediana(serieNumeri))
}