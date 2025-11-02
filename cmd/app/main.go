package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!!")
}

func main() {
	fmt.Println("String server...")

	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)

}
