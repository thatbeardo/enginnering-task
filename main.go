package main

import (
	"engineering-task/infrastructure"
	"engineering-task/interfaces"
	"engineering-task/usecases"
	"fmt"
	"log"
	"net/http"
)

func main() {
	carRepository := interfaces.NewCarRepository()
	searchInteractor := usecases.SearchInteractor{
		CarRepository: carRepository,
		Logger:        infrastructure.Logger{},
	}

	http.HandleFunc("/api/search", interfaces.HandleRequest(searchInteractor))
	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
