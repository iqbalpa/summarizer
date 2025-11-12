package repo

import (
	"fmt"
	"summarizer/internal/model"

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

func MigrateDb(db *gorm.DB) {
	fmt.Println("Delete existing data...")
	// db.Where("1 = 1").Delete(&model.User{})
	db.Where("1 = 1").Delete(&model.Summary{})
	db.Where("1 = 1").Delete(&model.Job{})
	fmt.Println("Migrating a new table...")
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Summary{})
	db.AutoMigrate(&model.Job{})
}
