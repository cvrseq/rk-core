package models


type UserModel struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	NickName string `gorm:"column:nick_name" json:"nick_name"`
}
