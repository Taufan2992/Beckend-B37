package models

import "time"

// User model struct
type Transaction struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	UserID    int                  `json:"user_id"`
	User      UsersProfileResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Amount    int                  `json:"amount"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}

type TransactionUserResponse struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	Amount int `json:"amount"`
}

type TransactionResponse struct {
	UserID int `json:"user_id"`
	Amount int `json:"amount"`
}

func (TransactionUserResponse) TableName() string {
	return "transactions"
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
