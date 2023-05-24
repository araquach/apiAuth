package models

type User struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Email      string `json:"email" gorm:"unique_index"`
	Password   string `json:"password"`
	AdminLevel uint   `json:"admin_level"`
	Salon      uint   `json:"salon"`
	Token      string `json:"token" gorm:"-"`
}
