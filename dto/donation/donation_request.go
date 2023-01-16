package donationdto

type DonationRequest struct {
	DonationAmt int    `json:"donation_amt"`
	Status      string `json:"status"`
	FundID      int    `json:"fund_id"`
	UserID      int    `json:"user_id"`
}

// type BookRequest struct {
// 	Qty    int    `json:"qty"`
// 	Total  int    `json:"total"`
// 	Status string `json:"status"`
// 	TripID int    `json:"trip_id"`
// 	UserID int    `json:"user_id"`
// }

type CheckRequest struct {
	Status string `json:"status"`
}
