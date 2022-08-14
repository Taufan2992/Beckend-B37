package transactiondto

type TransactionRequest struct {
	ProductID int    `json:"product_id"`
	BuyerID   int    `json:"buyer_id"`
	SellerID  int    `json:"seller_id"`
	Price     int    `json:"price" gorm:"type: varchar(255)"`
	Status    string `json:"status" gorm:"type: varchar(255)"`
}
