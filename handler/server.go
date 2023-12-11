package handler

import (
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
)

type Server struct {
	repo repository.RepositoryInterface
}

func NewServer(repo repository.RepositoryInterface) generated.ServerInterface {
	return &Server{repo}
}
