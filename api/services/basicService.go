package services

import (
	"apiServer/api/util"
	"fmt"

	"github.com/go-easylog/el"
)

//getReserve は
func getEntity(tableName string, entity interface{}) interface{} {

	dbCon, _ := util.Connect()
	defer dbCon.Close()

	el.Info(fmt.Printf("get %s start", tableName))

	dbCon.Table(tableName).Find(&entity, 1)

	return entity
}

//persistEntity は登録
func persistEntity(tableName string, entity interface{}) error {

	dbCon, err := util.Connect()
	defer dbCon.Close()
	if err != nil {
		el.Errorf("登録処理の接続処理に失敗ました。")
		return nil
	}
	el.Info(fmt.Sprintf("persist %s start", tableName))
	dbCon = dbCon.Create(entity)
	if dbCon.Error != nil {
		el.Errorf("登録処理の登録処理に失敗しました %s", dbCon.Error)
	}
	return nil
}
