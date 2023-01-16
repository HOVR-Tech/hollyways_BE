package models

type User struct {
	ID       int        `json:"id"`
	Name     string     `json:"name" gorm:"type: varchar(255)"`
	Email    string     `json:"email" gorm:"type: varchar(255)"`
	Password string     `json:"password" gorm:"type: varchar(255)"`
	Phone    string     `json:"phone" gorm:"type: varchar(255)"`
	Donation []Donation `json:"donation"`
	Image    string     `json:"image" gorm:"type: varchar(255)"`
	Role     string     `json:"role" gorm:"type: varchar(255)"`
}
