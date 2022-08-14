package transactiondto

type TransactionRequest struct {
	Status    string `json:"status" gorm:"type: varchar(255)"`
	BuyerID   int    `json:"buyer_id"`
	ProductID int    `json:"product_id"`
}

type UpdateTransactionRequest struct {
	Status    string `json:"status" gorm:"type: varchar(255)"`
	BuyerID   int    `json:"buyer_id"`
	ProductID int    `json:"product_id"`
}
