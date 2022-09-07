package models

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	Id        int       `json:"id" gorm:"id"`
	Name      string    `json:"name" gorm:"name"`
	Email     string    `json:"email" gorm:"email"`
	Password  string    `json:"password" gorm:"password"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt NullTime  `json:"deleted_at" gorm:"deleted_at"`
}

type NullTime struct {
	sql.NullTime
}

func GetUser(param interface{}, DB *gorm.DB) (User, error) {

	var user User

	GetUser := DB.Debug().Where("email =? AND deleted_at is null", param).First(&user)
	if GetUser.Error != nil {
		return user, GetUser.Error
	}
	return user, nil
}

func AddUser(data User, DB *gorm.DB) (User, error) {

	log.Println("create user")

	date, _ := time.Parse("2006-01-02 15:04:05", time.Now().Local().Format("2006-01-02 15:04:05"))
	data.DeletedAt = NullTime{sql.NullTime{Time: date, Valid: false}}
	result := DB.Create(&data)
	if result.Error != nil {
		return data, result.Error
	}

	return data, nil

}

func CheckUser(email string, DB *gorm.DB) (bool, error) {

	var user User

	GetUser := DB.Raw("select id, name ,email, password , created_at,  updated_at ,deleted_at from users where email = ? ", email).First(&user)
	if GetUser.RowsAffected > 0 {
		return false, errors.New("email already used")
	} else {
		return true, nil
	}

}
