package Models

import (
	"POS/Config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Roles struct {
	gorm.Model
	ID         uint   `gorm:"primary_key" id:"id"`
	Name       string `json:"name"`
	Permission string `json:"permission"`
}

func (roles *Roles) TableName() string {
	return "roles"
}

func GetAllRoles(roles *[]Roles) (err error) {
	if err = Config.DB.Find(roles).Error; err != nil {
		return err
	}
	return nil
}

func AddNewRoles(roles *Roles) (err error) {
	if err = Config.DB.Create(roles).Error; err != nil {
		return err
	}
	return nil
}

func GetOneRoles(roles *Roles, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(roles).Error; err != nil {
		return err
	}
	return nil
}

func PutOneRoles(roles *Roles, id string) (err error) {
	fmt.Println(roles)
	Config.DB.Save(roles)
	return nil
}

func DeleteRoles(roles *Roles, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(roles)
	return nil
}
