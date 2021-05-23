package services

import (
	"apiServer/api/entity"
	"apiServer/api/util"
	"math/rand"
	"testing"
	"time"

	"github.com/go-easylog/el"
)

func TestPersistRoom(t *testing.T) {
	el.SetLogLevel(el.TRACE)

	dbCon, err := util.Connect()
	defer dbCon.Close()
	if err != nil {
		el.Errorf("予約登録処理の接続処理に失敗ました。")
		t.Error("error")
	}

	rand.Seed(time.Now().UnixNano())
	roomTypes := []entity.RoomType{entity.FORROOM, entity.ATHOME}
	roomStatus := []entity.RoomStatus{entity.USING, entity.ISSTOPPED, entity.ISDELETE}

	for i := 0; i < 5; i++ {
		room := entity.Room{}
		room.RoomName = "aaaa"
		room.Address = "千葉県船橋市"

		roomTypesIndex := rand.Intn(len(roomTypes))
		room.Type = roomTypes[roomTypesIndex]

		roomStatusIndex := rand.Intn(len(roomStatus))
		room.Status = roomStatus[roomStatusIndex]
		room.Note = "aaaaaa"
		room.Created = time.Now()
		room.Modified = time.Now()
		persistRoom(dbCon, room)
	}
}
