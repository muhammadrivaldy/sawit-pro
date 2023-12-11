// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"context"

	"github.com/SawitProRecruitment/UserService/models"
)

type RepositoryInterface interface {
	UsersRepositoryInterface
}

type UsersRepositoryInterface interface {
	InsertUsers(ctx context.Context, user models.User) (models.User, error)
	UpdateUsers(ctx context.Context, user models.User) (models.User, error)
	SelectUsersById(ctx context.Context, id int) (models.User, error)
	SelectUsersByPhoneNumber(ctx context.Context, phoneNumber string) (models.User, error)
}

type SessionRepositoryInterface interface {
	InsertSessions(ctx context.Context, session models.Session) (models.Session, error)
}
