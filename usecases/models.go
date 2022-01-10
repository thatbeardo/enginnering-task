package usecases

import "engineering-task/domain"

type pricingData struct {
	vehicleCount int
	price        int
}

// PricingStatistic holds lowest, highest and median price for a given vehicle
type PricingStatistic struct {
	Vehicle      string
	LowestPrice  int
	MedianPrice  int
	HighestPrice int
}

// SearchResult encapsulates all computed data requested by the user
type SearchResult struct {
	TotalCount          int
	MakeModelMatchCount int
	PricingStatistics   []PricingStatistic
	Suggestions         []Car
}

// SearchRepository provides methods to scan underlying persistence layer for data
type SearchRepository interface {
	Search(make, model, year string, budget int) []SearchResult
}

// Logger interface is to be injected during initialization of the SearchIterator struct
type Logger interface {
	Log(args ...interface{})
}

// Car encapsulates detailes pertaining to a vehicle
type Car struct {
	Make  string
	Model string
	Year  int
	Price int
}

// SearchInteractor provides concrete implementation of SearchRepository interface
type SearchInteractor struct {
	CarRepository domain.CarRepository
	Logger        Logger
}
