package service

import (
	"go/domain/model"
	"go/domain/repository"
)

type circleService struct {
	circleRepo repository.CircleRepository
}

type CircleInterface interface {
	GetCircle(string) ([]model.Circle, error)
	GetAllCircle() ([]model.Circle, error)
}

func NewCircleService(c repository.CircleRepository) CircleInterface {
	return &circleService{circleRepo: c}
}

func (c *circleService) GetCircle(keyword string) ([]model.Circle, error) {
	return c.circleRepo.GetCircleData(keyword)
}

func (c *circleService) GetAllCircle() ([]model.Circle, error) {
	return c.circleRepo.GetAllCircleData()
}
