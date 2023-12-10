package repository

import (
	"context"

	"github.com/SawitProRecruitment/UserService/models"
)

func (r *Repository) InsertUsers(ctx context.Context, user models.User) (models.User, error) {

	return models.User{}, nil

}

func (r *Repository) SelectUsersById(ctx context.Context, id int) (models.User, error) {

	return models.User{}, nil

}
