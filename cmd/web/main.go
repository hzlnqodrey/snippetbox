package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Chapter 3.1 - CLI Flag
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	flag.Parse()

	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
