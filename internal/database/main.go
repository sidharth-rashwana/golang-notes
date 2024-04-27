package internal

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	var db *gorm.DB
	var err error
	for attempt := 1; attempt <= 5; attempt++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		fmt.Printf("Attempt %d: Unable to connect to the database. Retrying...\n", attempt)
		time.Sleep(time.Duration(attempt) * time.Second)
	}
	if err != nil {
		fmt.Println("Failed to connect to the database after multiple attempts.")
		return nil
	}

	return db
}
