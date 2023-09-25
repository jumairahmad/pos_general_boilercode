package main

import (
	Config "POS/Config"
	Models "POS/Models"
	"POS/Routers"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	//dbURL := "postgres://postgres:1234@localhost:5432/crud"
	dbURL := "gasmanagmentsystem.db"
	// db, err = gorm.Open(dbURL, &gorm.Config{})
	db, err := gorm.Open(sqlite.Open(dbURL), &gorm.Config{})
	Config.DB = db
	if err != nil {
		fmt.Println("status: ", err)
	}
	Config.DB.AutoMigrate(&Models.Book{}, &Models.Orders{}, &Models.Quantity{}, &Models.Roles{}, &Models.Stock{}, &Models.Supplier{}, &Models.User{}, &Models.Warehouse{}, &Models.Product{})

	r := Routers.SetupRouter()
	// running
	r.Run()
}
