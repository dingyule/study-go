package main

type inners struct {
	int1 int
	int2 int
}

type outers struct {
	b int
	c float32
	int
	inners
}

func main() {
	//outer := new(outers)
	//outer.b = 1
	//outer.c = 2.0
	//outer.int = 2
	//outer.int1 = 3
	//outer.int2 = 4
	////fmt.Println(outer.int2)
	//fmt.Println(outer)
	//
	//outer1 := outers{
	//	b:      2,
	//	c:      3,
	//	int:    4,
	//	inners: inners{4 ,6},
	//}
	//fmt.Println("outer1: ", outer1)

	//b := B{
	//	A:  A{1,2},
	//	bx: 2.2,
	//	by: 2.3,
	//}
	//fmt.Println(b.ax, b.ay, b.bx, b.by)
	//fmt.Println(b.A)

	//e := E{
	//	C: C{2},
	//	D: D{3,4},
	//}
	//fmt.Println(e.C.a, e.D.a)

}

type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by float32
}

type C struct {
	a int
}

type D struct {
	a, b int
}

type E struct {
	C
	D
}
