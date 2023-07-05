package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
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

	// Chapter 4.3 - Database Connection Pool
	dsn := flag.String("dsn", "web:qodri123@/snippetbox?parseTime=true", "MySQL Database")
	flag.Parse()

	// Chapter 3.2 - Levelled Logging
	// info logging
	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// error logging
	errorlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Chapter 4.3 - Database Connection Pool
	db, err := openDB(*dsn)
	if err != nil {
		errorlog.Fatal(err)
	}

	// Chapter 4.3 - Database Connection Pool
	defer db.Close()

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
	ERR := srv.ListenAndServe()
	errorlog.Fatal(ERR)
}

// Chapter 4.3 Database Connection Pool
// the openDB() function wraps sql.Open() and returns a sql.DB connection pool for a given DSN
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
