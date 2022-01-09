package interfaces

import (
	"encoding/json"
	"engineering-task/usecases"
	"net/http"
)

type searchInput struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Year   string `json:"year"`
	Budget int    `json:"budget"`
}

type SearchInteractor interface {
	Search(make, model, year string, budget int) []usecases.SearchResult
}

type WebserviceHandler struct {
	SearchInteractor SearchInteractor
}

func HandleRequest(SearchInteractor SearchInteractor) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		var payload searchInput
		if req.Method == "POST" {
			err := json.NewDecoder(req.Body).Decode(&payload)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			wh := WebserviceHandler{SearchInteractor: SearchInteractor}
			searchResults := wh.SearchInteractor.Search(payload.Make, payload.Model, payload.Year, payload.Budget)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(searchResults)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}
}
