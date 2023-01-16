package repositories

import (
	"hollyways/models"

	"gorm.io/gorm"
)

type FundRepository interface {
	FindFund() ([]models.Fund, error)
	GetFund(ID int) (models.Fund, error)
	AddFund(fund models.Fund) (models.Fund, error)
	EditFund(fund models.Fund, ID int) (models.Fund, error)
	DeleteFund(fund models.Fund, ID int) (models.Fund, error)
}

func RepositoryFund(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFund() ([]models.Fund, error) {
	var fund []models.Fund
	err := r.db.Find(&fund).Error

	return fund, err
}

func (r *repository) GetFund(ID int) (models.Fund, error) {
	var fund models.Fund
	err := r.db.First(&fund, ID).Error

	return fund, err
}

func (r *repository) AddFund(fund models.Fund) (models.Fund, error) {
	err := r.db.Create(&fund).Error

	return fund, err
}

func (r *repository) EditFund(fund models.Fund, ID int) (models.Fund, error) {
	err := r.db.Model(&fund).Updates(fund).Error

	// err := r.db.Raw("UPDATE Funds SET title=?, country_id=?, accomodation=?, transportation=?,  eat=?, day=?, night=?, date_Fund=?, price=?, quota=?, description=?, image=? WHERE id=?", Fund.Title, Fund.CountryID, Fund.Accomodation, Fund.Transportation, Fund.Eat,  Fund.Day, Fund.Night, Fund.Date_Fund, Fund.Price, Fund.Quota, Fund.Description, Fund.Image, Fund.ID).Scan(&Fund).Error

	return fund, err
}

func (r *repository) DeleteFund(fund models.Fund, ID int) (models.Fund, error) {
	err := r.db.Delete(&fund).Error

	return fund, err
}
