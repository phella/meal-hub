package userRouter

import (
	"Bete/pkg/httputils"
	userService "Bete/services/user"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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
	var ensureUserParams userService.EnsureUserParams
	err := json.NewDecoder(req.Body).Decode(&ensureUserParams)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	ensureUserParams.Token = uuid.New().String()
	token := r.userService.EnsureUser(ensureUserParams)

	err = httputils.JSON(w, token)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
