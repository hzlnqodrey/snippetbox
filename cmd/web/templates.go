package main

import "github.com/hzlnqodrey/snippetbox.git/pkg/models"

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}
