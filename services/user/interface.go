package userService

type Service interface {
	CreateUser(CreateUserParams) User
}

type User struct {
	Token string
}

type CreateUserParams struct {
	Token string
}
