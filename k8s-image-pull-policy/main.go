package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "<h1>Hello World</h1>")
}


func main() {
	http.HandleFunc("/", index)
	_ = http.ListenAndServe(":8888", nil)
}
