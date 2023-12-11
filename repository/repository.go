// This file contains the repository implementation layer.
package repository

import (
	"github.com/SawitProRecruitment/UserService/configs"
	_ "github.com/lib/pq"
	goutil "github.com/muhammadrivaldy/go-util"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type Repository struct {
	GormDb *gorm.DB
}

func NewRepository(config configs.Configuration) RepositoryInterface {

	pgDb, err := goutil.NewPostgreSQL(config.Database.User, config.Database.Password, config.Database.Address, config.Database.Schema, null.String{}, config.Database.Port)
	if err != nil {
		panic(err)
	}

	gormDb, err := goutil.NewGorm(pgDb, "postgres", goutil.LoggerSilent)
	if err != nil {
		panic(err)
	}

	return &Repository{GormDb: gormDb}

}
