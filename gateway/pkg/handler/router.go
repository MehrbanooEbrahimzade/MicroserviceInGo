package handler

import (
	"github.com/MehrbanooEbrahimzade/MicroserviceInGo/gateway/pkg/models"
	userSvc "github.com/MehrbanooEbrahimzade/MicroserviceInGo/gateway/pkg/service/user"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type jsonResponse struct {
	Data interface{} `json:"data,omitempty"`
	Err  string      `json:"error,omitempty"`
}
type Handler struct {
	userSvc *userSvc.UserSvc
}

func NewHandler(u *userSvc.UserSvc) *Handler {
	return &Handler{
		userSvc: u,
	}
}

func NewHandlerRoute(c models.Cors, h *Handler) http.Handler {
	mux := mux.NewRouter()
	allowedOrigins := handlers.AllowedOrigins(c.AllowedOrigins)
	allowedMethods := handlers.AllowedMethods(c.AllowedMethods)
	allowedHeaders := handlers.AllowedHeaders(c.AllowedHeaders)
	exposedHeaders := handlers.ExposedHeaders(c.ExposedHeaders)
	maxAge := handlers.MaxAge(c.MaxAge)
	allowCredentials := handlers.AllowCredentials()

	cors := handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders, exposedHeaders, allowCredentials, maxAge)

	getRouter := mux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/users", h.getAllUsers)

	putRouter := mux.Methods(http.MethodPost).Subrouter()
	putRouter.HandleFunc("/users", h.CreateUser)

	return cors(mux)
}
