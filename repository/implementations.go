package repository

import (
	"context"

	"github.com/SawitProRecruitment/UserService/models"
)

func (r *Repository) InsertUsers(ctx context.Context, user models.User) (models.User, error) {
	err := r.gormDb.Create(&user).Error
	return user, err
}

func (r *Repository) UpdateUsers(ctx context.Context, user models.User) (models.User, error) {
	err := r.gormDb.Updates(&user).Error
	// err := r.gormDb.Model(&models.User{}).Where("id = ?", user.Id).Updates(&user).Error
	return user, err
}

func (r *Repository) SelectUsersById(ctx context.Context, id int) (user models.User, err error) {
	err = r.gormDb.Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *Repository) SelectUsersByPhoneNumber(ctx context.Context, phoneNumber string) (user models.User, err error) {
	err = r.gormDb.Where("phone_number = ?", phoneNumber).First(&user).Error
	return user, err
}
