package models

import "time"

type Donation struct {
	ID          int       `json:"id"`
	DonationAmt int       `json:"donation_amt" `
	Status      string    `json:"status" gorm:"type: varchar(255)"`
	UserID      int       `json:"user_id" `
	User        User      `json:"user"`
	FundID      int       `json:"fund_id"`
	Fund        Fund      `json:"fund"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
