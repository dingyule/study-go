package main

import "fmt"

type Shaper interface {
	Area() float32

	//hello() float64
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r *Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	//sq1 := new(Square)
	//sq1.side = 5
	//var areaIntf Shaper
	//areaIntf = sq1
	//fmt.Printf("Square area: %f\n", areaIntf.Area())
	//fmt.Printf("sql1: %f", sq1.Area())

	r := &Rectangle{
		length: 5,
		width:  3,
	}

	q := &Square{3}

	//shapes := []Shaper{r, q}
	shapes := []Shaper{Shaper(r), Shaper(q)}
	for n, _ := range shapes {
		fmt.Println(shapes[n])
		fmt.Println("area: ", shapes[n].Area())
	}

}
