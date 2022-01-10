package main

import (
	"engineering-task/infrastructure"
	"engineering-task/interfaces"
	"engineering-task/usecases"
	"fmt"
	"net/http"
)

func main() {
	carRepository := interfaces.NewCarRepository()
	searchInteractor := usecases.SearchInteractor{
		CarRepository: carRepository,
		Logger:        infrastructure.Logger{},
	}

	http.HandleFunc("/api/search", interfaces.HandleRequest(searchInteractor))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server up and running on port 8080")
	}
}
