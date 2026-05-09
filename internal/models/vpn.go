package models

type VpnConfigModel struct {
	OrderID   uint      `gorm:"primaryKey" json:"order_id"`
	UserID    uint      `gorm:"index"`
	User      UserModel `gorm:"foreignKey:UserID"`
	Region    string    `gorm:"column:region" json:"region"`
	IP        string    `gorm:"unique" json:"ip"`
	IsActive  bool      `gorm:"column:is_active" json:"is_active"`
	ExpiresAt int64     `gorm:"index" json:"expires_at"`
}
