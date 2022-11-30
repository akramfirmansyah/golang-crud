package initializers

import (
	"fmt"
	"os"

	"github.com/akramfirmansyah/go-crud/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed Connect to Database")
	}
}

func Migrate() {
	DB.AutoMigrate(&models.Mahasiswa{})
	fmt.Println("Migrate Database Successful!")
}
