package main

import (
	"engineering-task/infrastructure"
	"engineering-task/interfaces"
	"engineering-task/usecases"
	"net/http"
)

func main() {
	carRepository := interfaces.NewCarRepository()
	searchInteractor := usecases.SearchInteractor{
		CarRepository: carRepository,
		Logger:        infrastructure.Logger{},
	}

	http.HandleFunc("/search", interfaces.HandleRequest(searchInteractor))
	http.ListenAndServe(":8080", nil)
}
