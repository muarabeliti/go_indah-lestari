package services

import (
	"fmt"
	"myapp/models"

	"gorm.io/gorm"
)

func (d *db) ProductTypeCreate(input models.ProductType) (*models.ProductType, error) {

	err := d.Tx.Model(&models.ProductType{}).Create(&input).Error
	if err != nil {
		d.Tx.Rollback()
		panic(err)
	}

	return &input, nil
}

func (d *db) ProductTypeGetByID(id int) (*models.ProductType, error) {

	var productType models.ProductType
	err := d.Tx.Model(&models.ProductType{}).Where("id = ?", id).Take(&productType).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			d.Tx.Rollback()
			panic(err)
		} else {
			d.Tx.Rollback()
			return nil, fmt.Errorf("product type not found")
		}
	}

	return &productType, nil
}

func (d *db) ProductTypeGetAll() ([]*models.ProductType, error) {

	var productTypes []*models.ProductType
	err := d.Tx.Model(&models.ProductType{}).Find(&productTypes).Error
	if err != nil {
		d.Tx.Rollback()
		panic(err)
	}

	return productTypes, nil
}

func (d *db) ProductTypeUpdate(input models.ProductType) (*models.ProductType, error) {

	err := d.Tx.Model(&models.ProductType{}).Where("id = ?", input.ID).Updates(&input).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			d.Tx.Rollback()
			panic(err)
		} else {
			d.Tx.Rollback()
			return nil, fmt.Errorf("product type not found")
		}
	}

	return &input, nil
}

func (d *db) ProductTypeDelete(id int) (*int, error) {

	err := d.Tx.Model(&models.ProductType{}).Where("id = ?", id).Delete(&models.ProductType{}).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			d.Tx.Rollback()
			panic(err)
		} else {
			d.Tx.Rollback()
			return nil, fmt.Errorf("product type not found")
		}
	}

	return &id, nil
}
