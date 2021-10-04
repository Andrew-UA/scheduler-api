package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"scheduler/internal/helpers"
	"scheduler/internal/model"
	mock_repository "scheduler/internal/repository/mocks"
	"scheduler/pkg/logger"
	"testing"
)

func TestService_Show(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockISchedule, scheduleEvent model.ScheduleEvent, ID int)
	l := logger.LoggerMock{}

	testTable := []struct {
		name          string
		scheduleEvent model.ScheduleEvent
		user          model.User
		mockBehavior  mockBehavior
	}{
		{
			name: "OK",
			scheduleEvent: model.ScheduleEvent{
				1,
				1,
				"Test Schedule event",
				360,
				1660794400,
				1660494400,
				1660494400,
			},
			user: model.User{
				ID:       1,
				Timezone: "UTC",
			},
			mockBehavior: func(s *mock_repository.MockISchedule, scheduleEvent model.ScheduleEvent, ID int) {
				s.EXPECT().Show(ID).Return(scheduleEvent, nil)
			},
		},
	}

	for _, test := range testTable {
		// Init Dependencies
		c := gomock.NewController(t)
		defer c.Finish()

		repo := mock_repository.NewMockISchedule(c)
		ctx := helpers.SetUserToContext(test.user, context.Background())
		test.mockBehavior(repo, test.scheduleEvent, test.scheduleEvent.ID)
		service := NewScheduleService(repo, l)

		// Test
		result, err := service.Show(ctx, test.scheduleEvent.ID)

		// Assert
		assert.Equal(t, test.scheduleEvent, result)
		assert.Nil(t, err)
	}
}

func TestService_List(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockISchedule, scheduleEvents []model.ScheduleEvent, params map[string]string)
	l := logger.LoggerMock{}

	testTable := []struct {
		name           string
		scheduleEvents []model.ScheduleEvent
		user           model.User
		mockBehavior   mockBehavior
		params         map[string]string
	}{
		{
			name: "OK",
			scheduleEvents: []model.ScheduleEvent{
				{
					1,
					1,
					"First schedule event",
					360,
					1660794400,
					1660494400,
					1660494400,
				},
				{
					2,
					1,
					"Second schedule event",
					360,
					1660794400,
					1660494400,
					1660494400,
				},
			},
			user: model.User{
				ID:       1,
				Timezone: "UTC",
			},
			mockBehavior: func(s *mock_repository.MockISchedule, scheduleEvents []model.ScheduleEvent, params map[string]string) {
				s.EXPECT().List(params).Return(scheduleEvents, nil)
			},
			params: map[string]string{},
		},
	}

	for _, test := range testTable {
		// Init Dependencies
		c := gomock.NewController(t)
		defer c.Finish()

		repo := mock_repository.NewMockISchedule(c)
		ctx := helpers.SetUserToContext(test.user, context.Background())
		test.mockBehavior(repo, test.scheduleEvents, test.params)
		service := NewScheduleService(repo, l)

		// Test
		result, err := service.List(ctx, test.params)

		// Assert
		assert.Equal(t, test.scheduleEvents, result)
		assert.Nil(t, err)
	}
}

func TestService_Create(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockISchedule, inputScheduleEvents model.ScheduleEvent, outputScheduleEvents model.ScheduleEvent)
	l := logger.LoggerMock{}

	testTable := []struct {
		name                string
		inputScheduleEvent  model.ScheduleEvent
		outputScheduleEvent model.ScheduleEvent
		user                model.User
		mockBehavior        mockBehavior
		params              map[string]string
	}{
		{
			name: "OK",
			inputScheduleEvent: model.ScheduleEvent{
				0,
				1,
				"Schedule event",
				360,
				1660794400,
				0,
				0,
			},
			outputScheduleEvent: model.ScheduleEvent{
				1,
				1,
				"Schedule event",
				360,
				1660794400,
				1660494400,
				1660494400,
			},
			user: model.User{
				ID:       1,
				Timezone: "UTC",
			},
			mockBehavior: func(s *mock_repository.MockISchedule, inputScheduleEvents model.ScheduleEvent, outputScheduleEvents model.ScheduleEvent) {
				s.EXPECT().Create(inputScheduleEvents).Return(outputScheduleEvents, nil)
			},
			params: map[string]string{},
		},
	}

	for _, test := range testTable {
		// Init Dependencies
		c := gomock.NewController(t)
		defer c.Finish()

		repo := mock_repository.NewMockISchedule(c)
		ctx := helpers.SetUserToContext(test.user, context.Background())
		test.mockBehavior(repo, test.inputScheduleEvent, test.outputScheduleEvent)
		service := NewScheduleService(repo, l)

		// Test
		result, err := service.Create(ctx, test.inputScheduleEvent)

		// Assert
		assert.Equal(t, test.outputScheduleEvent, result)
		assert.Nil(t, err)
	}
}

func TestService_Update(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockISchedule, ID int, inputScheduleEvents model.ScheduleEvent, outputScheduleEvents model.ScheduleEvent)
	l := logger.LoggerMock{}

	testTable := []struct {
		name                string
		inputScheduleEvent  model.ScheduleEvent
		outputScheduleEvent model.ScheduleEvent
		user                model.User
		mockBehavior        mockBehavior
		params              map[string]string
	}{
		{
			name: "OK",
			inputScheduleEvent: model.ScheduleEvent{
				0,
				1,
				"Schedule event",
				360,
				1660794400,
				0,
				0,
			},
			outputScheduleEvent: model.ScheduleEvent{
				1,
				1,
				"Schedule event",
				360,
				1660794400,
				1660494400,
				1660494400,
			},
			user: model.User{
				ID:       1,
				Timezone: "UTC",
			},
			mockBehavior: func(
				s *mock_repository.MockISchedule,
				ID int,
				inputScheduleEvent model.ScheduleEvent,
				outputScheduleEvent model.ScheduleEvent,
			) {
				s.EXPECT().
					Update(ID, inputScheduleEvent).
					Return(outputScheduleEvent, nil)
			},
			params: map[string]string{},
		},
	}

	for _, test := range testTable {
		// Init Dependencies
		c := gomock.NewController(t)
		defer c.Finish()

		repo := mock_repository.NewMockISchedule(c)
		ctx := helpers.SetUserToContext(test.user, context.Background())
		test.mockBehavior(repo, test.outputScheduleEvent.ID, test.inputScheduleEvent, test.outputScheduleEvent)
		service := NewScheduleService(repo, l)

		// Test
		result, err := service.Update(ctx, test.outputScheduleEvent.ID, test.inputScheduleEvent)

		// Assert
		assert.Equal(t, test.outputScheduleEvent, result)
		assert.Nil(t, err)
	}
}

func TestService_Delete(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockISchedule, ID int)
	l := logger.LoggerMock{}

	testTable := []struct {
		name         string
		inputId      int
		user         model.User
		mockBehavior mockBehavior
		params       map[string]string
	}{
		{
			name:    "OK",
			inputId: 1,
			user: model.User{
				ID:       1,
				Timezone: "UTC",
			},
			mockBehavior: func(s *mock_repository.MockISchedule, ID int) {
				s.EXPECT().
					Delete(ID).
					Return(nil)
			},
			params: map[string]string{},
		},
	}

	for _, test := range testTable {
		// Init Dependencies
		c := gomock.NewController(t)
		defer c.Finish()

		repo := mock_repository.NewMockISchedule(c)
		ctx := helpers.SetUserToContext(test.user, context.Background())
		test.mockBehavior(repo, test.inputId)
		service := NewScheduleService(repo, l)

		// Test
		err := service.Delete(ctx, test.inputId)

		// Assert
		assert.Nil(t, err)
	}
}
