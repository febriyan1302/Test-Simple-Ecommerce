package model

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type User struct {
	ID       int    `gorm:"primaryKey" sql:"AUTO_INCREMENT" json:"id"`
	Email    string `gorm:"unique;not null;type:varchar(100)" json:"email"`
	Password string `gorm:"not null;type:varchar(50)" json:"password"`
	Name     string `gorm:"not null;type:varchar(100)" json:"name"`
}
