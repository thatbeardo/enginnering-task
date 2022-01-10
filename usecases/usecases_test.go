package usecases_test

import (
	"engineering-task/domain"
	"engineering-task/infrastructure"
	"engineering-task/mocks"
	"engineering-task/usecases"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cars = []domain.Car{
	{Make: "Tesla", Model: "Model Y", Year: 2019, Price: 50000, VehicleCount: 30},
	{Make: "Acura", Model: "IDX", Year: 2020, Price: 20000, VehicleCount: 30},
	{Make: "Acura", Model: "IDX", Year: 2018, Price: 21000, VehicleCount: 30},
	{Make: "Acura", Model: "IDX", Year: 2019, Price: 19000, VehicleCount: 30},
	{Make: "Acura", Model: "IDZ", Year: 2017, Price: 21000, VehicleCount: 30},
	{Make: "Volkswagen", Model: "Golf", Year: 2018, Price: 22000, VehicleCount: 100},
	{Make: "Chevrolet", Model: "Bolt", Year: 2015, Price: 19000, VehicleCount: 90},
	{Make: "Mazda", Model: "CX3", Year: 2017, Price: 21000, VehicleCount: 65},
	{Make: "Kia", Model: "Seltos", Year: 2019, Price: 22500, VehicleCount: 70},
	{Make: "Honda", Model: "CRV", Year: 2018, Price: 60000, VehicleCount: 150},
	{Make: "Honda", Model: "CRV", Year: 2018, Price: 40000, VehicleCount: 50},
	{Make: "Kia", Model: "EV6", Year: 2019, Price: 50000, VehicleCount: 30},
	{Make: "Kia", Model: "EV6", Year: 2020, Price: 60000, VehicleCount: 30},
	{Make: "Ford", Model: "Mach E", Year: 2020, Price: 60000, VehicleCount: 80},
	{Make: "Tesla", Model: "Model Y", Year: 2018, Price: 40000, VehicleCount: 60},
	{Make: "Hyundai", Model: "Kona", Year: 2018, Price: 50000, VehicleCount: 60},
	{Make: "Hyundai", Model: "Kona", Year: 2019, Price: 40000, VehicleCount: 60},
	{Make: "Hyundai", Model: "Kona", Year: 2019, Price: 70000, VehicleCount: 55},
	{Make: "Audi", Model: "Qtron", Year: 2021, Price: 49000, VehicleCount: 55},
	{Make: "Lucid", Model: "Air", Year: 2021, Price: 55000, VehicleCount: 55},
	{Make: "Polestar", Model: "2", Year: 2019, Price: 46000, VehicleCount: 55},
}

func instantiateSearchInteractor(cars []domain.Car) usecases.SearchInteractor {
	carRepository := mocks.CarRepository{Cars: cars}
	searchInteractor := usecases.SearchInteractor{
		CarRepository: carRepository,
		Logger:        infrastructure.Logger{},
	}
	return searchInteractor
}

func validateResults(t *testing.T, searchResult usecases.SearchResult,
	expectedTotalCount,
	expectedMakeModelMatchCount int,
	expectedPricingStatistic []usecases.PricingStatistic,
	expectedSuggestions []usecases.Car,
) {
	assert.Equal(t, expectedTotalCount, searchResult.TotalCount)
	assert.Equal(t, expectedMakeModelMatchCount, searchResult.MakeModelMatchCount)
	assert.Equal(t, expectedPricingStatistic, searchResult.PricingStatistics)
	assert.Equal(t, expectedSuggestions, searchResult.Suggestions)
}

func TestSearch_ComputesTotalCarsFound_SearchResultUpdated(t *testing.T) {
	searchInteractor := instantiateSearchInteractor(cars)

	searchResult, _ := searchInteractor.Search("TeSlA", "ModEl Y", 2019, 50000)
	expectedPricingStatistics := []usecases.PricingStatistic{
		{Vehicle: "TeslaModel Y", LowestPrice: 40000, HighestPrice: 50000, MedianPrice: 40000},
		{Vehicle: "KiaEV6", LowestPrice: 50000, HighestPrice: 60000, MedianPrice: 55000},
		{Vehicle: "HyundaiKona", LowestPrice: 40000, HighestPrice: 70000, MedianPrice: 50000},
	}

	expectedSuggestions := []usecases.Car{
		{Make: "Polestar", Model: "2", Year: 2019, Price: 46000},
		{Make: "Audi", Model: "Qtron", Year: 2021, Price: 49000},
		{Make: "Tesla", Model: "Model Y", Year: 2019, Price: 50000},
		{Make: "Kia", Model: "EV6", Year: 2019, Price: 50000},
		{Make: "Hyundai", Model: "Kona", Year: 2018, Price: 50000},
	}

	validateResults(t, searchResult, 30, 90, expectedPricingStatistics, expectedSuggestions)
}

func TestSearch_VehicleSuggestions_SearchResultUpdated(t *testing.T) {
	searchInteractor := instantiateSearchInteractor(cars)

	searchResult, _ := searchInteractor.Search("Acura", "IDX", 2020, 20000)
	expectedPricingStatistics := []usecases.PricingStatistic{
		{Vehicle: "AcuraIDX", LowestPrice: 19000, HighestPrice: 21000, MedianPrice: 20000},
	}
	expectedSuggestions := []usecases.Car{
		{Make: "Chevrolet", Model: "Bolt", Year: 2015, Price: 19000},
		{Make: "Acura", Model: "IDX", Year: 2020, Price: 20000},
		{Make: "Mazda", Model: "CX3", Year: 2017, Price: 21000},
		{Make: "Volkswagen", Model: "Golf", Year: 2018, Price: 22000},
	}

	validateResults(t, searchResult, 30, 90, expectedPricingStatistics, expectedSuggestions)
}

func TestSearch_ErrorOccured_ErrorReported(t *testing.T) {
	carRepository := mocks.CarRepository{Cars: cars, Err: errors.New("test-error")}
	searchInteractor := usecases.SearchInteractor{
		CarRepository: carRepository,
		Logger:        infrastructure.Logger{},
	}

	_, err := searchInteractor.Search("", "", 0, 0)
	assert.Error(t, err, errors.New("test-error"))
}
