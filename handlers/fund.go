package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	funddto "hollyways/dto/fund"
	dto "hollyways/dto/result"
	"hollyways/models"
	"hollyways/repositories"
	"net/http"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerFund struct {
	FundRepository repositories.FundRepository
}

func HandlerFund(FundRepository repositories.FundRepository) *handlerFund {
	return &handlerFund{FundRepository}
}

func (h *handlerFund) AddFund(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContext := r.Context().Value("dataFile")
	filePath := dataContext.(string)

	ctx := context.Background()
	CLOUD_NAME := "cloudme19"
	API_KEY := "281282276658861"
	API_SECRET := "uxlAfli9ExpwY2o6j8qTS0gRJ9g"

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
	resp, err := cld.Upload.Upload(ctx, filePath, uploader.UploadParams{Folder: "hollyways"})
	if err != nil {
		fmt.Println("Upload Gagal!", err.Error())
	}

	days, _ := strconv.Atoi(r.FormValue("days"))
	donation_limit, _ := strconv.Atoi(r.FormValue("donation_limit"))
	request := funddto.FundRequest{
		Title:         r.FormValue("title"),
		Days:          days,
		DonationLimit: donation_limit,
		Description:   r.FormValue("description"),
		Image:         resp.SecureURL,
	}
	validation := validator.New()
	errValidation := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: errValidation.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	fund := models.Fund{

		Title:         request.Title,
		Days:          request.Days,
		DonationLimit: request.DonationLimit,
		Description:   request.Description,
		Image:         request.Image,
	}
	fund, err = h.FundRepository.AddFund(fund)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	fund, _ = h.FundRepository.GetFund(fund.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: fund}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFund) FindFund(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fund, err := h.FundRepository.FindFund()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: fund}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFund) GetFund(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "multipart/form-data")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	fund, err := h.FundRepository.GetFund(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: fund}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerFund) EditFund(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContext := r.Context().Value("dataFile")
	filePath := dataContext.(string)

	ctx := context.Background()
	CLOUD_NAME := "cloudme19"
	API_KEY := "281282276658861"
	API_SECRET := "uxlAfli9ExpwY2o6j8qTS0gRJ9g"

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	dataFund, err := h.FundRepository.GetFund(int(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
	resp, err := cld.Upload.Upload(ctx, filePath, uploader.UploadParams{Folder: "dewe_tour"})
	if err != nil {
		fmt.Println("Upload Gagal!", err.Error())
	}

	days, _ := strconv.Atoi(r.FormValue("days"))
	donation_limit, _ := strconv.Atoi(r.FormValue("donation_limit"))
	request := funddto.FundRequest{
		Title:         r.FormValue("title"),
		Days:          days,
		DonationLimit: donation_limit,
		Description:   r.FormValue("description"),
		Image:         resp.SecureURL,
	}

	if request.Title != "" {
		dataFund.Title = request.Title
	}
	if request.Days != 0 {
		dataFund.Days = request.Days
	}
	if request.DonationLimit != 0 {
		dataFund.DonationLimit = request.DonationLimit
	}
	if request.Description != "" {
		dataFund.Description = request.Description
	}
	if request.Image != "" {
		dataFund.Image = request.Image
	}

	data, err := h.FundRepository.EditFund(dataFund, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFund) DeleteFund(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	fund, err := h.FundRepository.GetFund(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.FundRepository.DeleteFund(fund, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

// func convertDataFund(u []models.Fund) funddto.FundResponse {
// 	return funddto.FundResponse{
// 		ID:            ,
// 		Title:         u.Title,
// 		Days:          u.Days,
// 		DonationLimit: u.DonationLimit,
// 		Description:   u.Description,
// 		UserID:        u.UserID,
// 		User:          u.User,
// 		Image:         u.Image,
// 		Donatur:       u.Donatur,
// 	}
// }
