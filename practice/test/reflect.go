package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x float64 = 3.4
	fmt.Println("x type :", reflect.TypeOf(x))

	v := reflect.ValueOf(x)
	fmt.Println("v type :", v.Type())

	fmt.Println("kind is float", v.Kind() == reflect.Float64)

	fmt.Println("kind :", v.Kind())

	p := reflect.ValueOf(&x)
	fmt.Println("type of p", p.Type())
	fmt.Println("set p", p.CanSet())

	v1 := p.Elem()
	fmt.Println("can set p", v1.CanSet())

	v1.SetFloat(7.1)
	fmt.Println(v1.Interface())
	fmt.Println(v1)

	t := T{200, "mh200"}
	s := reflect.ValueOf(&t).Elem()
	fmt.Println(s)
	typeofT := s.Type()
	fmt.Println("s type ", typeofT)

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeofT.Field(i).Name, f.Type(), f.Interface())
	}

}

type T struct {
	A int
	B string
}
