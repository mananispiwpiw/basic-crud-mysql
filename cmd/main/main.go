package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	address := "localhost:8080"
	fmt.Printf("Server started at %s\n", address)
	err := http.ListenAndServe(address, mux)
	if err != nil {
		fmt.Println(err.Error())
	}
}
