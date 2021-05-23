package services

import (
	"apiServer/api/entity"

	"github.com/go-easylog/el"
	"github.com/jinzhu/gorm"
)

//persistRoom は部屋の登録
func persistRoom(dbCon *gorm.DB, room entity.Room) error {

	el.Info("部屋の登録を開始します。")
	dbCon = dbCon.Create(&room)

	if dbCon.Error != nil {
		el.Errorf("予約登録処理の登録処理に失敗しました %s", dbCon.Error)
	}

	return nil
}
