package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/i-vasilkov/go-todo-app/internal/domain"
	mock_service "github.com/i-vasilkov/go-todo-app/internal/service/mocks"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestTaskService_Get(t *testing.T) {
	type mockBehaviour func(rep *mock_service.MockTaskRepositoryI, taskId, userId string, task domain.Task)

	testCases := []struct {
		name   string
		taskId string
		userId string
		mock   mockBehaviour
		task   domain.Task
		err    error
	}{
		{
			name:   "OK",
			taskId: "taskId",
			userId: "userId",
			mock: func(rep *mock_service.MockTaskRepositoryI, taskId, userId string, task domain.Task) {
				rep.EXPECT().Get(context.Background(), taskId, userId).Return(task, nil)
			},
			task: domain.Task{
				Id:     "taskId",
				UserId: "userId",
			},
			err: nil,
		},
		{
			name:   "Repository error",
			taskId: "taskId",
			userId: "userId",
			mock: func(rep *mock_service.MockTaskRepositoryI, taskId, userId string, task domain.Task) {
				rep.EXPECT().Get(context.Background(), taskId, userId).Return(task, errors.New("repository error"))
			},
			task: domain.Task{},
			err:  errors.New("repository error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repository := mock_service.NewMockTaskRepositoryI(ctrl)
			testCase.mock(repository, testCase.taskId, testCase.userId, testCase.task)

			service := NewTaskService(repository)
			task, err := service.Get(context.Background(), testCase.taskId, testCase.userId)

			assert.Equal(t, task, testCase.task)
			assert.Equal(t, err, testCase.err)
		})
	}
}

func TestTaskService_GetAll(t *testing.T) {
	type mockBehaviour func(rep *mock_service.MockTaskRepositoryI, userId string, tasks []domain.Task)

	testCases := []struct {
		name   string
		userId string
		mock   mockBehaviour
		tasks  []domain.Task
		err    error
	}{
		{
			name:   "OK",
			userId: "userId",
			mock: func(rep *mock_service.MockTaskRepositoryI, userId string, tasks []domain.Task) {
				rep.EXPECT().GetAll(context.Background(), userId).Return(tasks, nil)
			},
			tasks: []domain.Task{
				{
					Id:     "taskId",
					UserId: "userId",
				},
			},
			err: nil,
		},
		{
			name:   "Repository error",
			userId: "userId",
			mock: func(rep *mock_service.MockTaskRepositoryI, userId string, tasks []domain.Task) {
				rep.EXPECT().GetAll(context.Background(), userId).Return(tasks, errors.New("repository error"))
			},
			tasks: []domain.Task{},
			err:   errors.New("repository error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repository := mock_service.NewMockTaskRepositoryI(ctrl)
			testCase.mock(repository, testCase.userId, testCase.tasks)

			service := NewTaskService(repository)
			tasks, err := service.GetAll(context.Background(), testCase.userId)

			assert.Equal(t, tasks, testCase.tasks)
			assert.Equal(t, err, testCase.err)
		})
	}
}

func TestTaskService_Delete(t *testing.T) {
	type mockBehaviour func(rep *mock_service.MockTaskRepositoryI, taskId, userId string)

	testCases := []struct {
		name   string
		taskId string
		userId string
		mock   mockBehaviour
		err    error
	}{
		{
			name:   "OK",
			taskId: "taskId",
			userId: "userId",
			mock: func(rep *mock_service.MockTaskRepositoryI, taskId, userId string) {
				rep.EXPECT().Delete(context.Background(), taskId, userId).Return(nil)
			},
			err: nil,
		},
		{
			name:   "Repository error",
			taskId: "taskId",
			userId: "userId",
			mock: func(rep *mock_service.MockTaskRepositoryI, taskId, userId string) {
				rep.EXPECT().Delete(context.Background(), taskId, userId).Return(errors.New("repository error"))
			},
			err: errors.New("repository error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repository := mock_service.NewMockTaskRepositoryI(ctrl)
			testCase.mock(repository, testCase.taskId, testCase.userId)

			service := NewTaskService(repository)
			err := service.Delete(context.Background(), testCase.taskId, testCase.userId)

			assert.Equal(t, err, testCase.err)
		})
	}
}

func TestTaskService_Update(t *testing.T) {
	type mockBehaviour func(rep *mock_service.MockTaskRepositoryI, taskId, userId string, in domain.UpdateTaskInput, task domain.Task)

	testCases := []struct {
		name   string
		taskId string
		userId string
		input  domain.UpdateTaskInput
		mock   mockBehaviour
		task   domain.Task
		err    error
	}{
		{
			name:   "OK",
			taskId: "taskId",
			userId: "userId",
			input:  domain.UpdateTaskInput{Name: "updated"},
			mock: func(rep *mock_service.MockTaskRepositoryI, taskId, userId string, in domain.UpdateTaskInput, task domain.Task) {
				rep.EXPECT().Update(context.Background(), taskId, userId, in).Return(task, nil)
			},
			task: domain.Task{
				Id:     "taskId",
				UserId: "userId",
				Name:   "updated",
			},
			err: nil,
		},
		{
			name:   "Repository error",
			taskId: "taskId",
			userId: "userId",
			input:  domain.UpdateTaskInput{Name: "updated"},
			mock: func(rep *mock_service.MockTaskRepositoryI, taskId, userId string, in domain.UpdateTaskInput, task domain.Task) {
				rep.EXPECT().Update(context.Background(), taskId, userId, in).Return(task, errors.New("repository error"))
			},
			task: domain.Task{},
			err:  errors.New("repository error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repository := mock_service.NewMockTaskRepositoryI(ctrl)
			testCase.mock(repository, testCase.taskId, testCase.userId, testCase.input, testCase.task)

			service := NewTaskService(repository)
			task, err := service.Update(context.Background(), testCase.taskId, testCase.userId, testCase.input)

			assert.Equal(t, task, testCase.task)
			assert.Equal(t, err, testCase.err)
		})
	}
}

func TestTaskService_Create(t *testing.T) {
	type mockBehaviour func(rep *mock_service.MockTaskRepositoryI, userId string, in domain.CreateTaskInput, task domain.Task)

	testCases := []struct {
		name   string
		userId string
		input  domain.CreateTaskInput
		mock   mockBehaviour
		task   domain.Task
		err    error
	}{
		{
			name:   "OK",
			userId: "userId",
			input:  domain.CreateTaskInput{Name: "updated"},
			mock: func(rep *mock_service.MockTaskRepositoryI, userId string, in domain.CreateTaskInput, task domain.Task) {
				rep.EXPECT().Create(context.Background(), userId, in).Return(task, nil)
			},
			task: domain.Task{
				Id:     "taskId",
				UserId: "userId",
			},
			err: nil,
		},
		{
			name:   "Repository error",
			userId: "userId",
			input:  domain.CreateTaskInput{Name: "updated"},
			mock: func(rep *mock_service.MockTaskRepositoryI, userId string, in domain.CreateTaskInput, task domain.Task) {
				rep.EXPECT().Create(context.Background(), userId, in).Return(task, errors.New("repository error"))
			},
			task: domain.Task{},
			err:  errors.New("repository error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repository := mock_service.NewMockTaskRepositoryI(ctrl)
			testCase.mock(repository, testCase.userId, testCase.input, testCase.task)

			service := NewTaskService(repository)
			task, err := service.Create(context.Background(), testCase.userId, testCase.input)

			assert.Equal(t, task, testCase.task)
			assert.Equal(t, err, testCase.err)
		})
	}
}
