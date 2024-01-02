package interfaces

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64
}

// ? Uma struct implementa uma interface de forma implicita, ou seja, basta implementar todos os métodos
// ? de uma interface para o Golang entender que está implementando.
// ? Uma struct pode implementar mais que uma interface.
type retangulo struct {
	largura, altura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	radius float64
}

func (c circulo) area() float64 {
	return math.Pi * c.radius * c.radius
}

func showArea(shape geometria) {
	fmt.Println("Area: ", shape.area())
}

func Test() {

	retangulo := retangulo{
		altura:  1,
		largura: 2,
	}
	circulo := circulo{
		radius: 1,
	}

	showArea(retangulo)
	showArea(circulo)
}
