package interfaces

import (
	"encoding/json"
	"engineering-task/usecases"
	"fmt"
	"net/http"
)

type searchInput struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Year   int    `json:"year"`
	Budget int    `json:"budget"`
}

// SearchResult is the final payload sent back to the user
type SearchResult struct {
	Data usecases.SearchResult `json:"data"`
}

// SearchInteractor shadows the definition of underlying usecase layer
// allowing any instance of this interface eligible to passed
// to usecases
type SearchInteractor interface {
	Search(make, model string, year, budget int) (usecases.SearchResult, error)
}

// WebserviceHandler is used to make calls to underlying usecases
type WebserviceHandler struct {
	SearchInteractor SearchInteractor
}

// HandleRequest acts as a handler for incoming requests on the /search path
func HandleRequest(SearchInteractor SearchInteractor) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var payload searchInput
		if req.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			jsonResp := constructErrorResponse("Invalid Method Called", fmt.Sprintf("Method %s not allowed", req.Method))
			w.Write(jsonResp)
			return
		}

		if req.URL.Path != "/api/search" {
			w.WriteHeader(http.StatusNotFound)
			jsonResp := constructErrorResponse("Invalid Path Called", fmt.Sprintf("Path %s not defined", req.URL.Path))
			w.Write(jsonResp)
			return
		}

		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			jsonResp := constructErrorResponse("Request body parsing failed", err.Error())
			w.Write(jsonResp)
			return
		}

		wh := WebserviceHandler{SearchInteractor: SearchInteractor}
		data, err := wh.SearchInteractor.Search(payload.Make, payload.Model, payload.Year, payload.Budget)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			jsonResp := constructErrorResponse("Internal Server Error Occured", err.Error())
			w.Write(jsonResp)
			return
		}
		json.NewEncoder(w).Encode(SearchResult{Data: data})
	}
}
