package userService

import (
	"Bete/models"
)

type Service interface {
	EnsureUser(EnsureUserParams) models.User
}

type User struct {
	Token string
	Name  string
}

type EnsureUserParams struct {
	Token string
	Name  string
}
