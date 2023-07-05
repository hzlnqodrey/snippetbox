package mysql

import (
	"database/sql"

	"github.com/hzlnqodrey/snippetbox.git/pkg/models"
)

// Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// POST SNIPPETS
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// GET SNIPPETS
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// GET 10 RECENT SNIPPETS
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
