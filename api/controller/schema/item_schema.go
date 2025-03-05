package schema

type CreateItemInput struct {
	Name        string                 `json:"name" binding:"required",min=2"`
	Price       uint                   `json:"price" binding:"required",min=1,max=999999`
	Description map[string]interface{} `gorm:"type:jsonb"`
}
type UpdateItemInput struct {
	Name        *string                `json:"name" binding:"required",min=2"`
	Price       *uint                  `json:"price" binding:"required",min=1,max=999999`
	Description map[string]interface{} `gorm:"type:jsonb"`
	SoldOut     *bool                  `json:"soldOut"`
}
