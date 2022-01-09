package main

import (
	"engineering-task/interfaces"
	"engineering-task/usecases"
	"net/http"
)

func main() {
	carRepository := interfaces.NewCarRepository()
	searchInteractor := usecases.SearchInteractor{
		CarRepository: carRepository,
	}

	http.HandleFunc("/search", interfaces.HandleRequest(searchInteractor))
	http.ListenAndServe(":8080", nil)
}
