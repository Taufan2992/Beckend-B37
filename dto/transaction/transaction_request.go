package transactiondto

type TransactionRequest struct {
	UserID int `json:"user_id"`
	Amount int `json:"amount"`
}

type UpdateTransactionRequest struct {
	UserID int `json:"user_id"`
	Amount int `json:"amount"`
}
