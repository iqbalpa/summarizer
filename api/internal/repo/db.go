package repo

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "iqbalpahlevi"
	dbname   = "summarizer"
	port     = "5432"
)

func ConnectDb() *gorm.DB {
	fmt.Println("Connecting to DB...")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbname,
		port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect to db")
		return nil
	}
	fmt.Println("Connected to DB successfully!")
	return db
}
