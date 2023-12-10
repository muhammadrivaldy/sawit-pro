package main

import (
	"fmt"

	"github.com/SawitProRecruitment/UserService/configs"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/labstack/echo/v4"
	goutil "github.com/muhammadrivaldy/go-util"
)

func main() {

	config := loadConfig()
	e := echo.New()

	repo := repository.NewRepository(config)
	serv := handler.NewServer(repo)
	generated.RegisterHandlers(e, serv)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}

func loadConfig() configs.Configuration {

	osFile, err := goutil.OpenFile("./../configs", "local.conf")
	if err != nil {
		panic(err)
	}
	defer osFile.Close()

	var config configs.Configuration
	if err := goutil.Configuration(osFile, &config); err != nil {
		panic(err)
	}

	return config

}
