package usecases_test

import (
	"engineering-task/domain"
	"engineering-task/infrastructure"
	"engineering-task/mocks"
	"engineering-task/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cars = []domain.Car{
	{Make: "Tesla", Model: "Model Y", Year: 2019, Price: 50000, VehicleCount: 30},
	{Make: "Acura", Model: "IDX", Year: 2017, Price: 20000, VehicleCount: 30},
	{Make: "Honda", Model: "CRV", Year: 2018, Price: 60000, VehicleCount: 50},
	{Make: "Kia", Model: "EV6", Year: 2019, Price: 50000, VehicleCount: 30},
	{Make: "Kia", Model: "EV6", Year: 2020, Price: 60000, VehicleCount: 30},
	{Make: "Ford", Model: "Mach E", Year: 2020, Price: 60000, VehicleCount: 4},
	{Make: "Tesla", Model: "Model Y", Year: 2018, Price: 40000, VehicleCount: 60},
	{Make: "Hyundai", Model: "Kona", Year: 2018, Price: 50000, VehicleCount: 60},
	{Make: "Hyundai", Model: "Kona", Year: 2019, Price: 40000, VehicleCount: 60},
	{Make: "Hyundai", Model: "Kona", Year: 2019, Price: 70000, VehicleCount: 60},
}

func instantiateSearchInteractor(cars []domain.Car) usecases.SearchInteractor {
	carRepository := mocks.CarRepository{Cars: cars}
	searchInteractor := usecases.SearchInteractor{
		CarRepository: carRepository,
		Logger:        infrastructure.Logger{},
	}
	return searchInteractor
}

func validateResults(t *testing.T, searchResult usecases.SearchResult, expectedTotalCount, expectedMakeModelMatchCount int, expectedPricingStatistic []usecases.PricingStatistic) {
	assert.Equal(t, expectedTotalCount, searchResult.TotalCount)
	assert.Equal(t, expectedMakeModelMatchCount, searchResult.MakeModelMatchCount)
	assert.Equal(t, expectedPricingStatistic, searchResult.PricingStatistics)
}

func TestSearch_ComputesTotalCarsFound_SearchResultUpdated(t *testing.T) {
	searchInteractor := instantiateSearchInteractor(cars)

	searchResult := searchInteractor.Search("Tesla", "Model Y", 2019, 50000)
	expectedPricingStatistics := []usecases.PricingStatistic{
		{Vehicle: "TeslaModel Y", LowestPrice: 40000, HighestPrice: 50000, MedianPrice: 40000},
		{Vehicle: "KiaEV6", LowestPrice: 50000, HighestPrice: 60000, MedianPrice: 55000},
		{Vehicle: "HyundaiKona", LowestPrice: 40000, HighestPrice: 70000, MedianPrice: 50000},
	}
	validateResults(t, searchResult, 30, 90, expectedPricingStatistics)
}

func TestSearch_CaseSensitiveNamesPassed_SearchResultUpdated(t *testing.T) {
	searchInteractor := instantiateSearchInteractor(cars)

	searchResult := searchInteractor.Search("tEsLa", "MoDeL y", 2019, 70000)
	expectedPricingStatistics := []usecases.PricingStatistic{
		{Vehicle: "HyundaiKona", LowestPrice: 40000, HighestPrice: 70000, MedianPrice: 50000},
	}

	validateResults(t, searchResult, 0, 90, expectedPricingStatistics)
}
