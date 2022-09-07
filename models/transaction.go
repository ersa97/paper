package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/ersa97/paper-test/paginations"
	"github.com/jinzhu/gorm"
)

type Transaction struct {
	TrxId     int       `json:"trx_id" gorm:"trx_id"`
	Code      int       `json:"code" gorm:"code"`
	UserId    int       `json:"user_id" gorm:"user_id"`
	Amount    float64   `json:"amount" gorm:"amount"`
	Status    int       `json:"status" gorm:"status"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt NullTime  `json:"deleted_at" gorm:"deleted_at"`
}

func AddTransaction(data Transaction, DB *gorm.DB) (Transaction, error) {

	create := DB.Create(&data)
	if create.Error != nil {
		return data, create.Error
	}
	return data, nil
}

func GetDetailTransaction(trxid int, DB *gorm.DB) (Transaction, error) {
	var trx Transaction

	result := DB.Where("trx_id = ? and deleted_at is null", trxid).First(&trx)
	if result.Error != nil {
		return Transaction{}, errors.New("something is wrong while getting account data")
	}
	if result.RecordNotFound() {
		return Transaction{}, errors.New("data not found")
	}
	return trx, nil
}

func GetListTransaction(limit, page, code int, DB *gorm.DB) (*paginations.Paginator, error) {
	var trx Transaction

	counttrx := DB.Where("code = ? and deleted_at is null", code)
	gettrx := DB.Where("code = ? deleted_at is null", code)
	if gettrx.Error != nil {
		return nil, gettrx.Error
	}
	paginator := paginations.Paging(&paginations.Param{
		DbSelect: gettrx,
		DbCount:  counttrx,
		Page:     page,
		Limit:    limit,
		OrderBy:  []string{"created_at asc"},
	}, &trx)

	return paginator, nil
}

func UpdateTransaction(data Transaction, DB *gorm.DB) (Transaction, error) {
	var trx Transaction

	trx.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", time.Now().Local().Format("2006-01-02 15:04:05"))

	update := DB.Where("trx_id = ? and deleted_at is null", data.TrxId).Update(&data)
	if update.Error != nil {
		return Transaction{}, errors.New("something is wrong while updating transaction data")
	}
	return data, nil

}

func DeleteTransaction(data Transaction, DB *gorm.DB) error {
	var trx Transaction
	Get := DB.Where("trx_id = ? and deleted_at is null", data.TrxId).First(&trx)
	if Get.RecordNotFound() {
		return errors.New("data not found")

	}

	date, _ := time.Parse("2006-01-02 15:04:05", time.Now().Local().Format("2006-01-02 15:04:05"))

	trx.UpdatedAt = date
	trx.DeletedAt = NullTime{sql.NullTime{Time: date, Valid: true}}

	delete := DB.Where("code = ?").Update(trx.DeletedAt)
	if delete.Error != nil {
		return errors.New("something is wrong while deleting account data")
	}

	return nil

}
