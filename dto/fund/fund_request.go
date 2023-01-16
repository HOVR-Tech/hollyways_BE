package funddto

type FundRequest struct {
	Title         string `json:"title" `
	Days          int    `json:"days" `
	DonationLimit int    `json:"donation_limit" `
	Description   string `json:"description" `
	Image         string `json:"image"`
}
