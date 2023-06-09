package repository

import (
	"context"
	"database/sql"

	"github.com/Alfeenn/article/model"
)

type Repository interface {
	Create(ctx context.Context, tx *sql.Tx, category model.User) model.User
	Update(ctx context.Context, tx *sql.Tx, category model.User) model.User
	Delete(ctx context.Context, tx *sql.Tx, id string)
	FindAll(ctx context.Context, tx *sql.Tx) []model.User
	Find(ctx context.Context, tx *sql.Tx, id string) (model.User, error)
}
