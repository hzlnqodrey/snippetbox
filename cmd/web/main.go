package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Chapter 3.3 - dependency Injection
// For now we'll only include fields for the two custom logger
// we'll add more to it as the build progresses.
type application struct {
	errorlog *log.Logger
	infolog  *log.Logger
}

func main() {
	// Chapter 3.1 - CLI Flag
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	flag.Parse()

	// Chapter 3.2 - Levelled Logging
	// info logging
	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// error logging
	errorlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Chapter 3.3 - dependency Injection
	// Initialize a new instance of application containing the dependencies
	app := &application{
		errorlog: errorlog,
		infolog:  infolog,
	}

	// Chapter 3.5 - Isolation the Application Routes
	// This is quite a bit neater. The routes for our application are now isolated and encapsulated in the app.routes() method,
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorlog,
		Handler:  app.routes(),
	}

	infolog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorlog.Fatal(err)
}
