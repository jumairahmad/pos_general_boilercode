package Models

import (
	"POS/Config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Quantity struct {
	gorm.Model
	ID            uint        `gorm:"primary_key" id:"id"`
	TotalQuantity int         `json:"total_quantity"`
	Warehouse     []Warehouse ` gorm:"foreignKey:ID" json:"warehouse" `
	Product       []Product   `gorm:"foreignKey:ID" json:"product" `
}

func (quantity *Quantity) TableName() string {
	return "quantity"
}

func GetAllQuantity(quantity *[]Quantity) (err error) {
	if err = Config.DB.Find(quantity).Error; err != nil {
		return err
	}
	return nil
}

func AddNewQuantity(quantity *Quantity) (err error) {
	if err = Config.DB.Create(quantity).Error; err != nil {
		return err
	}
	return nil
}

func GetOneQuantity(quantity *Quantity, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(quantity).Error; err != nil {
		return err
	}
	return nil
}

func PutOneQuantity(quantity *Quantity, id string) (err error) {
	fmt.Println(quantity)
	Config.DB.Save(quantity)
	return nil
}

func DeleteQuantity(quantity *Quantity, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(quantity)
	return nil
}
