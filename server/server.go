package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"example/ferros/utils"
)

const prompt = ">>"

func start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		dir := scanner.Text()
		reader.Reader(dir)

		// fmt.Println(line)
	}
}

func main() {
	start(os.Stdin, os.Stdout)
}
