package utilities

import (
	"fmt"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Reads byte from scanner
// text input and prints
// back in text form
func Reader(dir string) {
	dat, err := os.ReadFile(dir)
	Check(err)
	fmt.Print(string(dat))

}
