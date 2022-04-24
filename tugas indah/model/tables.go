package model

// TableName overrides the table name used by User to `profiles`
func (User) TableName() string {
	return "users"
}

func (ProductType) TableName() string {
	return "product_type"
}

func (Product) TableName() string {
	return "product"
}

func (LogProduk) TableName() string {
	return "log_product"
}

func (Pivot) TableName() string {
	return "pivot"
}
