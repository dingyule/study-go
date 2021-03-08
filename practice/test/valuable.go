package main

import "fmt"

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showvalue(valuable2 valuable) {
	fmt.Println(valuable2.getValue())
}

func main() {
	var o valuable = stockPosition{
		ticker:     "today",
		sharePrice: 20,
		count:      4,
	}
	showvalue(o)
	o = car{
		make:  "dd",
		model: "hh",
		price: 10,
	}
	showvalue(o)
}
