package handler

import "github.com/web-dev137/test-task/repository"

type Handler struct {
	Repo *repository.Repo
}

func NewHandler(repo *repository.Repo) *Handler {
	return &Handler{Repo: repo}
}
