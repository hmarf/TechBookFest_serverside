package repository

import (
	"go/domain/model"
)

type CircleRepository interface {
	GetCircleData(string) ([]model.Circle, error)
}
