package Models

import (
	"POS/Config"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Warehouse struct {
	gorm.Model
	ID        uint   `gorm:"primary_key" id:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	ContactNo string `json:"contact"`
	Email     int    `json:"email"`
	Image     string `json:"image"`
}

func (warehouse *Warehouse) GetTotalQuantity() int {
	var quantity Quantity
	Config.DB.Find("WarehouseID=?", warehouse.ID).First(&quantity)
	return quantity.TotalQuantity
}

func (warehouse *Warehouse) TableName() string {
	return "warehouse"
}

func GetAllWarehouse(warehouse *[]Warehouse) (err error) {
	if err = Config.DB.Find(warehouse).Error; err != nil {
		return err
	}
	return nil
}

func AddNewWarehouse(warehouse *Warehouse) (err error) {
	if err = Config.DB.Create(warehouse).Error; err != nil {
		return err
	}
	return nil
}

func GetOneWarehouse(warehouse *Warehouse, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(warehouse).Error; err != nil {
		return err
	}
	return nil
}

func PutOneWarehouse(warehouse *Warehouse, id string) (err error) {
	fmt.Println(warehouse)
	Config.DB.Save(warehouse)
	return nil
}

func DeleteWarehouse(warehouse *Warehouse, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(warehouse)
	return nil
}
