package mocks

import (
	"engineering-task/domain"
	"engineering-task/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

// SearchInteractor provides mock implementation of the Search method.
type SearchInteractor struct {
	ExpectedMake   string
	ExpectedModel  string
	ExpectedYear   int
	ExpectedBudget int
	Result         usecases.SearchResult
	T              *testing.T
}

// Search inspects all the arguments passed to ensure data isn't lost in transit
func (si SearchInteractor) Search(make, model string, year, budget int) usecases.SearchResult {
	assert.Equal(si.T, si.ExpectedMake, make)
	assert.Equal(si.T, si.ExpectedModel, model)
	assert.Equal(si.T, si.ExpectedYear, year)
	assert.Equal(si.T, si.ExpectedBudget, budget)
	return si.Result
}

// CarRepostiory provides mocked implementation for the GetAllCar() method
type CarRepository struct {
	Cars []domain.Car
}

func (cr CarRepository) GetAllCars() []domain.Car {
	return cr.Cars
}
