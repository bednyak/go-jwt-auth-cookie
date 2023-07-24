package database

import (
	"github.com/bednyak/go-react-jwt-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("andrew:12345678@/gojwtauth"), &gorm.Config{})

	if err != nil {
		println("Could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
