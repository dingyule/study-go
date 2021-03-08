package main

import (
	"fmt"
)

func main() {
	//fmt.Println("hello,world !")
	//const freezingF, boilingF = 32.0, 212.0
	//fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	//fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))

	//array := [5]int{1,2,3,4,5}
	//modify(array)
	//fmt.Println("in main(), array values", array)

	//var myarray [10]int  = [10]int{1,2,3,4,5,6,7,8,9,10}
	//var myslice []int  = myarray[:5]
	//fmt.Println("myarray: " )
	//for _, v := range myarray{
	//	fmt.Print(v, " ")
	//}
	//fmt.Println("\nmyslice : ")
	//
	//for _, v:= range myslice{
	//	fmt.Print(v, " ")
	//}

	//var person map[string] PersonInfo
	//person = make(map[string] PersonInfo)
	//person["12345"] = PersonInfo{"12345" , "xiaoming"}
	//person["1"] = PersonInfo{"1", "xiaohu"}
	//
	//per , ok := person["12345"]
	//if ok {
	//	fmt.Println("Found person", per.Name, "with ID 1234.")
	//}else {
	//	fmt.Println("Did not find person with ID 1234.")
	//}

	//sum := 0
	//
	//for i := 0; i < 10; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)

	//myfunc(2, 3, 4)
	//var v1 int = 12
	//var v2 int = 14
	//var v3 string = "hello"
	//var v4 bool = false
	//myfunc2(v1, v2, v3, v4)

	//var j int  = 5
	//a := func()(func()) {
	//	var i int  = 10
	//	return func() {
	//		fmt.Printf("i , j : ,%d ,%d\n", i, j)
	//	}
	//}()
	//
	//a()
	//j *= 2
	//fmt.Println("------")
	//a()
	//fmt.Printf("j = %d", j)

	//var a Integer = 1
	//if Interger_less(a ,2) {
	//	fmt.Println(a, "less 2")
	//}

	//var a Integer  =  1
	////a.Add(2)
	//a.Add1(2)
	//fmt.Println("a = " , a)

	//var a  = [3]int{1, 2, 3}
	//var b  = a
	//b[1] ++
	//fmt.Println(a , b)
	//var a  = [3]int{1, 2, 3}
	//var b  = &a
	//b[1] ++
	//fmt.Println(a , *b)

	//var v1 interface{} = 1
	//var v2 interface{} = "abc"
	//var v3 interface{} = &v2
	//fmt.Println(v2)
	//fmt.Println(v3)

	//go Add(1, 1)

	//c := make(chan int, 1024)
	//
	//for i:= range c {
	//	fmt.Println("received" , i)
	//}
	//gobook := Book{"go" ,"xiaoming", "ccc" , true, 9.99}
	//b, _ := json.Marshal(gobook)
	//fmt.Println(b)

	//var movies  = []Movie{
	//	{Title: "Casablanca", Year: 1942, Color: false,
	//		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	//	{Title: "Cool Hand Luke", Year: 1967, Color: true,
	//		Actors: []string{"Paul Newman"}},
	//	{Title: "Bullitt", Year: 1968, Color: true,
	//		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	//}
	//
	//data, err := json.Marshal(movies)
	//
	//if err != nil {
	//	log.Fatalf("json failed :", err)
	//}
	//fmt.Printf(" ", data)
	//fmt.Println(movies)

	//go func() {
	//	bank.Deposit(200)
	//	fmt.Println("=" , bank.Balance())
	//}()

	//go bank.Deposit(100)

	//fmt.Println(",.....")

	const (
		a = iota
		b
		c
	)
	fmt.Println(a, b, c)

}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func modify(array [5]int) {
	array[0] = 10
	fmt.Println("in modify(), array values ", array)
}

type PersonInfo struct {
	ID   string
	Name string
}

func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func myfunc2(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is a int value")
		case string:
			fmt.Println(arg, "is a string value")
		case bool:
			fmt.Println(arg, "is a bool value")
		}
	}
}

type Integer int

func Interger_less(a Integer, b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}

func (a Integer) Add1(b Integer) {
	a += b
}

type Rect struct {
	x, y          float64
	width, heigth float64
}

func (r *Rect) Area() float64 {
	return r.heigth * r.width
}

func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

type Base struct {
	Name string
}

func (base *Base) Foo() {

}

func (base *Base) Bar() {

}

type Foo struct {
	Base
}

func (foo *Foo) Bar() {
	foo.Base.Bar()
}

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

type Book struct {
	Tittle      string
	Authors     string
	Publisher   string
	IsPublished bool
	Price       float64
}

type Movie struct {
	Title  string
	Year   int
	Color  bool
	Actors []string
}
