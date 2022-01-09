package usecases

import (
	"engineering-task/domain"
	"fmt"
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

func (si SearchInteractor) Search(make, model string, year, price int) []SearchResult {
	cars := si.CarRepository.GetAllCars()
	si.Logger.Log(fmt.Sprintf("Scanning through %d records", len(cars)))
	return []SearchResult{}
}
