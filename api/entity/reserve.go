package entity

import "time"

//Reserve は予約データ
type Reserve struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	RoomID           int       `json:"room_id"`
	StartReserveDate time.Time `json:"start_reserve_date"`
	EndReserveDate   time.Time `json:"end_reserve_date"`
	Note             string    `json:"note"`
	Created          time.Time `json:"created"`
	Modified         time.Time `json:"modified"`
}
