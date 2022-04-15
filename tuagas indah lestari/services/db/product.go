package services

import (
	"fmt"
	"myapp/models"

	"gorm.io/gorm"
)

func (d *db) ProductCreate(input models.Product) (*models.Product, error) {

	err := d.Tx.Model(&models.Product{}).Create(&input).Error
	if err != nil {
		d.Tx.Rollback()
		panic(err)
	}

	return d.ProductGetByID(input.ID)
}

func (d *db) ProductGetByID(id int) (*models.Product, error) {

	var product models.Product
	err := d.Tx.Model(&models.Product{}).Where("id = ?", id).Take(&product).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			d.Tx.Rollback()
			panic(err)
		} else {
			d.Tx.Rollback()
			return nil, fmt.Errorf("product not found")
		}
	}

	return &product, nil
}

func (d *db) ProductGetByIDForLogProduct(id int) (*models.Product, error) {

	var product models.Product
	result := d.Tx.Model(&models.Product{}).Where("id = ?", id).First(&product)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("product not found")
	}

	return &product, nil
}

func (d *db) ProductGetByCode(code string) (*models.Product, error) {

	var product models.Product
	err := d.Tx.Model(&models.Product{}).Where("code = ?", code).Take(&product).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			d.Tx.Rollback()
			panic(err)
		} else {
			d.Tx.Rollback()
			return nil, fmt.Errorf("product not found")
		}
	}

	return &product, nil
}

func (d *db) ProductGetAll() ([]*models.Product, error) {

	var products []*models.Product
	err := d.Tx.Model(&models.Product{}).Find(&products).Error
	if err != nil {
		d.Tx.Rollback()
		panic(err)
	}

	return products, nil
}

func (d *db) ProductUpdate(input models.Product) (*models.Product, error) {

	err := d.Tx.Model(&models.Product{}).Where("id = ?", input.ID).Updates(&input).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			d.Tx.Rollback()
			panic(err)
		} else {
			d.Tx.Rollback()
			return nil, fmt.Errorf("product not found")
		}
	}

	return d.ProductGetByID(input.ID)
}

func (d *db) ProductDelete(id int) (*int, error) {

	err := d.Tx.Model(&models.Product{}).Where("id = ?", id).Delete(&models.Product{}).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			d.Tx.Rollback()
			panic(err)
		} else {
			d.Tx.Rollback()
			return nil, fmt.Errorf("product not found")
		}
	}

	return &id, nil
}
