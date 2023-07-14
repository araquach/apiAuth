package models

type User struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	StaffID    uint   `json:"staff_id"`
	Email      string `json:"email" gorm:"unique_index"`
	Password   string `json:"password"`
	AdminLevel uint   `json:"admin_level"`
	Token      string `json:"token" gorm:"-"`
}
