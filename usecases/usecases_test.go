package usecases_test

import (
	"engineering-task/domain"
	"engineering-task/infrastructure"
	"engineering-task/mocks"
	"engineering-task/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func instantiateSearchInteractor(cars []domain.Car) usecases.SearchInteractor {
	carRepository := mocks.CarRepository{Cars: cars}
	searchInteractor := usecases.SearchInteractor{
		CarRepository: carRepository,
		Logger:        infrastructure.Logger{},
	}
	return searchInteractor
}

func TestSearch_ComputesTotalCarsFound_SearchResultUpdated(t *testing.T) {
	cars := []domain.Car{
		{Make: "Tesla", Model: "Model Y", Year: 2019, Price: 50000, VehicleCount: 30},
		{Make: "Acura", Model: "IDX", Year: 2017, Price: 20000, VehicleCount: 30},
		{Make: "Honda", Model: "CRV", Year: 2018, Price: 60000, VehicleCount: 50},
		{Make: "Kia", Model: "EV6", Year: 2019, Price: 50000, VehicleCount: 30},
		{Make: "Ford", Model: "Mach E", Year: 2020, Price: 60000, VehicleCount: 10},
	}
	searchInteractor := instantiateSearchInteractor(cars)

	searchResult := searchInteractor.Search("Tesla", "Model Y", 2019, 50000)
	assert.Equal(t, 30, searchResult.TotalCount)
}

func TestSearch_CaseSensitiveNamesPassed_SearchResultUpdated(t *testing.T) {
	cars := []domain.Car{
		{Make: "Tesla", Model: "Model y", Year: 2019, Price: 50000, VehicleCount: 30},
	}
	searchInteractor := instantiateSearchInteractor(cars)

	searchResult := searchInteractor.Search("tEsLa", "MoDeL y", 2019, 50000)
	assert.Equal(t, 30, searchResult.TotalCount)
}
