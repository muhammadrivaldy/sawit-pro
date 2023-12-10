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
	InsertUsers(ctx context.Context, user models.User) (models.User, error)
	SelectUsersById(ctx context.Context, id int) (models.User, error)
}
