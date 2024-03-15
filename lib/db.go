package lib

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase() (*gorm.DB, error) {
	connectionString := "host=localhost port=5432 user=postgres  password=password dbname=transaksi_db sslmode=disable"
	return gorm.Open(postgres.Open(connectionString), &gorm.Config{})
}
