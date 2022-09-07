package models

import (
	"database/sql"
	"errors"

	"github.com/jinzhu/gorm"
)

type UserToken struct {
	Id     int        `json:"id" gorm:"id"`
	UserId int        `json:"userid" gorm:"userid"`
	Uuid   nullString `json:"uuid" gorm:"uuid"`
}

type nullString struct {
	sql.NullString
}

func AddUserToken(param int, DB *gorm.DB) error {
	var token UserToken

	token.UserId = param
	token.Uuid = nullString{sql.NullString{String: "", Valid: false}}
	resulttoken := DB.Create(&token)

	if resulttoken.Error != nil {
		return resulttoken.Error
	}
	return nil
}

func GetUserToken(userid int, DB *gorm.DB) (string, error) {

	var token UserToken

	result := DB.Where("user_id = ?", userid).First(&token)
	if result.Error != nil {
		return "", errors.New("error while finding user token")
	}
	if result.RecordNotFound() {
		return "", errors.New("data not found")
	}
	return token.Uuid.String, nil
}

func UpdateUserToken(userid int, uid string, DB *gorm.DB) error {
	var token UserToken

	token.UserId = userid

	if uid == "nil" {
		token.Uuid = nullString{sql.NullString{String: "", Valid: false}}
		result := DB.Exec("UPDATE user_tokens SET uuid = ? WHERE user_id = ?", token.Uuid, token.UserId)
		if result.Error != nil {
			return result.Error
		}
	}
	token.Uuid = nullString{sql.NullString{String: uid, Valid: true}}
	result := DB.Debug().Where("user_id=?", userid).Model(&token).Update(token)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteUserToken(userid int, DB *gorm.DB) error {

	var token UserToken

	token.Id = userid

	result := DB.Where("user_id=?", userid).Delete(&token)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
