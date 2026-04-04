package service

type Model struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	OrderID  uint   `gorm:"primaryKey" json:"order_id"`
	NickName string `gorm:"column" json:"nick_name"`
	Region   string `gorm:"column" json:"region"`
	IsDone   bool   `gorm:"column" json:"is_done"`
}
