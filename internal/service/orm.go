package service

type Model struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	OrderID  uint   `gorm:"primaryKey" json:"order_id"`
	NickName string `gorm:"column:nick_name" json:"nick_name"`
	Region   string `gorm:"column:region" json:"region"`
	IsDone   bool   `gorm:"column:is_done" json:"is_done"`
}
