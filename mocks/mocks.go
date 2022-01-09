package mocks

import (
	"engineering-task/domain"
	"engineering-task/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SearchInteractor struct {
	ExpectedMake   string
	ExpectedModel  string
	ExpectedYear   string
	ExpectedBudget int
	Results        []usecases.SearchResult
	T              *testing.T
}

func (si SearchInteractor) Search(make, model string, year, budget int) []usecases.SearchResult {
	assert.Equal(si.T, si.ExpectedMake, make)
	assert.Equal(si.T, si.ExpectedModel, model)
	assert.Equal(si.T, si.ExpectedYear, year)
	assert.Equal(si.T, si.ExpectedBudget, budget)
	return []usecases.SearchResult{}
}

type CarRepository struct {
	Cars []domain.Car
}

func (cr CarRepository) GetAllCars() []domain.Car {
	return cr.Cars
}
