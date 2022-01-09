package interfaces

import (
	"engineering-task/domain"
)

type carRepository struct{}

func (cr carRepository) GetAllCars() []domain.Car {
	return allCars()
}

func NewCarRepository() carRepository {
	return carRepository{}
}
