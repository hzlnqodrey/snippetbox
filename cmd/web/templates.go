package main

import "github.com/hzlnqodrey/snippetbox.git/pkg/models"
import "html/template"
import "path/filepath"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

// Chap 5.3 - Caching Templates
func newTemplateCache(dir string) (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache
	cache := map[string]*template.Template{}

	// filepath.Glob() to get a slice of all filepaths with the ext. 'page.tmpl'.
	// gives us a slice of all the 'page' templates for the application
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	// Loop through pages one-by-one
	for _, page := range pages {
		// Extract the file name (like 'home.page.tmpl') from the full file path
		// and assign it to the name variable.
		name := filepath.Base(page)

		// parse the page template
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Use the ParseGlob method to add any 'layout' templates to the
		// template set (in our case, it's just the 'footer' layout at the
		// moment).
		ts, err = template.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		// Use the ParseGlob method to add any 'partial' templates to the
		// template set (in our case, it's just the 'footer' partial at the
		// moment).
		ts, err = template.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		// ts = template set | add template set to cache by page name as a key
		cache[name] = ts
	}

	// return map
	return cache, nil
}
