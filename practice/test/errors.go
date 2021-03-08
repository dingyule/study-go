package main

import (
	"errors"
	"fmt"
)

var errornotfound = errors.New("not found error")

func main() {
	fmt.Printf("err: %v", errornotfound)
}
