package funddto

import "hollyways/models"

type FundResponse struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Days          int    `json:"days"`
	DonationLimit int    `json:"donation_limit"`
	Description   string `json:"description"`
	// UserID        int               `json:"user_id"`
	// User          models.User       `json:"user"`
	Image   string            `json:"image"`
	Donatur []models.Donation `json:"donatur"`
}
