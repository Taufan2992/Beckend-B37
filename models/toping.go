package models

import "time"

// User model struct
type Toping struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Name      string    `json:"name" gorm:"type: varchar(255)"`
	Desc      string    `json:"desc" gorm:"type: varchar(255)"`
	Price     int       `json:"price" gorm:"type: int"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
