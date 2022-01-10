package interfaces_test

import (
	"engineering-task/interfaces"
	"engineering-task/mocks"
	"engineering-task/usecases"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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

func inspectResponse(t *testing.T, rr *httptest.ResponseRecorder, expectedStatusCode int, expectedBody string) {
	responseBody := rr.Body.String()
	if strings.TrimSpace(responseBody) != strings.TrimSpace(expectedBody) {
		t.Errorf("handler returned wrong body: got %v expected %v",
			responseBody, expectedBody)
	}
	if responseStatusCode := rr.Code; responseStatusCode != expectedStatusCode {
		t.Errorf("handler returned wrong status code: got %v expected %v",
			responseStatusCode, http.StatusOK)
	}
}

func TestHandleRequest_InvalidMethodGet_StatusMethodNotAllowed(t *testing.T) {
	rr := performRequest(t, "GET", "/", "", mocks.SearchInteractor{})
	inspectResponse(t, rr, http.StatusMethodNotAllowed, `{"error":"Method GET not allowed","status":"Invalid Method Called"}`)
}

func TestHandleRequest_InvalidPath_StatusMethodNotAllowed(t *testing.T) {
	rr := performRequest(t, "POST", "/", "", mocks.SearchInteractor{})
	inspectResponse(t, rr, http.StatusNotFound, `{"error":"Path / not defined","status":"Invalid Path Called"}`)
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
	rr := performRequest(t, "POST", "/api/search", validPayload, msi)
	inspectResponse(t, rr, http.StatusOK, `{"data":{"totalCount":0,"makeModelMatchCount":0,"pricingStatistics":null,"suggestions":null}}`)
}

func TestHandleRequest_PostInvalidPayload_StatusBadRequest(t *testing.T) {
	rr := performRequest(t, "POST", "/api/search", invalidPayload, mocks.SearchInteractor{})

	inspectResponse(t, rr, http.StatusBadRequest, `{"error":"json: cannot unmarshal string into Go struct field searchInput.budget of type int","status":"Request body parsing failed"}`)
}

func TestHandleRequest_EmptyPayload_StatusOK(t *testing.T) {
	rr := performRequest(t, "POST", "/api/search", emptyPayload, mocks.SearchInteractor{
		T: t,
	})

	inspectResponse(t, rr, http.StatusOK, `{"data":{"totalCount":0,"makeModelMatchCount":0,"pricingStatistics":null,"suggestions":null}}`)
}

func TestHandleRequest_PopulatedPayload_ResultPresentStatusOK(t *testing.T) {

	rr := performRequest(t, "POST", "/api/search", emptyPayload, mocks.SearchInteractor{
		Result: mockResult,
		T:      t,
	})
	inspectResponse(t, rr, http.StatusOK, `{"data":{"totalCount":500,"makeModelMatchCount":600,"pricingStatistics":[{"vehicle":"TeslaModel 3","lowestPrice":40000,"medianPrice":45000,"highestPrice":50000}],"suggestions":[]}}`)
}
