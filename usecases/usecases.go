package usecases

import (
	"engineering-task/domain"
	"fmt"
	"strings"
)

type Suggestion struct {
	Make  string
	Model string
	Year  string
	Price string
}

type SearchResult struct {
	TotalCount          int
	MakeModelMatchCount int
	LowestPrice         int
	HighestPrice        int
	MedianPrice         int
	Suggestions         []Suggestion
}

type SearchRepository interface {
	Search(make, model, year string, budget int) []SearchResult
}

type Logger interface {
	Log(args ...interface{})
}

type Car struct {
	Make  string
	Model string
	Year  string
	Price int
}

type SearchInteractor struct {
	CarRepository domain.CarRepository
	Logger        Logger
}

func (si SearchInteractor) Search(make, model string, year, price int) SearchResult {
	cars := si.CarRepository.GetAllCars()
	si.Logger.Log(fmt.Sprintf("Scanning through %d records", len(cars)))

	totalCars := 0
	for _, car := range cars {
		if isCompleteMatch(make, model, year, price, car) {
			totalCars += car.VehicleCount
		}
	}

	return SearchResult{
		TotalCount: totalCars,
	}
}

func isCompleteMatch(make, model string, year, price int, car domain.Car) bool {
	return strings.HasPrefix(strings.ToLower(car.Make), strings.ToLower(make)) &&
		strings.HasPrefix(strings.ToLower(car.Model), strings.ToLower(model)) &&
		car.Year == year && car.Price == price
}
