package conn

import (
	"fmt"
	"github.com/alfaysal/go-pet-project/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/url"
)

var DB *gorm.DB

func ConnectDefaultDB() error {
	cfg := config.DefaultDB()

	uri := url.URL{
		Scheme: "postgres",
		Host:   cfg.Host,
		Path:   cfg.Name,
		User:   url.UserPassword(cfg.Username, cfg.Password),
	}

	db, err := gorm.Open(postgres.Open(uri.String()), &gorm.Config{})
	DB = db
	fmt.Println(uri.String())
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	return nil
}

func GetDefaultDB() *gorm.DB {
	return DB.Debug()
}
