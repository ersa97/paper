package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/ersa97/paper-test/paginations"
	"github.com/jinzhu/gorm"
)

type Account struct {
	Code      int       `json:"code" gorm:"code"`
	Name      string    `json:"name" gorm:"name"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt NullTime  `json:"deleted_at" gorm:"deleted_at"`
}

func GetAccountDetail(code int, DB *gorm.DB) (Account, error) {

	var account Account

	result := DB.Where("code = ? and deleted_at is null", code).First(&account)
	if result.Error != nil {
		return Account{}, errors.New("something is wrong while getting account data")
	}
	if result.RecordNotFound() {
		return Account{}, errors.New("data not found")
	}
	return account, nil
}

func GetAccountList(limit, page int, DB *gorm.DB) (*paginations.Paginator, error) {
	var acc Account

	countAcc := DB.Where("deleted_at is null")
	getAcc := DB.Where("deleted_at is null")
	if getAcc.Error != nil {
		return nil, getAcc.Error
	}
	paginator := paginations.Paging(&paginations.Param{
		DbSelect: getAcc,
		DbCount:  countAcc,
		Page:     page,
		Limit:    limit,
		OrderBy:  []string{"code asc"},
	}, &acc)

	return paginator, nil
}

func AddAccount(data Account, DB *gorm.DB) (Account, error) {

	var name Account
	resultGet := DB.Debug().Where("name=? and deleted_at is null", data.Name).Find(&name)
	if !resultGet.RecordNotFound() {
		return Account{}, errors.New("name already in use")
	}

	create := DB.Create(&data)
	if create.Error != nil {
		return data, create.Error
	}
	return data, nil

}
func UpdateAccount(data Account, DB *gorm.DB) (Account, error) {
	var acc Account
	resultGet := DB.Where("name=? and deleted_at is null", data.Name).Find(&acc)
	if !resultGet.RecordNotFound() {
		return Account{}, errors.New("name already in use")
	}

	acc.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", time.Now().Local().Format("2006-01-02 15:04:05"))

	update := DB.Where("code = ?", data.Code).Update(&data)
	if update.Error != nil {
		return Account{}, errors.New("something is wrong while updating account data")
	}
	return data, nil
}

func DeleteAccount(data Account, DB *gorm.DB) error {

	var acc Account
	Get := DB.Where("code = ? and deleted_at is null", data.Code).First(&acc)
	if Get.RecordNotFound() {
		return errors.New("data not found")

	}

	date, _ := time.Parse("2006-01-02 15:04:05", time.Now().Local().Format("2006-01-02 15:04:05"))

	acc.UpdatedAt = date
	acc.DeletedAt = NullTime{sql.NullTime{Time: date, Valid: true}}

	delete := DB.Where("code = ?").Update(acc.DeletedAt)
	if delete.Error != nil {
		return errors.New("something is wrong while deleting account data")
	}

	return nil
}
