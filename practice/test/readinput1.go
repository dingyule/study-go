package main

import "fmt"

var (
	firstname, lastname, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func main() {
	//fmt.Println("please enter your full name")
	//fmt.Scanln(&firstname, &lastname)
	//fmt.Printf("%s, %s", firstname, lastname)
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
}
