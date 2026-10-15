package main

import (
	"fmt"
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
	const pgConfig = "host=localhost user=postgres password=postgres dbname=mydb sslmode=disable port=5432"
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

	// newPhone := Phone{
	// 	Vendor:    "nothingphone",
	// 	ModelName: "model 2",
	// 	Price:     50_000,
	// }

	// res := db.Create(&newPhone)
	// if res.Error != nil {
	// 	log.Println(err)
	// 	return
	// }

	var phones []Phone
	res := db.Where("Vendor = ?", "nothingphone").Find(&phones)
	if res.Error != nil {
		log.Println(err)
		return
	}

	for _, v := range phones {
		fmt.Printf("%s %s %d\n", v.Vendor, v.ModelName, v.Price)
	}

}
