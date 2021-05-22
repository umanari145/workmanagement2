package entity

import "time"

const (
	//USING 使う
	USING RoomStatus = 10
	//ISSTOPPED 停止
	ISSTOPPED RoomStatus = 50
	//ISDELETE 削除
	ISDELETE RoomStatus = 99

	//COMMUTING 通勤
	COMMUTING RoomType = 10
	//ISHOME 在宅
	ISHOME RoomType = 20
)

//Room は部屋データ
type Room struct {
	ID       int        `json:"id"`
	RoomName int        `json:"room_name"`
	Address  int        `json:"address"`
	Type     RoomType   `json:"type"`
	Status   RoomStatus `json:"status"`
	Note     time.Time  `json:"note"`
	Created  time.Time  `json:"created"`
	Modified time.Time  `json:"modified"`
}

//RoomStatus は部屋の状態
type RoomStatus int

//RoomType は部屋の系統
type RoomType int
