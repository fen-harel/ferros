package main

import (
	"example/ferros/server"
	"os"
)

func main() {
	server.Start(os.Stdin, os.Stdout)
}
