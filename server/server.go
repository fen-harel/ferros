package server

import (
	"bufio"
	"fmt"
	"io"

	"example/ferros/utils"
)

const prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		dir := scanner.Text()
        utilities.Reader(dir)

		fmt.Println(dir)
	}
}

