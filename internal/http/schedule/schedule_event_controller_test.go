package schedule

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"scheduler/internal/http"
	"scheduler/internal/model"
	mock_service "scheduler/internal/service/mocks"
	"scheduler/pkg/router"
	"testing"
)

func TestController_Show(t *testing.T) {
	type mockBehavior func(s *mock_service.MockIScheduleService, event model.ScheduleEvent, ID int)

	testTable := []struct {
		name               string
		scheduleEvent      model.ScheduleEvent
		mockBehavior       mockBehavior
		expectedStatusCode int
		url                string
	}{
		{
			name: "OK",
			scheduleEvent: model.ScheduleEvent{
				ID:      1,
				Name:    "Test schedule event",
				Time:    360,
				StartAt: 1660694400,
			},
			mockBehavior: func(s *mock_service.MockIScheduleService, event model.ScheduleEvent, ID int) {
				s.EXPECT().
					Show(event.ID).
					Return(event, nil)
			},
			expectedStatusCode: 200,
			url:                "/schedule-events/1",
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_service.NewMockIScheduleService(c)
			test.mockBehavior(service, test.scheduleEvent, test.scheduleEvent.ID)
			controller := http.NewController(
				router.NewRouter(),
				NewController(service),
			)
			controller.Init()

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(
				"GET",
				test.url,
				bytes.NewBufferString(""),
			)

			// Make request
			controller.ServeHTTP(w, req)

			// Assert
			m := &model.ScheduleEventJson{}
			json.Unmarshal(w.Body.Bytes(), m)
			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.scheduleEvent.ID, m.ID)
		})
	}
}

func TestController_List(t *testing.T) {
	type mockBehavior func(s *mock_service.MockIScheduleService, scheduleEvents []model.ScheduleEvent, params map[string]string)

	testTable := []struct {
		name                      string
		scheduleEvents            []model.ScheduleEvent
		mockBehavior              mockBehavior
		expectedStatusCode        int
		expectScheduleEventsCount int
		url                       string
		params                    map[string]string
	}{
		{
			name: "OK",
			scheduleEvents: []model.ScheduleEvent{
				{
					ID:        1,
					Name:      "First event",
					Time:      120,
					StartAt:   1660694400,
					CreatedAt: 1660594400,
					UpdatedAt: 1660594400,
				},
				{
					ID:        2,
					Name:      "Second event",
					Time:      360,
					StartAt:   1660794400,
					CreatedAt: 1660494400,
					UpdatedAt: 1660494400,
				},
			},
			mockBehavior: func(s *mock_service.MockIScheduleService, scheduleEvents []model.ScheduleEvent, params map[string]string) {
				s.EXPECT().
					List(params).
					Return(scheduleEvents, nil)
			},
			expectedStatusCode:        200,
			expectScheduleEventsCount: 2,
			url:                       "/schedule-events",
			params:                    make(map[string]string),
		},
	}

	for _, test := range testTable {
		// Init Dependencies
		c := gomock.NewController(t)
		defer c.Finish()

		service := mock_service.NewMockIScheduleService(c)
		test.mockBehavior(service, test.scheduleEvents, test.params)
		controller := http.NewController(
			router.NewRouter(),
			NewController(service),
		)
		controller.Init()

		// Create Request
		w := httptest.NewRecorder()
		req := httptest.NewRequest(
			"GET",
			test.url,
			bytes.NewBufferString(""),
		)

		// Make request
		controller.ServeHTTP(w, req)

		// Assert
		m := make([]model.ScheduleEventJson, 2)
		json.Unmarshal(w.Body.Bytes(), &m)
		assert.Equal(t, test.expectedStatusCode, w.Code)
		assert.Equal(t, test.expectScheduleEventsCount, len(m))
	}
}

func TestController_Create(t *testing.T) {
	type mockBehavior func(s *mock_service.MockIScheduleService, scheduleEvent model.ScheduleEvent)

	testTable := []struct {
		name               string
		inputBody          string
		scheduleEvent      model.ScheduleEvent
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:      "OK",
			inputBody: `{"name": "Test schedule event","time": 360,"start_at": "08/17/22 00:00:00"}`,
			scheduleEvent: model.ScheduleEvent{
				Name:    "Test schedule event",
				Time:    360,
				StartAt: 1660694400,
			},
			mockBehavior: func(s *mock_service.MockIScheduleService, scheduleEvent model.ScheduleEvent) {
				s.EXPECT().
					Create(scheduleEvent).
					Return(scheduleEvent, nil)
			},
			expectedStatusCode: 201,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_service.NewMockIScheduleService(c)
			test.mockBehavior(service, test.scheduleEvent)
			controller := http.NewController(
				router.NewRouter(),
				NewController(service),
			)
			controller.Init()

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(
				"POST",
				"/schedule-events/",
				bytes.NewBufferString(test.inputBody),
			)

			// Make request
			controller.ServeHTTP(w, req)

			// Assert
			m := &model.ScheduleEventJson{}
			json.Unmarshal(w.Body.Bytes(), m)
			scheduleEventJson := test.scheduleEvent.ToScheduleEventJson()
			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, scheduleEventJson.Time, m.Time)
			assert.Equal(t, scheduleEventJson.StartAt, m.StartAt)
		})
	}
}

func TestController_Update(t *testing.T) {
	type mockBehavior func(
		s *mock_service.MockIScheduleService,
		ID int,
		inputScheduleEvent model.ScheduleEvent,
		outputScheduleEvent model.ScheduleEvent,
	)

	testTable := []struct {
		name                string
		inputBody           string
		inputScheduleEvent  model.ScheduleEvent
		outputScheduleEvent model.ScheduleEvent
		mockBehavior        mockBehavior
		expectedStatusCode  int
		url                 string
	}{
		{
			name:      "OK",
			inputBody: `{"name": "Test schedule event","time": 360,"start_at": "08/17/22 00:00:00"}`,
			inputScheduleEvent: model.ScheduleEvent{
				Name:    "Test schedule event",
				Time:    360,
				StartAt: 1660694400,
			},
			outputScheduleEvent: model.ScheduleEvent{
				ID:      1,
				Name:    "Test schedule event",
				Time:    360,
				StartAt: 1660694400,
			},
			mockBehavior: func(s *mock_service.MockIScheduleService, ID int, iScheduleEvent model.ScheduleEvent, oScheduleEvent model.ScheduleEvent) {
				s.EXPECT().
					Update(ID, iScheduleEvent).
					Return(oScheduleEvent, nil)
			},
			expectedStatusCode: 200,
			url:                "/schedule-events/1",
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_service.NewMockIScheduleService(c)
			test.mockBehavior(service, test.outputScheduleEvent.ID, test.inputScheduleEvent, test.outputScheduleEvent)
			controller := http.NewController(
				router.NewRouter(),
				NewController(service),
			)
			controller.Init()

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(
				"PUT",
				test.url,
				bytes.NewBufferString(test.inputBody),
			)

			// Make request
			controller.ServeHTTP(w, req)

			// Assert
			m := &model.ScheduleEventJson{}
			json.Unmarshal(w.Body.Bytes(), m)
			scheduleEventJson := test.outputScheduleEvent.ToScheduleEventJson()
			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, scheduleEventJson.Time, m.Time)
			assert.Equal(t, scheduleEventJson.StartAt, m.StartAt)
		})
	}
}

func TestController_Delete(t *testing.T) {
	type mockBehavior func(s *mock_service.MockIScheduleService, ID int)

	testTable := []struct {
		name               string
		mockBehavior       mockBehavior
		expectedStatusCode int
		url                string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mock_service.MockIScheduleService, ID int) {
				s.EXPECT().
					Delete(ID).
					Return(nil)
			},
			expectedStatusCode: 204,
			url:                "/schedule-events/1",
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_service.NewMockIScheduleService(c)
			test.mockBehavior(service, 1)
			controller := http.NewController(
				router.NewRouter(),
				NewController(service),
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
