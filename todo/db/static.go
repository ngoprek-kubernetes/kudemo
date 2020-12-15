package db

import (
	"time"

	"github.com/ngoprek-kubernetes/kudemo/todo/schema"
)

type Static struct{}

func (s *Static) GetAll() ([]schema.Todo, error) {
	todoList := []schema.Todo{
		{
			ID:       1,
			Title:    "Beli mie instan",
			Note:     "",
			Deadline: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:       2,
			Title:    "Isi pulsa",
			Note:     "",
			Deadline: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:       3,
			Title:    "Ambil uang di atm",
			Note:     "",
			Deadline: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	return todoList, nil
}

func (s *Static) Insert(todo *schema.Todo) (int, error) {
	return 0, nil
}

func (s *Static) Delete(id int) error {
	return nil
}

func (s *Static) Close() {}
