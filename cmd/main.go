package main

import (
	"fmt"

	"github.com/SawitProRecruitment/UserService/configs"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/utils"
	"github.com/labstack/echo/v4"
	goutil "github.com/muhammadrivaldy/go-util"
)

func main() {

	conf := loadConfig()
	e := echo.New()
	e.Use(utils.ValidateToken)

	repo := repository.NewRepository(conf)
	serv := handler.NewServer(repo)
	generated.RegisterHandlers(e, serv)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))

}

func loadConfig() configs.Configuration {

	osFile, err := goutil.OpenFile("./../configs", "local.conf")
	if err != nil {
		panic(err)
	}
	defer osFile.Close()

	var conf configs.Configuration
	if err := goutil.Configuration(osFile, &conf); err != nil {
		panic(err)
	}

	return conf

}
