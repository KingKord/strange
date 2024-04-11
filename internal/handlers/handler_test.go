package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	router := setupTest()
	ts := httptest.NewTLSServer(router)
	defer ts.Close()

	tests := []struct {
		name               string
		expectedStatusCode int
		URL                string
		method             string
	}{
		{
			name:               "assign meet",
			expectedStatusCode: http.StatusOK,
			URL:                "/api/v1/schedule/reserve",
			method:             "post",
		},
		{
			name:               "root",
			expectedStatusCode: http.StatusOK,
			URL:                "/api/v1/",
			method:             "get",
		},
		{
			name:               "day schedule",
			expectedStatusCode: http.StatusOK,
			URL:                "/api/v1/schedule/day?date=2024-04-11&user-id=1234",
			method:             "get",
		},
		{
			name:               "day schedule internal error",
			expectedStatusCode: http.StatusInternalServerError,
			URL:                "/api/v1/schedule/day?date=2024-04-11&user-id=4321",
			method:             "get",
		},
		{
			name:               "day schedule bad request",
			expectedStatusCode: http.StatusBadRequest,
			URL:                "/api/v1/schedule/day?date=1111-31-31&user-id=4321",
			method:             "get",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			var (
				response *http.Response
				err      error
			)

			if tt.method == "post" {
				response, err = ts.Client().Post(ts.URL+tt.URL, "application/json", bytes.NewBuffer([]byte(``)))
			} else {
				response, err = ts.Client().Get(ts.URL + tt.URL)
			}
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if response.StatusCode != tt.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", tt.name, tt.expectedStatusCode, response.StatusCode)
			}
		})
	}
}
