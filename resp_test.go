package ginkit

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lantonster/cerrors"
	"github.com/stretchr/testify/assert"
)

// mockResp is a helper function to create a Resp object
func mockResp(data interface{}, err error, msg string) Resp {
	return Resp{
		Data:    data,
		Error:   err,
		Message: msg,
	}
}

// TestResponse tests the Response function
func TestResponse(t *testing.T) {
	// Setup Gin
	gin.SetMode(gin.TestMode)

	// Define the test cases
	tests := []struct {
		name       string
		resp       Resp
		wantStatus int
		wantCode   int
		wantMsg    string
		wantData   interface{}
	}{
		{
			name:       "Success",
			resp:       mockResp("data", nil, ""),
			wantStatus: http.StatusOK,
			wantCode:   http.StatusOK,
			wantMsg:    "Succcess",
			wantData:   "data",
		},
		{
			name:       "ErrorWithMessage",
			resp:       mockResp(nil, errors.New("error message"), ""),
			wantStatus: http.StatusOK,
			wantCode:   http.StatusInternalServerError,
			wantMsg:    "error message",
			wantData:   nil,
		},
		{
			name:       "ErrorWithoutMessage",
			resp:       mockResp(nil, cerrors.WithCode(1001, "with code"), ""),
			wantStatus: http.StatusOK,
			wantCode:   1001,
			wantMsg:    "with code",
			wantData:   nil,
		},
		{
			name:       "CustomMessage",
			resp:       mockResp("data", nil, "custom message"),
			wantStatus: http.StatusOK,
			wantCode:   http.StatusOK,
			wantMsg:    "custom message",
			wantData:   "data",
		},
	}

	// Run the test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a response recorder to record the response
			w := httptest.NewRecorder()

			// Create a new gin context
			c, _ := gin.CreateTestContext(w)

			// Call the Response function
			tt.resp.Response(c)

			// Assert the response status code
			assert.Equal(t, tt.wantStatus, w.Code)

			// Assert the response body
			var gResp ginResp
			err := json.Unmarshal(w.Body.Bytes(), &gResp)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantCode, gResp.Code)
			assert.Equal(t, tt.wantMsg, gResp.Message)
			assert.Equal(t, tt.wantData, gResp.Data)
		})
	}
}
