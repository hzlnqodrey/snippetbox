package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/hzlnqodrey/snippetbox.git/pkg/models"
)

// Chapter 3.3 - dependency Injection
// Change the signature of the home handler so it is defined as a method agains application.

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// Chapter 3.4 - Centralized Error Handling - Use NotFound() helper
		app.notFound(w)
		return
	}

	// Comment Out
	// s, err := app.snippets.Latest()
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	// for _, snippet := range s {
	// 	fmt.Fprintf(w, "%v\n", snippet)
	// }

	// Initialize a slice containing the paths to the two files. Note that the
	// home.page.tmpl file must be the *first* file in the slice.
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// turn files into variadic
	ts, err := template.ParseFiles(files...)

	if err != nil {
		// Chapter 3.4 - Centralized Error Handling - Use serverError() helper
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)

	if err != nil {
		// Chapter 3.4 - Centralized Error Handling - Use serverError() helper
		app.serverError(w, err)
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	// Chap 4.6 - Using the Models
	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	// Chap 4.6 - Write the snippet data as a plain-text HTTP response body.
	// fmt.Fprintf(w, "%v", s)

	// Chap 5.1 - Rendering Multiple Data
	// Create an instance of a templateData struct holding the snippet data.
	data := &templateData{Snippet: s}


	// Chap 5.1 - Displaying Dynamic Data
	files := []string{
		"./ui/html/show.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// Parse the template files
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Pass in the templateData struct when executing the template.
	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "Post")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// Chap 4.5 - dummy data
	title := "O snail"
	content := "O snail l\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi"
	expires := "7"

	// Chap 4.5 - Pass the data to the SnippetModel.Insert() method, receiving the ID
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)

		return
	}

	// Chap 4.5 - redirect the user to the relevant page for the snippet
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)

	w.Write([]byte("Create a new snippet"))
}
