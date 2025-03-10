package handlers

import (
	"my_website/internal/repository"
)

type Handler struct {
	Repository *repository.Repository
}
