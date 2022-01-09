package interfaces_test

import (
	"engineering-task/interfaces"
	"engineering-task/usecases"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const validPayload = `{"make":"Tesla", "model":"Model Y", "year": "2019", "budget":50000}`
const invalidPayload = `{"make":"Tesla", "model":"Model Y", "year": "2019", "budget":"50000"}`
const emptyPayload = `{"make":"", "model":"", "year": "", "budget":0}`

type MockSearchInteractor struct {
	expectedMake   string
	expectedModel  string
	expectedYear   string
	expectedBudget int
	results        []usecases.SearchResult
	t              *testing.T
}

func (msi MockSearchInteractor) Search(make, model, year string, budget int) []usecases.SearchResult {
	assert.Equal(msi.t, msi.expectedMake, make)
	assert.Equal(msi.t, msi.expectedModel, model)
	assert.Equal(msi.t, msi.expectedYear, year)
	assert.Equal(msi.t, msi.expectedBudget, budget)
	return []usecases.SearchResult{}
}

func performRequest(t *testing.T, method, path, body string, msi MockSearchInteractor) *httptest.ResponseRecorder {
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
	rr := performRequest(t, "GET", "/", "", MockSearchInteractor{})
	inspectResponse(t, rr, http.StatusMethodNotAllowed)
}

func TestHandleRequest_PostValidPayload_StatusOK(t *testing.T) {
	msi := MockSearchInteractor{
		expectedMake:   "Tesla",
		expectedModel:  "Model Y",
		expectedYear:   "2019",
		expectedBudget: 50000,
		results:        []usecases.SearchResult{},
		t:              t,
	}
	rr := performRequest(t, "POST", "/", validPayload, msi)
	inspectResponse(t, rr, http.StatusOK)
}

func TestHandleRequest_PostInvalidPayload_StatusBadRequest(t *testing.T) {
	rr := performRequest(t, "POST", "/", invalidPayload, MockSearchInteractor{})

	inspectResponse(t, rr, http.StatusBadRequest)
}

func TestHandleRequest_EmptyPayload_StatusOK(t *testing.T) {
	rr := performRequest(t, "POST", "/", emptyPayload, MockSearchInteractor{
		t: t,
	})

	inspectResponse(t, rr, http.StatusOK)
}

// Check the response body is what we expect.
// expected := `{"alive": true}`
// if rr.Body.String() != expected {
// 	t.Errorf("handler returned unexpected body: got %v want %v",
// 		rr.Body.String(), expected)
// }
