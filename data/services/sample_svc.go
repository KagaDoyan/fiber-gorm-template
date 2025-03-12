package services

import (
	"go-fiber/data/repositories"
	"go-fiber/domain/models"
)

type SampleService interface {
	FindAll() ([]models.Sample, error)
}

type sampleService struct {
	samplerepo repositories.SampleRepository
}

func (s sampleService) FindAll() ([]models.Sample, error) {
	var result []models.Sample
	samples, err := s.samplerepo.FindAll()
	if err != nil {
		return nil, err
	}
	for _, sample := range samples {
		result = append(result, models.Sample{
			ID:    sample.ID,
			Title: sample.Title,
			Body:  sample.Body,
		})
	}
	return result, nil
}

func NewSampleServices(samplerepo repositories.SampleRepository) SampleService {
	return &sampleService{samplerepo: samplerepo}
}
