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
	SearchErr      error
	T              *testing.T
}

// Search inspects all the arguments passed to ensure data isn't lost in transit
func (si SearchInteractor) Search(make, model string, year, budget int) (usecases.SearchResult, error) {
	assert.Equal(si.T, si.ExpectedMake, make)
	assert.Equal(si.T, si.ExpectedModel, model)
	assert.Equal(si.T, si.ExpectedYear, year)
	assert.Equal(si.T, si.ExpectedBudget, budget)
	return si.Result, si.SearchErr
}

// CarRepostiory provides mocked implementation for the GetAllCar() method
type CarRepository struct {
	Cars []domain.Car
	Err  error
}

func (cr CarRepository) GetAllCars() ([]domain.Car, error) {
	return cr.Cars, cr.Err
}
