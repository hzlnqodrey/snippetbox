package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// Chapter 3.1 - CLI Flag
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	flag.Parse()

	// Chapter 3.2 - Levelled Logging
	// info logging
	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// error logging
	errorlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorlog,
		Handler:  mux,
	}

	infolog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorlog.Fatal(err)
}
