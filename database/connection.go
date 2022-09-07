package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var sqlDb *sql.DB
var gormDb *gorm.DB

func connectToDb() *sql.DB {
	var err error

	if sqlDb != nil {
		return sqlDb
	}

	log.Println("create pool database connection")

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	sqlDb, err = sql.Open("mysql", dbURL)
	if err != nil {
		log.Println("failed to connect database", err)
		return nil
	}

	sqlDb.SetMaxIdleConns(5)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Minute)

	log.Println("pool database connection is created")

	return sqlDb
}
func poolDbConnection() *gorm.DB {

	var err error

	if sqlDb == nil {
		sqlDb = connectToDb()
		if sqlDb == nil {
			return nil
		}
	}

	log.Println("create connection by gorm framework")

	gormDb, err = gorm.Open("mysql", sqlDb)
	if err != nil {
		log.Println("error on creating gorm connection ", err)
		return nil
	}

	log.Println("gorm connection is created")

	return gormDb
}

func Connection() *gorm.DB {

	return poolDbConnection()
	//return singletonDbConnection()

}
