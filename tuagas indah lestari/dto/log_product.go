package dto

type CreateFromIncreaseParam struct {
	ProductId 		int		`json:"product_id" example:"1"`
	ProductCode 	string  `json:"product_code" example:"001"`
	ProductTypeId 	int		`json:"product_type_id" example:"1"`
	ProductName 	string  `json:"product_name" binding:"max=255" example:"Product 1"`
	ProductQty 		float32 	`json:"product_qty" binding:"min=0" example:"10"`
	ProductPrice 	float32 `json:"product_price" binding:"min=0" example:"10.00"`
}

type CreateFromDecreaseParam struct {
	ProductId 		int `json:"product_id" example:"1"`
	ProductQty 		float32 `json:"product_qty" binding:"min=0" example:"10"`
}
