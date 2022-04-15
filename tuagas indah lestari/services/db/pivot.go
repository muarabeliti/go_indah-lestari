package services

import (
	"myapp/models"
)

func (d *db) PivotCreate(input models.Pivot) (*models.Pivot, error) {

	err := d.Tx.Model(&models.Pivot{}).Create(&input).Error
	if err != nil {
		d.Tx.Rollback()
		panic(err)
	}

	return &input, nil
}
