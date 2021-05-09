package services

import (
	"apiServer/api/entity"
	"apiServer/api/util"
	"time"

	"github.com/go-easylog/el"
)

//getReserve は
func getReserve() entity.Reserve {

	dbCon, _ := util.Connect()
	defer dbCon.Close()

	el.Info("getReserve start")
	reserve := entity.Reserve{}

	dbCon.Table("reserves").Find(&reserve, 1)

	return reserve
}

//persistReserve は予約の登録
func persistReserve() error {

	dbCon, err := util.Connect()
	defer dbCon.Close()
	if err != nil {
		el.Errorf("予約登録処理の接続処理に失敗ました。")
		return nil
	}

	el.Info("persistReserve start")
	reserve := entity.Reserve{}
	reserve.RoomID = 1
	reserve.UserID = 1
	reserve.StartReserveDate = time.Now()
	reserve.EndReserveDate = time.Now()
	reserve.Created = time.Now()
	reserve.Modified = time.Now()

	dbCon = dbCon.Create(&reserve)

	if dbCon.Error != nil {
		el.Errorf("予約登録処理の登録処理に失敗しました %s", dbCon.Error)
	}
	return nil
}
