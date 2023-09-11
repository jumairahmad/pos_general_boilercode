package Models

import (
	"POS/Config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	ID   uint   `gorm:"primary_key" id:"id"`
	Name string `json:"name"`
	Unit string `json:"unit"`
}

func (product *Product) TableName() string {
	return "product"
}

func GetAllProduct(product *[]Product) (err error) {
	if err = Config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

func AddNewProduct(product *Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func GetOneProduct(product *Product, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

func PutOneProduct(product *Product, id string) (err error) {
	fmt.Println(product)
	Config.DB.Save(product)
	return nil
}

func DeleteProduct(product *Product, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(product)
	return nil
}
