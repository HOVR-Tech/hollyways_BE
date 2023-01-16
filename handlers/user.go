package handlers

import (
	"encoding/json"
	"fmt"
	dto "hollyways/dto/result"
	userdto "hollyways/dto/user"
	"hollyways/models"
	"hollyways/repositories"
	"net/http"
	"strconv"

	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

func (h *handlerUser) FindUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user, err := h.UserRepository.FindUser()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: user}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(user)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userID := int(userInfo["id"].(float64))

	// CHECK USER BY ID
	getUser, err := h.UserRepository.GetUser(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "WROOONGG!"}
		json.NewEncoder(w).Encode(response)
		return
	}

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

	request := userdto.UpdateUserRequest{
		Image: resp.SecureURL,
	}
	validation := validator.New()
	errValidation := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: errValidation.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	user := models.User{
		Image: request.Image,
	}
	user, err = h.UserRepository.UpdateUser(user, getUser.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// user, _ = h.UserRepository.GetUser(user.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: user}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.UserRepository.DeleteUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := userdto.DeleteUserResponse{Code: http.StatusOK, Data: data, Message: "THIS USER HAS BEEN DELETED!!"}
	json.NewEncoder(w).Encode(response)
}

func convertResponse(u models.User) userdto.UserResponse {
	return userdto.UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		// Password: u.Password,
		Phone: u.Phone,
		Image: u.Image,
		Role:  u.Role,
	}
}
