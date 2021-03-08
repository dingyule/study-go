package main

import (
	"fmt"
	"regexp"
)

func main() {

	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+"

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("match found !")
	}

	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)

}
