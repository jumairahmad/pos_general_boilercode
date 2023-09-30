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
	WarehouseID   uint      `json:"warehouse_id"`
	Quantity      int       `json:"quantity"`
	CustomerID    uint      `json:"customer_id"`
	User          User      `gorm:"foreignKey:ID" json:"user"`
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
