package usecases

import "engineering-task/domain"

type SearchRepository interface {
	Search(make, model, year string, budget int)
}

type Logger interface {
	Log(args ...interface{})
}

type Car struct {
	Make   string
	Model  string
	Year   string
	Budget string
}

type SearchInteractor struct {
	CarRepository domain.CarRepository
	Logger        Logger
}
