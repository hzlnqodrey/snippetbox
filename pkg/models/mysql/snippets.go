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

	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	
	if err != nil {
		return 0, nil
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, nil
	}

	return int(id), nil
}

// GET SNIPPETS
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// GET 10 RECENT SNIPPETS
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
