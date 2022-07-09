package config

import (
	"os"

	"github.com/virhanali/user-management/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
)

var DB *gorm.DB
var err error

func Database() {
	DB, err = gorm.Open(postgres.Open(os.Getenv("db_url")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func Migrate() {
	DB.AutoMigrate(&models.User{})
}
