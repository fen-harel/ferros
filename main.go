package main

import (
	"bufio"
	"fmt"
	"net/http"

	"example/ferros/server"
)

func main() {
	http.HandleFunc("/hello", server.Hello)
	http.HandleFunc("/headers", server.Headers)

	// http.ListenAndServe(":8090", nil)

	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
