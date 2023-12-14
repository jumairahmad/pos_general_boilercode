package Models

import (
	"POS/Config"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Orders struct {
	gorm.Model
	ID            uint      `gorm:"primary_key" id:"id"`
	Warehouse     Warehouse `gorm:"foreignKey:ID" json:"warehouse"`
	Quantity      int       `json:"quantity"`
	Customer      User      `gorm:"foreignKey:ID" json:"customer"`
	Seller        User      `gorm:"foreignKey:ID" json:"user"`
	Date          time.Time `json:"date"`
	PricePerLiter int       `json:"price_per_liter"`
	TotalBill     int       `json:"total_bill"`
	Discount      int       `json:"discount"`
	Status        string    `json:"status"`
	Product       []Product ` gorm:"foreignKey:ID" json:"product_id"`
	OrderNO       string    `json:"order_no"`
}

func (orders *Orders) TableName() string {
	return "orders"
}

func (order *Orders) getSales() int {
	var orders []Orders
	Config.DB.Where("id = ?", order.ID).Find(&orders)
	TotalSales := 0
	for _, v := range orders {
		TotalSales += v.TotalBill
	}
	return 0
}

func GetAllSales() int {
	var orders []Orders
	Config.DB.Find(&orders)
	TotalSales := 0
	for _, v := range orders {
		TotalSales += v.TotalBill
	}
	return TotalSales
}

func GetAllOrders(orders *[]Orders) (err error) {
	if err = Config.DB.Find(orders).Error; err != nil {
		return err
	}
	return nil
}

func AddNewOrders(orders *Orders) (err error) {
	if err = Config.DB.Create(orders).Error; err != nil {
		return err
	}
	return nil
}

func GetOneOrders(orders *Orders, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(orders).Error; err != nil {
		return err
	}
	return nil
}

func PutOneOrders(orders *Orders, id string) (err error) {
	fmt.Println(orders)
	Config.DB.Save(orders)
	return nil
}

func DeleteOrders(orders *Orders, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(orders)
	return nil
}
