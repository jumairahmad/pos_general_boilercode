package main

import (
	Config "POS/Config"
	database "POS/Database"
	Models "POS/Models"
	"POS/Routers"
)

func main() {
	//dbURL := "postgres://postgres:1234@localhost:5432/crud"
	dbURL := "gasmanagmentsystem.db"
	// // db, err = gorm.Open(dbURL, &gorm.Config{})

	db, err := database.DbinstanceHandler(dbURL)

	if err != nil {
		panic("cannot connect to database ")
	}

	Config.DB = db
	db.Migrator().DropTable(&Models.Orders{})
	db.Migrator().DropTable(&Models.Quantity{})
	db.Migrator().DropTable(&Models.Roles{})
	db.Migrator().DropTable(&Models.Stock{})
	db.Migrator().DropTable(&Models.Supplier{})
	db.Migrator().DropTable(&Models.Warehouse{})
	db.Migrator().DropTable(&Models.Product{})
	db.Migrator().DropTable(&Models.User{})

	Config.DB.AutoMigrate(&Models.Orders{}, &Models.Quantity{}, &Models.Roles{}, &Models.Stock{}, &Models.Supplier{}, &Models.User{}, &Models.Warehouse{}, &Models.Product{})

	r := Routers.SetupRouter()
	// running
	r.Run()
}
