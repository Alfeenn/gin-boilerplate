package repository

import (
	"context"
	"database/sql"

	"github.com/Alfeenn/article/helper"
	"github.com/Alfeenn/article/model"
	"github.com/google/uuid"
)

type RepoImpl struct{}

func NewRepository() Repository {
	return &RepoImpl{}
}

func (r *RepoImpl) Create(ctx context.Context, tx *sql.Tx, category model.User) model.User {
	SQL := "INSERT INTO article(id,name,status,visibility,details) VALUES(?,?,?,?,?,?)"
	category.Id = uuid.NewString()
	_, err := tx.ExecContext(ctx, SQL,
		category.Id, category.Email, category.Password,
		category.Role, category.CreatedAt, category.UpdatedAt)
	helper.PanicIfErr(err)
	return category
}

func (r *RepoImpl) Update(ctx context.Context, tx *sql.Tx, category model.User) model.User {
	SQL := "UPDATE article SET name=? WHERE id=?"

	_, err := tx.ExecContext(ctx, SQL, category.Email, category.Id)
	helper.PanicIfErr(err)

	return category

}

func (r *RepoImpl) Delete(ctx context.Context, tx *sql.Tx, id string) {
	SQL := "DELETE FROM article WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfErr(err)
}

func (r *RepoImpl) FindAll(ctx context.Context, tx *sql.Tx) []model.User {
	sql := "SELECT *FROM user"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfErr(err)
	defer rows.Close()

	var sliceArticle []model.User

	for rows.Next() {
		article := model.User{}
		err := rows.Scan(&article.Id, &article.Email,
			&article.Password, &article.Role, &article.CreatedAt, &article.UpdatedAt)
		helper.PanicIfErr(err)
		sliceArticle = append(sliceArticle, article)
	}
	return sliceArticle
}

func (r *RepoImpl) Find(ctx context.Context, tx *sql.Tx, id string) (model.User, error) {
	SQL := "SELECT *FROM article WHERE id =?"

	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfErr(err)
	defer rows.Close()
	article := model.User{}
	if rows.Next() {
		rows.Scan(&article.Id, &article.Email,
			&article.Password, &article.Role, &article.CreatedAt, &article.UpdatedAt)

		return article, nil
	} else {
		return article, err
	}

}
