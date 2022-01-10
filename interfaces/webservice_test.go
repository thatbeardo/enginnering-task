package interfaces_test

import (
	"encoding/json"
	"engineering-task/interfaces"
	"engineering-task/mocks"
	"engineering-task/usecases"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const validPayload = `{"make":"Tesla", "model":"Model Y", "year": 2019, "budget":50000}`
const invalidPayload = `{"make":"Tesla", "model":"Model Y", "year": 2019, "budget":"50000"}`
const emptyPayload = `{"make":"", "model":"", "year": 0, "budget":0}`

var mockResult = usecases.SearchResult{
	TotalCount:          500,
	MakeModelMatchCount: 600,
	PricingStatistics: []usecases.PricingStatistic{
		{Vehicle: "TeslaModel 3", LowestPrice: 40000, HighestPrice: 50000, MedianPrice: 45000},
	},
	Suggestions: []usecases.Car{},
}

func performRequest(t *testing.T, method, path, body string, msi mocks.SearchInteractor) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(interfaces.HandleRequest(msi))
	handler.ServeHTTP(rr, req)

	return rr
}

func inspectResponse(t *testing.T, rr *httptest.ResponseRecorder, expectedStatusCode int) {
	if status := rr.Code; status != expectedStatusCode {
		t.Errorf("handler returned wrong status code: got %v expected %v",
			status, http.StatusOK)
	}
}

func TestHandleRequest_Get_StatusMethodNotAllowed(t *testing.T) {
	rr := performRequest(t, "GET", "/", "", mocks.SearchInteractor{})
	inspectResponse(t, rr, http.StatusMethodNotAllowed)
}

func TestHandleRequest_PostValidPayload_StatusOK(t *testing.T) {
	msi := mocks.SearchInteractor{
		ExpectedMake:   "Tesla",
		ExpectedModel:  "Model Y",
		ExpectedYear:   2019,
		ExpectedBudget: 50000,
		Result:         usecases.SearchResult{},
		T:              t,
	}
	rr := performRequest(t, "POST", "/", validPayload, msi)
	inspectResponse(t, rr, http.StatusOK)
}

func TestHandleRequest_PostInvalidPayload_StatusBadRequest(t *testing.T) {
	rr := performRequest(t, "POST", "/", invalidPayload, mocks.SearchInteractor{})

	inspectResponse(t, rr, http.StatusBadRequest)
}

func TestHandleRequest_EmptyPayload_StatusOK(t *testing.T) {
	rr := performRequest(t, "POST", "/", emptyPayload, mocks.SearchInteractor{
		T: t,
	})

	inspectResponse(t, rr, http.StatusOK)
}

func TestHandleRequest_PopulatedPayload_ResultPresentStatusOK(t *testing.T) {

	rr := performRequest(t, "POST", "/", emptyPayload, mocks.SearchInteractor{
		Result: mockResult,
		T:      t,
	})
	var response interfaces.SearchResult
	json.Unmarshal([]byte(rr.Body.Bytes()), &response)

	assert.Equal(t, response.Data, mockResult)
	inspectResponse(t, rr, http.StatusOK)
}
