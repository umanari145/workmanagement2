package entity

import "time"

const (
	//USING 使う
	USING RoomStatus = 10
	//ISSTOPPED 停止
	ISSTOPPED RoomStatus = 50
	//ROOMISDELETE 削除
	ROOMISDELETE RoomStatus = 99

	//FORROOM 通勤
	FORROOM RoomType = 10
	//ATHOME 在宅
	ATHOME RoomType = 20
)

//Room は部屋データ
type Room struct {
	ID       int        `json:"id"`
	RoomName string     `json:"room_name"`
	Address  string     `json:"address"`
	Type     RoomType   `json:"type"`
	Status   RoomStatus `json:"status"`
	Note     string     `json:"note"`
	Created  time.Time  `json:"created"`
	Modified time.Time  `json:"modified"`
}

//RoomStatus は部屋の状態
type RoomStatus int

//RoomType は部屋の系統
type RoomType int
