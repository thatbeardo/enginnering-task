package interfaces

import (
	"engineering-task/domain"
)

type carRepository struct{}

// GetAllCars returns all vehicles present in data source
func (cr carRepository) GetAllCars() ([]domain.Car, error) {
	return allCars(), nil
}

// NewCarRepository is the factory function to return concrete implementations of the SearchInteractor interface
func NewCarRepository() carRepository {
	return carRepository{}
}
