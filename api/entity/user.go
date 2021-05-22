package entity

//User はユーザーデータ
type User struct {
	ID           int    `json:"id"`
	ChatGirlID   string `json:"chatgirl_id"`
	UserName     string `json:"username"`
	JapaneseName string `json:"japanese_name"`
	//まだ作成途中
}
