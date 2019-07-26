package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<!--not serving from our server-->
	<img src="https://saznajlako.com/wp-content/uploads/2014/04/Nemacki-ovcar.jpg">
	`)
}
