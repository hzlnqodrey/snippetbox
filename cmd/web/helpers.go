package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

// Server Error
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	// Frame Depth of trace-logging
	app.errorlog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Client Error - Bad Request of USER
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// Client Error - Bad Request of USER - 404 NOT FOUND
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

// Render - Retrieve the approriate template set. If no entries exist in the cache with the provided name, then call the serverError
func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("The templates %s does not exist", name))
		return
	}

	// Write to an in-memory buffer (instead of directly to a 'w') so 
	// to prevent incomplete HTML outputs in case of errors
	buf := new(bytes.Buffer) // A Buffer is a variable-sized buffer of bytes with Read and Write methods. The zero value for Buffer is an empty buffer ready to use

	// Execute the template set, passing in dynamic data,
	err := ts.Execute(buf, td)
	if err != nil {
		app.serverError(w, err)
	}
}
