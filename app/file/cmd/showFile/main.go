package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("../pool/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":9999", nil)
}
