package handler

import (
	"github.com/SawitProRecruitment/UserService/configs"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
)

type Server struct {
	conf configs.Configuration
	repo repository.RepositoryInterface
}

func NewServer(conf configs.Configuration, repo repository.RepositoryInterface) generated.ServerInterface {
	return &Server{conf, repo}
}
