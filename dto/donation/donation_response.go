package donationdto

import (
	"hollyways/models"
	"time"
)

type DonationResponse struct {
	ID          int         `json:"id"`
	DonationAmt int         `json:"donation_amt"`
	Status      string      `json:"status"`
	UserID      int         `json:"user_id"`
	User        models.User `json:"user"`
	FundID      int         `json:"fund_id"`
	Fund        models.Fund `json:"fund"`
	CreatedAt   time.Time   `json:"-"`
	UpdatedAt   time.Time   `json:"-"`
}
