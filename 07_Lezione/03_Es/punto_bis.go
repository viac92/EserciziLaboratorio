package main

import "fmt"

// Punto è un punto nel piano cartesiano
type Punto struct {
	x, y float64
}

func main() {
	var p1, p2 Punto
	var x3, y3 float64 = 3, 4

	fmt.Print("x e y di P1: ")
	fmt.Scan(&p1.x, &p1.y)

	fmt.Print("x e y di P2: ")
	fmt.Scan(&p2.x, &p2.y)

	fmt.Println(puntoMedioSegmento(p1, p2))

	traslaPuntoP1(&p1, 2, 3)
	fmt.Println(p1)

	p3 := newPunto(x3, y3)
	fmt.Println(p3)

	specchiaPunto(&p1)
	fmt.Println(p1)
}

func puntoMedioSegmento(p1, p2 Punto) Punto {
	var p Punto
	p.x = (p1.x + p2.x) / 2
	p.y = (p1.y + p2.y) / 2
	return p
}

func traslaPuntoP1(p *Punto, horiz, vert float64) {
	(*p).x += horiz
	(*p).y += vert
	return
}

func newPunto(x, y float64) Punto {
	var p Punto
	p.x = x
	p.y = y
	return p
}

func specchiaPunto(p *Punto) {
	p.x = p.x*-1
}