package Models

import (
	"POS/Config"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Stock struct {
	gorm.Model
	ID          uint      `gorm:"primary_key" id:"id"`
	Quantity    int       `json:"quantity"`
	Date        time.Time `json:"name"`
	RecievedBy  uint      `json:"recived_by"`
	SellerPrice string    `json:"seller_price"`
	Warehouse   Warehouse `gorm:"foreignKey:ID" json:"warehouse"`
	InvoiceNo   string    `json:"invoice_no"`
	Supplier    Supplier  `gorm:"foreignKey:ID" json:"suppiler"`
}

func (stock *Stock) TableName() string {
	return "stock"
}

func GetAllStock(stock *[]Stock) (err error) {
	if err = Config.DB.Find(stock).Error; err != nil {
		return err
	}
	return nil
}

func AddNewStock(stock *Stock) (err error) {
	if err = Config.DB.Create(stock).Error; err != nil {
		return err
	}
	return nil
}

func GetOneStock(stock *Stock, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(stock).Error; err != nil {
		return err
	}
	return nil
}

func PutOneStock(stock *Stock, id string) (err error) {
	fmt.Println(stock)
	Config.DB.Save(stock)
	return nil
}

func DeleteStock(stock *Stock, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(stock)
	return nil
}
