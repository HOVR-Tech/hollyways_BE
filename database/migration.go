package database

import (
	"fmt"
	"hollyways/models"
	"hollyways/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Fund{},
		&models.Donation{},
	)

	if err != nil {
		fmt.Println(err)
		panic("MIGRATION FAILED ")
	}

	fmt.Println("MIGRATION SUCCESS 😍😘")
}
