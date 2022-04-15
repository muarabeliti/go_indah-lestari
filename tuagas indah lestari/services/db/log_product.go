package services

import (
	"myapp/models"
)

func (d *db) LogProductCreate(input models.LogProduct) (*models.LogProduct, error) {
	err := d.Tx.Model(&models.LogProduct{}).Create(&input).Error
	if err != nil {
		d.Tx.Rollback()
		panic(err)
	}

	return &input, nil
}
