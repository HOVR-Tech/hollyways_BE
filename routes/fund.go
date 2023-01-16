package routes

import (
	"hollyways/handlers"
	"hollyways/pkg/middleware"
	"hollyways/pkg/mysql"
	"hollyways/repositories"

	"github.com/gorilla/mux"
)

func FundRoutes(r *mux.Router) {
	fundRepository := repositories.RepositoryFund(mysql.DB)
	h := handlers.HandlerFund(fundRepository)

	r.HandleFunc("/fund", h.FindFund).Methods("GET")
	r.HandleFunc("/fund/{id}", h.GetFund).Methods("GET")
	r.HandleFunc("/fund", middleware.Auth(middleware.UploadFile(h.AddFund))).Methods("POST")
	r.HandleFunc("/fund/{id}", middleware.Auth(middleware.UploadFile(h.EditFund))).Methods("PATCH")
	r.HandleFunc("/fund/{id}", middleware.Auth(h.DeleteFund)).Methods("DELETE")
}
