package main
import (
	"fmt"
	"math"
)

// shape interface
type shape interface {
	area() float64
	circumf() float64
}


//----------------------------------------------------------


type square struct {
	length float64
}
// square methods
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumf() float64 {
	return s.length * 4
}

//----------------------------------------------------------

type circle struct {
	radius float64
}
// circle methods
//mimic of polymorphism: circle implementing area and circumf function differently
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}
func (c circle) diam() float64 {
	return 2 * c.radius
}


//----------------------------------------------------------


//mimic of inheritance: circle and square is inheriting printShapeInfo() function
func printShapeInfo(s shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f \n", s, s.circumf())
}

//----------------------------------------------------------


func main() {
	shapes := []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}
	fmt.Println("diameter of circle", circle{radius: 12.3}.diam())

}