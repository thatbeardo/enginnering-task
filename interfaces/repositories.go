package interfaces

import (
	"engineering-task/domain"
)

type carRepository struct{}

func (cr carRepository) GetAllCars() ([]domain.Car, error) {
	return allCars(), nil
}

// NewCarRepository is the factory function to return concrete implementations of the SearchInteractor interface
func NewCarRepository() carRepository {
	return carRepository{}
}
