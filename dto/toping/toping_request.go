package topingdto

type TopingRequest struct {
	Name  string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Desc  string `json:"desc" gorm:"type:text" form:"desc"`
	Price int    `json:"price" form:"price" gorm:"type: int"`
}
