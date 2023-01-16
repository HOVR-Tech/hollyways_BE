package repositories

import (
	"hollyways/models"

	"gorm.io/gorm"
)

type DonationRepository interface {
	// BookTransaction(transaction models.Transaction) (models.Transaction, error)
	MakeDonation(donation models.Donation, DonationID int) (models.Donation, error)
	FindDonation() ([]models.Donation, error)
	GetDonation(ID int) (models.Donation, error)
	GetDonationByUserID(ID int) ([]models.Donation, error)
	GetDonationByFund(ID int) ([]models.Donation, error)
	GetOneDonation(ID string) (models.Donation, error)
	UpdateDonation(donation models.Donation, ID int) (models.Donation, error)
	Notification(status string, ID int) (models.Donation, error)

	// ADMIN
	CheckDonation(donation models.Donation, ID int) (models.Donation, error)
	DeleteDonation(donation models.Donation, ID int) (models.Donation, error)
}

func RepositoryDonation(db *gorm.DB) *repository {
	return &repository{db}
}

// func (r *repository) BookDonation(donation models.Donation) (models.Donation, error) {
// 	err := r.db.Preload("Trip").Preload("Trip.Country").Preload("User").Create(&donation).Error

// 	return donation, err
// }

func (r *repository) MakeDonation(donation models.Donation, DonationID int) (models.Donation, error) {
	err := r.db.Preload("Fund").Preload("User").Create(&donation).Error

	return donation, err
}

func (r *repository) FindDonation() ([]models.Donation, error) {
	var donation []models.Donation
	err := r.db.Preload("Fund").Preload("User").Find(&donation).Error

	return donation, err
}

func (r *repository) GetDonation(ID int) (models.Donation, error) {
	var donation models.Donation
	err := r.db.Preload("Fund").Preload("Fund.User").Preload("User").First(&donation, ID).Error

	return donation, err
}

func (r *repository) GetDonationByFund(ID int) ([]models.Donation, error) {
	var donation []models.Donation
	err := r.db.Where("fund_id=?", ID).Preload("Fund.User").Preload("User").Find(&donation).Error

	return donation, err
}

func (r *repository) GetDonationByUserID(ID int) ([]models.Donation, error) {
	var donation []models.Donation
	err := r.db.Debug().Where("user_id =?", ID).Preload("Fund").Preload("User").Find(&donation).Error

	return donation, err
}

func (r *repository) GetOneDonation(ID string) (models.Donation, error) {
	var donation models.Donation
	err := r.db.Preload("Fund").Preload("User").First(&donation, ID).Error

	return donation, err
}

func (r *repository) Notification(status string, ID int) (models.Donation, error) {
	var donation models.Donation
	r.db.Debug().Preload("Fund").First(&donation, ID)

	if status != donation.Status && status == "success" {
		var fund models.Fund
		r.db.First(&fund, donation.Fund.ID)
		fund.DonationLimit = fund.DonationLimit - donation.DonationAmt
		r.db.Model(&fund).Updates(fund)
	}

	donation.Status = status

	err := r.db.Preload("Fund").Preload("User").Model(&donation).Updates(donation).Error

	return donation, err
}

func (r *repository) UpdateDonation(donation models.Donation, ID int) (models.Donation, error) {
	err := r.db.Preload("Fund").Preload("User").Model(&donation).Updates(donation).Error

	return donation, err
}

// ADMIN
func (r *repository) CheckDonation(donation models.Donation, ID int) (models.Donation, error) {
	err := r.db.Preload("Fund").Preload("User").Model(&donation).Updates(donation).Error

	return donation, err
}

func (r *repository) DeleteDonation(donation models.Donation, ID int) (models.Donation, error) {
	err := r.db.Delete(&donation, ID).Error

	return donation, err
}
