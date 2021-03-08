package main

import (
	"fmt"
	"math"
)

type Square1 struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper1 interface {
	Area() float32
}

func (s Square1) Area() float32 {
	return s.side * s.side
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {
	//var areaIntf Shaper1
	//sq1 := new(Square1)
	//sq1.side = 5
	//
	//areaIntf = sq1

	//if t, ok := areaIntf.(*Square1); ok {
	//	fmt.Printf("type: %T\n", t)
	//}
	//
	//if u, ok := areaIntf.(*Circle); ok {
	//	fmt.Printf("type: %T", u)
	//}else {
	//	fmt.Printf("not circle ")
	//}

	//switch t := areaIntf.(type) {
	//case *Square1:
	//	fmt.Printf("Type Square %T with value %v\n", t, t)
	//case Circle:
	//	fmt.Printf("Type Circle %T with value %v\n", t, t)
	//}

	var v Stringer
	if sv, ok := v.(Stringer); ok {
		fmt.Printf(" v implements string(): %s\n", sv.String())
	}

}

type Stringer interface {
	String() string
}
