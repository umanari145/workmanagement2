package entity

import "time"

const (
	//REGULAR 正規
	REGULAR UserType = 10
	//EXTERNAL 外部
	EXTERNAL UserType = 20
	//ISOLATION 隔離
	ISOLATION UserType = 99

	//ISHOME 在宅
	ISHOME WorkingType = 10
	//COMMUTING 通勤
	COMMUTING WorkingType = 20

	//CHIBA 千葉
	CHIBA AreaType = 10
	//FUKUOKA 福岡
	FUKUOKA AreaType = 20

	//USERNORMAL は正常状態
	USERNORMAL UserStatus = 10
	//USERISDELETE は削除状態
	USERISDELETE UserStatus = 99
)

//User はユーザーデータ
type User struct {
	ID                    int         `json:"id"`
	UserName              string      `json:"username"`
	JapaneseName          string      `json:"japanese_name"`
	Email                 string      `json:"email"`
	Password              string      `json:"password"`
	UserType              UserType    `json:"user_type"`
	WorkingType           WorkingType `json:"working_type"`
	AreaType              AreaType    `json:"area_type"`
	Status                UserStatus  `json:"status"`
	IP                    string      `json:"ip"`
	LastAnnounceCheckTime time.Time   `json:"last_announce_check_time"`
	LastAccessTime        time.Time   `json:"last_access_time"`
	Created               time.Time   `json:"created"`
	Modified              time.Time   `json:"modified"`
}

//UserType はユーザーのタイプ
type UserType int

//WorkingType は働き方
type WorkingType int

//AreaType は働く場所
type AreaType int

//UserStatus はユーザーの状態
type UserStatus int
