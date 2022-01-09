package usecases

import (
	"engineering-task/domain"
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

func (si SearchInteractor) Search(make, model, year string, price int) []SearchResult {
	si.CarRepository.GetAllCars()
	return []SearchResult{}
}
