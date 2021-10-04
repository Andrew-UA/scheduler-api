package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"scheduler/internal/helpers"
	"scheduler/internal/model"
	mock_service "scheduler/internal/service/mocks"
	"scheduler/pkg/logger"
	"scheduler/pkg/router"
	"testing"
	"time"
)

func TestController_Show(t *testing.T) {
	type mockBehavior func(s *mock_service.MockIScheduleService, ctx context.Context, event model.ScheduleEvent)
	l := logger.LoggerMock{}

	testTable := []struct {
		name               string
		scheduleEvent      model.ScheduleEvent
		user               model.User
		mockBehavior       mockBehavior
		expectedStatusCode int
		url                string
	}{
		{
			name: "OK",
			scheduleEvent: model.ScheduleEvent{
				ID:      1,
				UserID:  1,
				Name:    "Test schedule event",
				Time:    360,
				StartAt: 1660694400,
			},
			user: model.User{
				ID:       1,
				Timezone: "UTC",
			},
			mockBehavior: func(s *mock_service.MockIScheduleService, ctx context.Context, event model.ScheduleEvent) {
				s.EXPECT().
					Show(ctx, event.ID).
					Return(event, nil)
			},
			expectedStatusCode: 200,
			url:                "/schedule-events/1",
		},
		{
			name: "NOT FOUND",
			scheduleEvent: model.ScheduleEvent{
				ID:      1,
				UserID:  1,
				Name:    "Test schedule event",
				Time:    360,
				StartAt: 1660694400,
			},
			mockBehavior: func(s *mock_service.MockIScheduleService, ctx context.Context, event model.ScheduleEvent) {
				s.EXPECT().
					Show(gomock.Any(), 2).
					Return(event, errors.New("NOT FOUND"))
			},
			expectedStatusCode: 404,
			url:                "/schedule-events/2",
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_service.NewMockIScheduleService(c)
			ctx := helpers.SetUserToContext(test.user, context.Background())
			test.mockBehavior(service, ctx, test.scheduleEvent)
			controller := NewController(
				router.NewRouter(),
				nil,
				NewScheduleController(service, nil, l),
				nil,
				nil,
				nil,
			)
			controller.Init()

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(
				"GET",
				test.url,
				bytes.NewBufferString(""),
			)
			req = req.WithContext(ctx)

			// Make request
			controller.ServeHTTP(w, req)

			// Assert
			m := &model.ScheduleEventJson{}
			json.Unmarshal(w.Body.Bytes(), m)
			assert.Equal(t, test.expectedStatusCode, w.Code)
			if test.name == "OK" {
				assert.Equal(t, test.scheduleEvent.ID, m.ID)
			}
		})
	}
}

func TestController_List(t *testing.T) {
	type mockBehavior func(s *mock_service.MockIScheduleService, ctx context.Context, scheduleEvents []model.ScheduleEvent, params map[string]string)
	l := logger.LoggerMock{}

	testTable := []struct {
		name                      string
		scheduleEvents            []model.ScheduleEvent
		user                      model.User
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
			user: model.User{
				ID:       1,
				Timezone: "UTC",
			},
			mockBehavior: func(s *mock_service.MockIScheduleService, ctx context.Context, scheduleEvents []model.ScheduleEvent, params map[string]string) {
				s.EXPECT().
					List(ctx, params).
					Return(scheduleEvents, nil)
			},
			expectedStatusCode:        200,
			expectScheduleEventsCount: 2,
			url:                       "/schedule-events",
			params:                    make(map[string]string),
		},
		{
			name: "OK DAY",
			scheduleEvents: []model.ScheduleEvent{
				{
					ID:        1,
					Name:      "First event",
					Time:      120,
					StartAt:   time.Now().Add(time.Hour * 24).Unix(),
					CreatedAt: 1660594400,
					UpdatedAt: 1660594400,
				},
			},
			mockBehavior: func(s *mock_service.MockIScheduleService, ctx context.Context, scheduleEvents []model.ScheduleEvent, params map[string]string) {
				s.EXPECT().
					List(ctx, params).
					Return(scheduleEvents, nil)
			},
			expectedStatusCode:        200,
			expectScheduleEventsCount: 1,
			url:                       "/schedule-events?interval=day",
			params: map[string]string{
				"interval": "day",
			},
		},
	}

	for _, test := range testTable {
		// Init Dependencies
		c := gomock.NewController(t)
		defer c.Finish()

		service := mock_service.NewMockIScheduleService(c)
		ctx := helpers.SetUserToContext(test.user, context.Background())
		test.mockBehavior(service, ctx, test.scheduleEvents, test.params)
		controller := NewController(
			router.NewRouter(),
			nil,
			NewScheduleController(service, nil, l),
			nil,
			nil,
			nil,
		)
		controller.Init()

		// Create Request
		w := httptest.NewRecorder()
		req := httptest.NewRequest(
			"GET",
			test.url,
			bytes.NewBufferString(""),
		)
		req = req.WithContext(ctx)

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
	type mockBehavior func(s *mock_service.MockIScheduleService, ctx context.Context, scheduleEvent model.ScheduleEvent)
	l := logger.LoggerMock{}

	testTable := []struct {
		name               string
		inputBody          string
		scheduleEvent      model.ScheduleEvent
		user               model.User
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
			user: model.User{
				ID:       1,
				Timezone: "UTC",
			},
			mockBehavior: func(s *mock_service.MockIScheduleService, ctx context.Context, scheduleEvent model.ScheduleEvent) {
				s.EXPECT().
					Create(ctx, scheduleEvent).
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
			ctx := helpers.SetUserToContext(test.user, context.Background())
			test.mockBehavior(service, ctx, test.scheduleEvent)
			controller := NewController(
				router.NewRouter(),
				nil,
				NewScheduleController(service, nil, l),
				nil,
				nil,
				nil,
			)
			controller.Init()

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(
				"POST",
				"/schedule-events/",
				bytes.NewBufferString(test.inputBody),
			)
			req = req.WithContext(ctx)

			// Make request
			controller.ServeHTTP(w, req)

			// Assert
			m := &model.ScheduleEventJson{}
			json.Unmarshal(w.Body.Bytes(), m)
			scheduleEventJson, err := test.scheduleEvent.ToScheduleEventJson("UTC")
			assert.Nil(t, err)
			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, scheduleEventJson.Time, m.Time)
			assert.Equal(t, scheduleEventJson.StartAt, m.StartAt)
		})
	}
}

func TestController_Update(t *testing.T) {
	type mockBehavior func(
		s *mock_service.MockIScheduleService,
		ctx context.Context,
		ID int,
		inputScheduleEvent model.ScheduleEvent,
		outputScheduleEvent model.ScheduleEvent,
	)
	l := logger.LoggerMock{}

	testTable := []struct {
		name                string
		inputBody           string
		inputScheduleEvent  model.ScheduleEvent
		outputScheduleEvent model.ScheduleEvent
		user                model.User
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
			user: model.User{
				ID:       1,
				Timezone: "UTC",
			},
			mockBehavior: func(s *mock_service.MockIScheduleService, ctx context.Context, ID int, iScheduleEvent model.ScheduleEvent, oScheduleEvent model.ScheduleEvent) {
				s.EXPECT().
					Update(ctx, ID, iScheduleEvent).
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
			ctx := helpers.SetUserToContext(test.user, context.Background())
			test.mockBehavior(service, ctx, test.outputScheduleEvent.ID, test.inputScheduleEvent, test.outputScheduleEvent)
			controller := NewController(
				router.NewRouter(),
				nil,
				NewScheduleController(service, nil, l),
				nil,
				nil,
				nil,
			)
			controller.Init()

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(
				"PUT",
				test.url,
				bytes.NewBufferString(test.inputBody),
			)
			req = req.WithContext(ctx)

			// Make request
			controller.ServeHTTP(w, req)

			// Assert
			m := &model.ScheduleEventJson{}
			json.Unmarshal(w.Body.Bytes(), m)
			scheduleEventJson, err := test.outputScheduleEvent.ToScheduleEventJson("UTC")
			assert.Nil(t, err)
			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, scheduleEventJson.Time, m.Time)
			assert.Equal(t, scheduleEventJson.StartAt, m.StartAt)
		})
	}
}

func TestController_Delete(t *testing.T) {
	type mockBehavior func(s *mock_service.MockIScheduleService, ctx context.Context, ID int)
	l := logger.LoggerMock{}

	testTable := []struct {
		name               string
		mockBehavior       mockBehavior
		user               model.User
		expectedStatusCode int
		url                string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mock_service.MockIScheduleService, ctx context.Context, ID int) {
				s.EXPECT().
					Delete(ctx, ID).
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
			ctx := helpers.SetUserToContext(test.user, context.Background())
			test.mockBehavior(service, ctx, 1)
			controller := NewController(
				router.NewRouter(),
				nil,
				NewScheduleController(service, nil, l),
				nil,
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
			req = req.WithContext(ctx)

			// Make request
			controller.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, test.expectedStatusCode, w.Code)
		})
	}
}
