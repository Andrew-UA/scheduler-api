package schedule

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"scheduler/internal/model"
	mock_repository "scheduler/internal/repository/mocks"
	"testing"
)

func TestService_Show(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockISchedule, scheduleEvent model.ScheduleEvent, ID int)

	testTable := []struct {
		name          string
		scheduleEvent model.ScheduleEvent
		mockBehavior  mockBehavior
	}{
		{
			name: "OK",
			scheduleEvent: model.ScheduleEvent{
				1,
				"Test Schedule event",
				360,
				1660794400,
				1660494400,
				1660494400,
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
		test.mockBehavior(repo, test.scheduleEvent, test.scheduleEvent.ID)
		service := NewService(repo)

		// Test
		result, err := service.Show(test.scheduleEvent.ID)

		// Assert
		assert.Equal(t, test.scheduleEvent, result)
		assert.Nil(t, err)

	}

}

func TestService_List(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockISchedule, scheduleEvents []model.ScheduleEvent, params map[string]string)

	testTable := []struct {
		name           string
		scheduleEvents []model.ScheduleEvent
		mockBehavior   mockBehavior
		params         map[string]string
	}{
		{
			name: "OK",
			scheduleEvents: []model.ScheduleEvent{
				{
					1,
					"First schedule event",
					360,
					1660794400,
					1660494400,
					1660494400,
				},
				{
					2,
					"Second schedule event",
					360,
					1660794400,
					1660494400,
					1660494400,
				},
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
		test.mockBehavior(repo, test.scheduleEvents, test.params)
		service := NewService(repo)

		// Test
		result, err := service.List(test.params)

		// Assert
		assert.Equal(t, test.scheduleEvents, result)
		assert.Nil(t, err)
	}
}

func TestService_Create(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockISchedule, inputScheduleEvents model.ScheduleEvent, outputScheduleEvents model.ScheduleEvent)

	testTable := []struct {
		name                string
		inputScheduleEvent  model.ScheduleEvent
		outputScheduleEvent model.ScheduleEvent
		mockBehavior        mockBehavior
		params              map[string]string
	}{
		{
			name: "OK",
			inputScheduleEvent: model.ScheduleEvent{
				0,
				"Schedule event",
				360,
				1660794400,
				0,
				0,
			},
			outputScheduleEvent: model.ScheduleEvent{
				1,
				"Schedule event",
				360,
				1660794400,
				1660494400,
				1660494400,
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
		test.mockBehavior(repo, test.inputScheduleEvent, test.outputScheduleEvent)
		service := NewService(repo)

		// Test
		result, err := service.Create(test.inputScheduleEvent)

		// Assert
		assert.Equal(t, test.outputScheduleEvent, result)
		assert.Nil(t, err)
	}
}

func TestService_Update(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockISchedule, ID int, inputScheduleEvents model.ScheduleEvent, outputScheduleEvents model.ScheduleEvent)

	testTable := []struct {
		name                string
		inputScheduleEvent  model.ScheduleEvent
		outputScheduleEvent model.ScheduleEvent
		mockBehavior        mockBehavior
		params              map[string]string
	}{
		{
			name: "OK",
			inputScheduleEvent: model.ScheduleEvent{
				0,
				"Schedule event",
				360,
				1660794400,
				0,
				0,
			},
			outputScheduleEvent: model.ScheduleEvent{
				1,
				"Schedule event",
				360,
				1660794400,
				1660494400,
				1660494400,
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
		test.mockBehavior(repo, test.outputScheduleEvent.ID, test.inputScheduleEvent, test.outputScheduleEvent)
		service := NewService(repo)

		// Test
		result, err := service.Update(test.outputScheduleEvent.ID, test.inputScheduleEvent)

		// Assert
		assert.Equal(t, test.outputScheduleEvent, result)
		assert.Nil(t, err)
	}
}

func TestService_Delete(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockISchedule, ID int)

	testTable := []struct {
		name                string
		inputId				int
		mockBehavior        mockBehavior
		params              map[string]string
	}{
		{
			name: "OK",
			inputId: 1,
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
		test.mockBehavior(repo, test.inputId)
		service := NewService(repo)

		// Test
		err := service.Delete(test.inputId)

		// Assert
		assert.Nil(t, err)
	}
}
