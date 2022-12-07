package reader

import (
	"fmt"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Reader(dir string) {
	dat, err := os.ReadFile(dir)
	Check(err)
	fmt.Print(string(dat))

}
