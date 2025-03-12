package repositories

import (
	"go-fiber/domain/entities"

	"gorm.io/gorm"
)

type sampleRepository struct {
	db *gorm.DB
}
type SampleRepository interface {
	FindAll() ([]entities.Sample, error)
}

func (s sampleRepository) FindAll() ([]entities.Sample, error) {
	var samples []entities.Sample
	err := s.db.Find(&samples).Error
	if err != nil {
		return nil, err
	}
	return samples, nil
}

func NewSampleRepository(db *gorm.DB) SampleRepository {
	return &sampleRepository{db: db}
}
