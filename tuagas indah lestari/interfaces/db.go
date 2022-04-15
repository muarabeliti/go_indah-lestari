package interfaces

import "myapp/models"

type DBInterface interface {

	// User
	UserCreate(models.User) (*models.User, error)
	UserGetByID(int) (*models.User, error)
	UserGetByEmail(string) (*models.User, error)

	// Product
	ProductCreate(models.Product) (*models.Product, error)
	ProductGetByID(int) (*models.Product, error)
	ProductGetByIDForLogProduct(int) (*models.Product, error)
	ProductGetByCode(string) (*models.Product, error)
	ProductGetAll() ([]*models.Product, error)
	ProductUpdate(models.Product) (*models.Product, error)
	ProductDelete(int) (*int, error)

	// ProductType
	ProductTypeCreate(models.ProductType) (*models.ProductType, error)
	ProductTypeGetByID(int) (*models.ProductType, error)
	ProductTypeGetAll() ([]*models.ProductType, error)
	ProductTypeUpdate(models.ProductType) (*models.ProductType, error)
	ProductTypeDelete(int) (*int, error)

	// LogProduct
	LogProductCreate(models.LogProduct) (*models.LogProduct, error)

	// Pivot
	PivotCreate(models.Pivot) (*models.Pivot, error)
}
