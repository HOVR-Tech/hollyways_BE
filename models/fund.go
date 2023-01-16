package models

type Fund struct {
	ID            int    `json:"id"`
	Title         string `json:"title" gorm:"type: varchar(255)"`
	Days          int    `json:"days" gorm:"type:int"`
	DonationLimit int    `json:"donation_limit" gorm:"type:int"`
	Description   string `json:"description" gorm:"type:longtext"`
	UserID        int    `json:"user_id"`
	User          User   `json:"user"`
	Image         string `json:"image" gorm:"type: varchar(255)"`
}
