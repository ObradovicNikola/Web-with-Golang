package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Obrad", "Served from Obrad server")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Super cool code here</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
