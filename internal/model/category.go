package model

import "time"

type Category struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime:true" json:"created_at"`
}
