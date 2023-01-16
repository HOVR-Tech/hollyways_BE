package handlers

import (
	"encoding/json"
	"fmt"
	donationdto "hollyways/dto/donation"
	dto "hollyways/dto/result"
	"hollyways/models"
	"hollyways/repositories"
	"math/rand"

	"net/http"
	// "os"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	// "gopkg.in/gomail.v2"
)

type handlerDonation struct {
	DonationRepository repositories.DonationRepository
}

func HandlerDonation(donationRepository repositories.DonationRepository) *handlerDonation {
	return &handlerDonation{donationRepository}
}

func (h *handlerDonation) MakeDonation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userID := int(userInfo["id"].(float64))

	donation_amt, _ := strconv.Atoi(r.FormValue("donation_amt"))
	fund_id, _ := strconv.Atoi(r.FormValue("fund_id"))
	request := donationdto.DonationRequest{
		DonationAmt: donation_amt,
		Status:      r.FormValue("status"),
		FundID:      fund_id,
		UserID:      userID,
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// validation := validator.New()
	// err = validation.Struct(request)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	donation := models.Donation{
		DonationAmt: request.DonationAmt,
		Status:      request.Status,
		FundID:      request.FundID,
		UserID:      request.UserID,
	}

	var DonationIDIsMatch = false
	var DonationID int
	for !DonationIDIsMatch {
		DonationID = donation.ID + rand.Intn(10000) - rand.Intn(100)
		transactionData, _ := h.DonationRepository.GetDonation(DonationID)
		if transactionData.ID == 0 {
			DonationIDIsMatch = true
		}
	}

	newDonation, err := h.DonationRepository.MakeDonation(donation, DonationID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// var DonateIdIsMatch = false
	// var DonationId int
	// for !DonateIdIsMatch {
	// 	DonationId = int(time.Now().Unix())
	// 	donationData, _ := h.DonationRepository.GetDonation(DonationId)
	// 	if donationData.ID == 0 {
	// 		DonateIdIsMatch = true
	// 	}
	// }

	donate, _ := h.DonationRepository.GetDonation(newDonation.ID)
	fmt.Println("data donation", newDonation.UserID)

	var s = snap.Client{}
	s.New("SB-Mid-server-yabfvfwIWlF_17N59lI1WuJ0", midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(donate.ID),
			GrossAmt: int64(donate.DonationAmt),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: donate.User.Name,
			Email: donate.User.Email,
		},
	}

	snapResp, _ := s.CreateTransaction(req)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: snapResp}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerDonation) Snap(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "multipart/form-data")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	dataDonation, err := h.DonationRepository.GetDonation(int(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	request := donationdto.DonationRequest{
		Status: r.FormValue("status"),
	}

	if request.Status != "" {
		dataDonation.Status = request.Status
	}

	data, err := h.DonationRepository.UpdateDonation(dataDonation, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	donate, _ := h.DonationRepository.GetDonation(data.ID)
	fmt.Println()

	var s = snap.Client{}
	s.New("SB-Mid-server-yabfvfwIWlF_17N59lI1WuJ0", midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(donate.ID),
			GrossAmt: int64(donate.DonationAmt),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: donate.User.Name,
			Email: donate.User.Email,
		},
	}

	snapResp, _ := s.CreateTransaction(req)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: snapResp}
	json.NewEncoder(w).Encode(response)
}

// func (h *handlerDonation) UpdateDonation(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "multipart/form-data")
// 	dataContext := r.Context().Value("dataFile")
// 	filepath := dataContext.(string)

// 	// ctx := context.Background()
// 	CLOUD_NAME := "cloudme19"
// 	API_KEY := "281282276658861"
// 	API_SECRET := "uxlAfli9ExpwY2o6j8qTS0gRJ9g"

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])
// 	dataTransaction, err := h.DonationRepository.GetDonation(int(id))
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	}

// 	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
// 	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "dewe_tour"})
// 	if err != nil {
// 		fmt.Println("Upload Gagal!", err.Error())
// 	}

// 	request := transactiondto.TransactionRequest{
// 		Image: resp.SecureURL,
// 	}

// 	if request.Image != "" {
// 		dataTransaction.Image = request.Image
// 	}

// 	data, err := h.TransactionRepository.UpdateTransaction(dataTransaction, id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	}

// 	trans, _ := h.TransactionRepository.GetTransaction(data.ID)

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: trans}
// 	json.NewEncoder(w).Encode(response)
// }

func (h *handlerDonation) Notification(w http.ResponseWriter, r *http.Request) {
	var notificationPayload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&notificationPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	donationStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)

	donation, _ := h.DonationRepository.GetOneDonation(orderId)
	fmt.Println(donationStatus, fraudStatus, orderId, donation)
	fmt.Println(notificationPayload)
	if donationStatus == "capture" {
		if fraudStatus == "challenge" {
			SendEmail("pending", donation)
			h.DonationRepository.Notification("pending", donation.ID)
		} else if fraudStatus == "accept" {
			SendEmail("Approved", donation)
			h.DonationRepository.Notification("Approved", donation.ID)
		} else if donationStatus == "settlement" {
			SendEmail("Approved", donation)
			h.DonationRepository.Notification("Approved", donation.ID)
		} else if donationStatus == "deny" {
			SendEmail("failed", donation)
			h.DonationRepository.Notification("failed", donation.ID)
		} else if donationStatus == "cancel" {
			SendEmail("failed", donation)
			h.DonationRepository.Notification("failed", donation.ID)
		} else if donationStatus == "pending" {
			SendEmail("pending", donation)
			h.DonationRepository.Notification("pending", donation.ID)
		}

		w.WriteHeader(http.StatusOK)
	}
}

func SendEmail(status string, transaction models.Donation) {
	// var CONFIG_SMTP_HOST = "smtp.gmail.com"
	// var CONFIG_SMTP_PORT = 587
	// var CONFIG_SENDER_NAME = "DeweTour <hydrilla.salim@gmail.com>"
	// var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL_SYSTEM")
	// var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD_SYSTEM")

	// var tripTitle = transaction.Trip.Title
	// var price = strconv.Itoa(transaction.Total)

	// mailer := gomail.NewMessage()
	// mailer.SetHeader("from", CONFIG_SENDER_NAME)
	// mailer.SetHeader("To", transaction.User.Email)
	// mailer.SetHeader("Subject", "Status Transaksi AMAAAN")
	// mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
	//  <html lang="en">
	//    <head>
	//    <meta charset="UTF-8" />
	//    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
	//    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
	//    <title>Document</title>
	//    <style>
	//      h1 {
	//      color: brown;
	//      }
	//    </style>
	//    </head>
	//    <body>
	//    <h2>Product payment :</h2>
	//    <ul style="list-style-type:none;">
	//      <li>Name : %s</li>
	//      <li>Total payment: Rp.%s</li>
	//      <li>Status : <b>%s</b></li>
	//      <li>Iklan : <b>%s</b></li>
	//    </ul>
	//    </body>
	//  </html>`, tripTitle, price, status, "TOUR SURGA DISKON 99%"))

	// dialer := gomail.NewDialer(
	// 	CONFIG_SMTP_HOST,
	// 	CONFIG_SMTP_PORT,
	// 	CONFIG_AUTH_EMAIL,
	// 	CONFIG_AUTH_PASSWORD,
	// )

	// err := dialer.DialAndSend(mailer)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// log.Println("Pesan Terkirim")
}

func (h *handlerDonation) FindDonation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	donation, err := h.DonationRepository.FindDonation()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: donation}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerDonation) GetDonation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	donation, err := h.DonationRepository.GetDonation(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: donation}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerDonation) GetDonationByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	donation, err := h.DonationRepository.GetDonationByUserID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: donation}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerDonation) UpdateDonation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// userId := int(userInfo["id"].(float64))

	// dataContext := r.Context().Value("dataFile")
	// filename := Path_File + dataContext.(string)

	var status string
	json.NewDecoder(r.Body).Decode(&status)

	fmt.Println(status)

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	donation, err := h.DonationRepository.GetDonation(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// if request.UserID != 0 {
	// 	donation.UserID = request.UserID
	// }
	if status != "" {
		donation.Status = status
	}

	data, err := h.DonationRepository.CheckDonation(donation, donation.ID)
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

func (h *handlerDonation) DeleteDonation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	trip, err := h.DonationRepository.GetDonation(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.DonationRepository.DeleteDonation(trip, id)
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

func (h *handlerDonation) GetDonationByFund(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	donation, err := h.DonationRepository.GetDonationByFund(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: donation}
	json.NewEncoder(w).Encode(response)
}

// func convertResponseDonation(u models.Donation) donationdto.DonationResponse {
// 	return donationdto.DonationResponse{
// 		ID:          u.ID,
// 		DonationAmt: u.DonationAmt,
// 		Status:      u.Status,
// 		UserID:      u.UserID,
// 		User:        u.User,
// 		FundID:      u.FundID,
// 		Fund:        u.Fund,
// 	}
// }
