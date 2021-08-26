package controller

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	schedule2 "scheduler/internal/http/controller/schedule"
	mock_service "scheduler/internal/service/mocks"
	"scheduler/pkg/router"
	"testing"
)

func TestController_BedURL(t *testing.T) {
	testTable := []struct {
		name               string
		expectedStatusCode int
		url                string
	}{
		{
			name: "OK",
			expectedStatusCode: 404,
			url:                "/schedule-events/",
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_service.NewMockIScheduleService(c)
			controller := NewController(
				router.NewRouter(),
				schedule2.NewController(service),
				nil,
				nil,
			)
			controller.Init()

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(
				"DELETE",
				test.url,
				bytes.NewBufferString(""),
			)

			// Make request
			controller.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, test.expectedStatusCode, w.Code)
		})
	}
}
