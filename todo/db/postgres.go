package db

import (
	"database/sql"

	"github.com/ngoprek-kubernetes/kudemo/todo/schema"
)

type Postgres struct {
	DB *sql.DB
}

func (p *Postgres) GetAll() ([]schema.Todo, error) {
	query := `
		SELECT *
		FROM todo
		ORDER BY id;
	`

	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var todoList []schema.Todo
	for rows.Next() {
		var t schema.Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Note, &t.Deadline); err != nil {
			return nil, err
		}
		todoList = append(todoList, t)
	}
	return todoList, nil
}

func (p *Postgres) Insert(todo *schema.Todo) (int, error) {
	query := `
		INSERT INTO todo (id, title, note, deadline)
		VALUES(nextval(todo_id), $1, $2, $3)
		RETURNING id;
	`

	rows, err := p.DB.Query(query, todo.ID, todo.Note, todo.Deadline)
	if err != nil {
		return -1, err
	}

	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return -1, err
		}
	}
	return id, nil
}

func (p *Postgres) Delete(id int) error {
	query := `
		DELETE FROM todo
		WHERE id = $1;
	`

	if _, err := p.DB.Exec(query, id); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) Close() {
	p.DB.Close()
}
