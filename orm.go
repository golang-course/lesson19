package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Phone struct {
	gorm.Model
	Vendor    string `gorm:"not null"`
	ModelName string `gorm:"not null"`
	Price     int
}

func main() {
	const pgConfig = "user=postgres password=postgres dbname=mydb sslmode=disable port=5432"
	db, err := gorm.Open(postgres.Open(pgConfig), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}

	err = db.AutoMigrate(&Phone{})
	if err != nil {
		log.Println(err)
		return
	}

}
