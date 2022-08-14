package models

import "time"

// User model struct
type Transaction struct {
  ID          	int				 	 `json:"id" gorm:"primary_key:auto_increment"`
  ProductID 	int				 	 `json:"product_id"`
  Product   	ProductResponse  	 `json:"product"`
  BuyerID		int				 	 `json:"buyer_id"`
  Buyer     	UsersProfileResponse `json:"buyer"`
  SellerID	  	int					 `json:"seller_id"`
  Seller    	UsersProfileResponse `json:"seller"`
  Price			int					 `json:"price" gorm:"type: varchar(255)"`
  Status		string				 `json:"status" gorm:"type: varchar(255)"`
  CreatedAt 	time.Time			 `json:"-"`
  UpdatedAt 	time.Time			 `json:"-"`
}