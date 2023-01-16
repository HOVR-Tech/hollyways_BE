package routes

import (
	"hollyways/handlers"
	"hollyways/pkg/middleware"
	"hollyways/pkg/mysql"
	"hollyways/repositories"

	"github.com/gorilla/mux"
)

func DonationRoutes(r *mux.Router) {
	donationRepository := repositories.RepositoryDonation(mysql.DB)
	h := handlers.HandlerDonation(donationRepository)

	r.HandleFunc("/donation", middleware.Auth(h.MakeDonation)).Methods("POST")
	r.HandleFunc("/donation", middleware.Auth(h.FindDonation)).Methods("GET")
	r.HandleFunc("/donation/{id}", middleware.Auth(h.GetDonation)).Methods("GET")
	r.HandleFunc("/donation-user/{id}", middleware.Auth(h.GetDonationByUserID)).Methods("GET")
	r.HandleFunc("/donation-fund/{id}", middleware.Auth(h.GetDonationByFund)).Methods("GET")
	r.HandleFunc("/snap/{id}", middleware.Auth(h.Snap)).Methods("GET")
	r.HandleFunc("/notification", h.Notification).Methods("POST")
	// r.HandleFunc("/transaction-upload/{id}", middleware.Auth(middleware.UploadFile(h.UpdateTransaction))).Methods("PATCH")
	r.HandleFunc("/donation/{id}", middleware.Auth(h.UpdateDonation)).Methods("PATCH")
	r.HandleFunc("/donation/{id}", middleware.Auth(h.DeleteDonation)).Methods("DELETE")
}
