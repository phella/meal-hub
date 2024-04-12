package userRouter

import (
	"Bete/pkg/httputils"
	"Bete/services/user"
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type params struct {
	fx.In

	UserService userService.Service
}

type Router interface {
	RegisterRoutes() chi.Router
}

type userRouter struct {
	userService userService.Service
}

func New(p params) Router {
	return &userRouter{
		userService: p.UserService,
	}
}

func (r userRouter) RegisterRoutes() chi.Router {
	router := chi.NewRouter()
	router.Post("/", r.createUser)
	return router
}

func (r userRouter) createUser(w http.ResponseWriter, req *http.Request) {
	var createUserParams userService.CreateUserParams
	err := json.NewDecoder(req.Body).Decode(&createUserParams)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	token := r.userService.CreateUser(createUserParams)

	httputils.JSON(w, token)
}
