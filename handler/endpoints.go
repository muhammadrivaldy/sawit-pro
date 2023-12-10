package handler

import (
	"fmt"

	"github.com/SawitProRecruitment/UserService/payload"
	"github.com/SawitProRecruitment/UserService/utils"
	"github.com/labstack/echo/v4"
)

// (POST /users)
func (s *Server) PostUsers(ctx echo.Context) error {

	payload := payload.RequestPostUsers{
		PhoneNumber: ctx.FormValue("phone_number"),
		FullName:    ctx.FormValue("full_name"),
		Password:    ctx.FormValue("password"),
	}

	err := utils.ValidatorNew().Struct(payload)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}

// (GET /users/:id)
func (s *Server) GetUsersId(ctx echo.Context) error {

	return nil

}

// (PUT /users/:id)
func (s *Server) PutUsersId(ctx echo.Context) error {

	return nil

}

// (POST /users/login)
func (s *Server) PostUsersLogin(ctx echo.Context) error {

	return nil

}
