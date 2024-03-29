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

	// Chap 4.6 - Single-record SQL Queries
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// Chap 4.6 - Single-record SQL Queries
	row := m.DB.QueryRow(stmt, id)

	// Chap 4.6 - Single-record SQL Queries
	// Initialize a pointer to a new zeroed Snippet struct
	s := &models.Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	// everything works, return the snippet object
	return s, nil
}

// GET 10 RECENT SNIPPETS
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {

	// Chap 4.7 - Multiple records sql queries
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := []*models.Snippet{}

	// Chap 4.7 - Use rows.Next() to iterate through the rows in resultset
	for rows.Next() {
		// Create zero snippet struct
		s := &models.Snippet{}

		// use rows.scan
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

		if err != nil {
			return nil, err
		}

		// Append to the slice
		snippets = append(snippets, s)
	}

	// when rows.Next() is finished call the rows.Err
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// everything well went, return snippets slice
	return snippets, nil
}
