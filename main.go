package main

import (
	"engineering-task/interfaces"
	"engineering-task/usecases"
	"net/http"
)

func main() {
	dbCarRepository := interfaces.NewDbCarRepository()
	searchInteractor := usecases.SearchInteractor{
		CarRepository: dbCarRepository,
	}

	http.HandleFunc("/search", interfaces.HandleRequest(searchInteractor))
	http.ListenAndServe(":8080", nil)
}
