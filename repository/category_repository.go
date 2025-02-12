package repository

import (
	"context"
	"database/sql"
	"rest-api-pzn-go/model/domain"
)

type CategoryRepository interface {
	// pake fungsi transaksional, sebenernya bisa pake atau engga, di contoh ini pake
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
