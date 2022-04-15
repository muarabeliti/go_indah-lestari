package dto

type ProductTypeParam struct {
	Name	string `json:"name" binding:"required,max=255" example:"Type 1"`
}

type ProductTypeUpdateParam struct {
	Name   	string `json:"name" binding:"required,max=255" example:"Type 1"`
}