//Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point
//с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

// инкапсулированные параметры - это те к которым нет доступа извне (в го на уровне пакетов)
type point struct {
	x float64
	y float64
}

// конструктор
func newPoint(x, y float64) *point {
	return &point{
		x: x,
		y: y,
	}
}

func (p1 *point) distance(p2 *point) float64 {
	return math.Sqrt(math.Pow((p2.x-p1.x), 2) + math.Pow((p2.y-p1.y), 2))
}

func main() {
	p1 := newPoint(4, 5)
	p2 := newPoint(6, 7)

	fmt.Println(p1.distance(p2))

}
