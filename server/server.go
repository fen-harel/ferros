package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dat, err := os.ReadFile("../README.md")
	check(err)
	fmt.Print(string(dat))

	// f, err := os.Open("../README.md")
	// check(err)
}
