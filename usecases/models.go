package usecases

import "engineering-task/domain"

type pricingData struct {
	vehicleCount int
	price        int
}

// PricingStatistic holds lowest, highest and median price for a given vehicle
type PricingStatistic struct {
	Vehicle      string `json:"vehicle"`
	LowestPrice  int    `json:"lowestPrice"`
	MedianPrice  int    `json:"medianPrice"`
	HighestPrice int    `json:"highestPrice"`
}

// SearchResult encapsulates all computed data requested by the user
type SearchResult struct {
	TotalCount          int                `json:"totalCount"`
	MakeModelMatchCount int                `json:"makeModelMatchCount"`
	PricingStatistics   []PricingStatistic `json:"pricingStatistics"`
	Suggestions         []Car              `json:"suggestions"`
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
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
	Price int    `json:"price"`
}

// SearchInteractor provides concrete implementation of SearchRepository interface
type SearchInteractor struct {
	CarRepository domain.CarRepository
	Logger        Logger
}
