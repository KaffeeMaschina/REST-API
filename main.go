package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello from a Handlefunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello from a Handlefunc #2!\n")
	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)

	log.Fatal(http.ListenAndServe("185.5.249.91:22", nil))
}
