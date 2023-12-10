// This file contains the repository implementation layer.
package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type Repository struct {
	db     *sql.DB
	gormDb *gorm.DB
}

type NewRepositoryOptions struct {
	Dsn string
}

func NewRepository(opts NewRepositoryOptions) RepositoryInterface {

	db, err := sql.Open("postgres", opts.Dsn)
	if err != nil {
		panic(err)
	}

	return &Repository{db: db, gormDb: nil}

}
