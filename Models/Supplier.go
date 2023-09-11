package Models

import (
	"POS/Config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Supplier struct {
	gorm.Model
	ID        uint   `gorm:"primary_key" id:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	ContactNo string `json:"contact"`
}

func (supplier *Supplier) TableName() string {
	return "supplier"
}

func GetAllSupplier(supplier *[]Supplier) (err error) {
	if err = Config.DB.Find(supplier).Error; err != nil {
		return err
	}
	return nil
}

func AddNewSupplier(supplier *Supplier) (err error) {
	if err = Config.DB.Create(supplier).Error; err != nil {
		return err
	}
	return nil
}

func GetOneSupplier(supplier *Supplier, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(supplier).Error; err != nil {
		return err
	}
	return nil
}

func PutOneSupplier(supplier *Supplier, id string) (err error) {
	fmt.Println(supplier)
	Config.DB.Save(supplier)
	return nil
}

func DeleteSupplier(supplier *Supplier, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(supplier)
	return nil
}
