package main

import (
	"fmt"
)

// 基本数据类型转化成string

func main() {
	var num1 int = 1
	var num2 float64 = 1.2
	var b bool = true
	var mychar byte = 'h'
	var str string

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%v\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%v\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%v\n", str, str)

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str type %T str=%v\n", str, str)

}
