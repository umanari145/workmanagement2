package util

import (
	"fmt"
	"os"

	"github.com/go-easylog/el"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Connect DBへの接続
func Connect() (*gorm.DB, error) {

	el.Info("DB connect start")

	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	DBNAME := os.Getenv("DB_NAME")
	HOST := os.Getenv("DB_HOST")
	PROTOCOL := fmt.Sprintf("tcp(%s:3306)", HOST)

	CONNECT := fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, PROTOCOL, DBNAME)

	el.Infof("DB dms %s", CONNECT)
	db, err := gorm.Open("mysql", CONNECT)
	if err != nil {
		el.Error("DB cannnot connect")
		el.Errorf(" error message %s", err.Error())
		return nil, err
	}
	db.LogMode(true)
	el.Info("DB connect end")

	return db, nil
}
