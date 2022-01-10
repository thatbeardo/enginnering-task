package usecases

import "engineering-task/domain"

type Suggestion struct {
	Make  string
	Model string
	Year  int
	Price int
}

type pricingData struct {
	vehicleCount int
	price        int
}

type PricingStatistic struct {
	Vehicle      string
	LowestPrice  int
	MedianPrice  int
	HighestPrice int
}

type SearchResult struct {
	TotalCount          int
	MakeModelMatchCount int
	PricingStatistics   []PricingStatistic
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
