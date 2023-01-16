package routes

import (
	"hollyways/handlers"
	"hollyways/pkg/middleware"
	"hollyways/pkg/mysql"
	"hollyways/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/user", h.FindUser).Methods("GET")
	r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
	// r.HandleFunc("/user/{id}", h.UpdateUser).Methods("PATCH")
	r.HandleFunc("/user/{id}", middleware.Auth(middleware.UploadFile(h.UpdateUser))).Methods("PATCH")
	r.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE")
}
