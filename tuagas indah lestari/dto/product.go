package dto

type ProductParam struct {
	Code			string 	`json:"code" binding:"required,max=255" example:"0001"`
	ProductTypeId	int 	`json:"product_type_id" binding:"required,min=1" example:"1"`
	Name   			string 	`json:"name" binding:"required,max=255" example:"Product 1"`
	Description 	string	`json:"description" binding:"max=255" example:"Product Description"`
	Price 			float32	`json:"price" binding:"required,min=1" example:"10000"`
}

type ProductUpdateParam struct {
	Code			string 	`json:"code" binding:"required,max=255" example:"0001"`
	ProductTypeId	int 	`json:"product_type_id" binding:"required,min=1" example:"1"`
	Name   			string 	`json:"name" binding:"required,max=255" example:"Product 1"`
	Description 	string	`json:"description" binding:"max=255" example:"Product Description"`
	Price 			float32	`json:"price" binding:"required,min=1" example:"10000"`
}