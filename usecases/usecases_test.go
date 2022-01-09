package usecases_test

import (
	"engineering-task/domain"
	"engineering-task/infrastructure"
	"engineering-task/mocks"
	"engineering-task/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch_ComputesTotalCarsFound_SearchResultUpdated(t *testing.T) {
	cars := []domain.Car{
		{Make: "Tesla", Model: "Model Y", Year: 2019, Price: 50000, VehicleCount: 30},
	}
	carRepository := mocks.CarRepository{Cars: cars}
	searchInteractor := usecases.SearchInteractor{
		CarRepository: carRepository,
		Logger:        infrastructure.Logger{},
	}

	searchResult := searchInteractor.Search("Tesla", "Model Y", 2019, 50000)
	assert.Equal(t, 1, len(searchResult))
}
