package handler

import (
	"net/http"
	"time"

	"github.com/SawitProRecruitment/UserService/models"
	"github.com/SawitProRecruitment/UserService/payloads"
	"github.com/SawitProRecruitment/UserService/utils"
	"github.com/labstack/echo/v4"
)

// (POST /users)
func (s *Server) PostUsers(ctx echo.Context) error {

	param := payloads.RequestPostUsers{
		PhoneNumber: ctx.FormValue("phone_number"),
		FullName:    ctx.FormValue("full_name"),
		Password:    ctx.FormValue("password"),
	}

	err := utils.ValidatorNew().Struct(param)
	if err != nil {
		ctx.Logger().Error(err)
		return err
	}

	passwordHash, err := utils.GeneratePassword(param.Password)
	if err != nil {
		ctx.Logger().Error(err)
		return err
	}

	user, err := s.repo.InsertUsers(ctx.Request().Context(), models.User{
		FullName:     param.FullName,
		PhoneNumber:  param.PhoneNumber,
		PasswordHash: passwordHash,
		CreatedBy:    0,
		CreatedAt:    time.Now(),
	})
	if err != nil {
		ctx.Logger().Error(err)
		return err
	}

	payloads.ResponseOK(ctx, http.StatusCreated, payloads.ResponsePostUsers{
		Id:          user.Id,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	})

	return nil

}

// (GET /users/:id)
func (s *Server) GetUsersId(ctx echo.Context, id int) error {

	return nil

}

// (PUT /users/:id)
func (s *Server) PutUsersId(ctx echo.Context, id int) error {

	return nil

}

// (POST /users/login)
func (s *Server) PostUsersLogin(ctx echo.Context) error {

	return nil

}
