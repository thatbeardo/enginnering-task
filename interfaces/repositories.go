package interfaces

import (
	"engineering-task/domain"
)

type dbCarRepository struct {
}

func (dcr dbCarRepository) GetAllCars() []domain.Car {
	return []domain.Car{{Make: "Tesla", Model: "Y", Year: "2010", Price: 123}}
}

func NewDbCarRepository() dbCarRepository {
	return dbCarRepository{}
}
