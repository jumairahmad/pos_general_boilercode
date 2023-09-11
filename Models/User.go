package Models

import (
	"POS/Config"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uint   `gorm:"primary_key" id:"id"`
	Name         string `json:"name"`
	Email        string `json:"author"`
	Phone        string `json:"category"`
	Cnic         int    `json:"cnic"`
	UserName     string `json:"username"`
	Password     string `json:"password"`
	ProfileImage string `json:"profileimage"`
	RoleId       uint   `json:"role_id"`
}

func (user *User) Role() Roles {
	var roles Roles
	Config.DB.Where("id = ?", user.RoleId).First(&roles)
	return roles
}

func (user *User) TableName() string {
	return "user"
}

func GetAllUser(user *[]User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func AddNewUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetOneUser(user *User, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func PutOneUser(user *User, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}

func DeleteUser(user *User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}
