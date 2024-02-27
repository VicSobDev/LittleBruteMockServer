package server_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vicsobdev/LittleBruteMockServer/server" // Update with the actual path to your server package

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Setup
	router := gin.Default()
	api := server.Api{} // Assuming you have a constructor or setup for your Api struct
	router.POST("/login", api.Login)

	// Define your test cases
	tests := []struct {
		description   string
		requestBody   string
		expectedCode  int
		expectedError string
	}{
		{
			description:   "Valid request",
			requestBody:   `{"username":"testuser","password":"testpass"}`,
			expectedCode:  http.StatusOK,
			expectedError: "",
		},
		{
			description:   "Invalid request body",
			requestBody:   `{{invalid json}}`,
			expectedCode:  http.StatusBadRequest,
			expectedError: "Invalid Request Body",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			// Create a request to pass to our handler.
			req, err := http.NewRequest("POST", "/login", bytes.NewBufferString(test.requestBody))
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			// Check the status code is what we expect.
			assert.Equal(t, test.expectedCode, rr.Code, test.description)

			// Check the response body is what we expect.
			if test.expectedError != "" {
				assert.Contains(t, rr.Body.String(), test.expectedError, test.description)
			} else {
			}
		})
	}
}
