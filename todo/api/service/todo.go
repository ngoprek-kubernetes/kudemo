package service

import (
	"context"

	"github.com/ngoprek-kubernetes/kudemo/todo/api/db"
	"github.com/ngoprek-kubernetes/kudemo/todo/api/schema"
)

func GetAll(ctx context.Context) ([]schema.Todo, error) {
	return db.GetAll(ctx)
}

func Insert(ctx context.Context, todo *schema.Todo) (int, error) {
	return db.Insert(ctx, todo)
}

func Update(ctx context.Context, todo *schema.Todo) error {
	return db.Update(ctx, todo)
}

func Delete(ctx context.Context, id int) error {
	return db.Delete(ctx, id)
}

func Close(ctx context.Context) {
	db.Close(ctx)
}
