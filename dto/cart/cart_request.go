package cartdto

type CartRequest struct {
	ProductID     int `json:"product_id"`
	TransactionID int `json:"transaction_id"`
	Qty           int `json:"qty" form:"qty"`
	SubAmount     int `json:"subamount"`
}
