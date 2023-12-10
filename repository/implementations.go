package repository

import (
	"context"

	"github.com/SawitProRecruitment/UserService/models"
)

func (r *Repository) InsertUsers(ctx context.Context, user models.User) (models.User, error) {
	err := r.gormDb.Create(&user).Error
	return user, err
}

func (r *Repository) SelectUsersById(ctx context.Context, id int) (models.User, error) {
	return models.User{}, nil
}
