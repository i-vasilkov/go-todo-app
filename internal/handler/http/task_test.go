package http

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/i-vasilkov/go-todo-app/internal/domain"
	mock_http "github.com/i-vasilkov/go-todo-app/internal/handler/http/mocks"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_taskGetOne(t *testing.T) {
	type mockBehavior func(s *mock_http.MockTaskServiceI, id, userId string, task domain.Task)

	testCases := []struct {
		name           string
		taskId         string
		userId         string
		task           domain.Task
		mockBehavior   mockBehavior
		respStatusCode int
		respBody       string
	}{
		{
			name:   "OK",
			taskId: "taskId",
			userId: "userId",
			task: domain.Task{
				Id:        "taskId",
				Name:      "test",
				UserId:    "userId",
				CreatedAt: time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
			},
			mockBehavior: func(s *mock_http.MockTaskServiceI, id, userId string, task domain.Task) {
				s.EXPECT().Get(context.Background(), id, userId).Return(task, nil)
			},
			respStatusCode: http.StatusOK,
			respBody:       `{"success":true,"data":{"id":"taskId","name":"test","user_id":"userId","created_at":"2020-01-01T00:00:00Z","updated_at":"2020-01-01T00:00:00Z"}}`,
		},
		{
			name:           "Empty taskId",
			taskId:         "",
			userId:         "userId",
			task:           domain.Task{},
			mockBehavior:   func(s *mock_http.MockTaskServiceI, id, userId string, task domain.Task) {},
			respStatusCode: http.StatusBadRequest,
			respBody:       `{"success":false,"messages":["empty task id"]}`,
		},
		{
			name:           "Empty UserId",
			taskId:         "taskId",
			userId:         "",
			task:           domain.Task{},
			mockBehavior:   func(s *mock_http.MockTaskServiceI, id, userId string, task domain.Task) {},
			respStatusCode: http.StatusUnauthorized,
			respBody:       `{"success":false,"messages":["not exists userId in context"]}`,
		},
		{
			name:   "Service error",
			taskId: "taskId",
			userId: "userId",
			task:   domain.Task{},
			mockBehavior: func(s *mock_http.MockTaskServiceI, id, userId string, task domain.Task) {
				s.EXPECT().Get(context.Background(), id, userId).Return(task, errors.New("service error"))
			},
			respStatusCode: http.StatusInternalServerError,
			respBody:       `{"success":false,"messages":["service error"]}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			taskService := mock_http.NewMockTaskServiceI(ctrl)
			testCase.mockBehavior(taskService, testCase.taskId, testCase.userId, testCase.task)

			services := &Services{Task: taskService}
			handler := NewHandler(services)

			router := gin.New()
			router.GET("/task/:id", func(ctx *gin.Context) {
				if testCase.userId != "" {
					ctx.Set(userCtx, testCase.userId)
				}
			}, handler.taskGetOne)

			w := httptest.NewRecorder()
			reqUrl := fmt.Sprintf("/task/%s", testCase.taskId)
			req := httptest.NewRequest("GET", reqUrl, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.respStatusCode)
			assert.Equal(t, w.Body.String(), testCase.respBody)
		})
	}
}

func TestHandler_taskDelete(t *testing.T) {
	type mockBehavior func(s *mock_http.MockTaskServiceI, id, userId string)

	testCases := []struct {
		name           string
		taskId         string
		userId         string
		mockBehavior   mockBehavior
		respStatusCode int
		respBody       string
	}{
		{
			name:   "OK",
			taskId: "taskId",
			userId: "userId",
			mockBehavior: func(s *mock_http.MockTaskServiceI, id, userId string) {
				s.EXPECT().Delete(context.Background(), id, userId).Return(nil)
			},
			respStatusCode: http.StatusOK,
			respBody:       `{"success":true,"data":null}`,
		},
		{
			name:           "Empty taskId",
			taskId:         "",
			userId:         "userId",
			mockBehavior:   func(s *mock_http.MockTaskServiceI, id, userId string) {},
			respStatusCode: http.StatusBadRequest,
			respBody:       `{"success":false,"messages":["empty task id"]}`,
		},
		{
			name:           "Empty UserId",
			taskId:         "taskId",
			userId:         "",
			mockBehavior:   func(s *mock_http.MockTaskServiceI, id, userId string) {},
			respStatusCode: http.StatusUnauthorized,
			respBody:       `{"success":false,"messages":["not exists userId in context"]}`,
		},
		{
			name:   "Service error",
			taskId: "taskId",
			userId: "userId",
			mockBehavior: func(s *mock_http.MockTaskServiceI, id, userId string) {
				s.EXPECT().Delete(context.Background(), id, userId).Return(errors.New("service error"))
			},
			respStatusCode: http.StatusInternalServerError,
			respBody:       `{"success":false,"messages":["service error"]}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			taskService := mock_http.NewMockTaskServiceI(ctrl)
			testCase.mockBehavior(taskService, testCase.taskId, testCase.userId)

			services := &Services{Task: taskService}
			handler := NewHandler(services)

			router := gin.New()
			router.DELETE("/task/:id", func(ctx *gin.Context) {
				if testCase.userId != "" {
					ctx.Set(userCtx, testCase.userId)
				}
			}, handler.taskDelete)

			w := httptest.NewRecorder()
			reqUrl := fmt.Sprintf("/task/%s", testCase.taskId)
			req := httptest.NewRequest("DELETE", reqUrl, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.respStatusCode)
			assert.Equal(t, w.Body.String(), testCase.respBody)
		})
	}
}

func TestHandler_taskUpdate(t *testing.T) {
	type mockBehavior func(s *mock_http.MockTaskServiceI, id, userId string, in domain.UpdateTaskInput, task domain.Task)

	testCases := []struct {
		name           string
		reqBody        string
		inputObj       domain.UpdateTaskInput
		taskId         string
		userId         string
		task           domain.Task
		mockBehavior   mockBehavior
		respStatusCode int
		respBody       string
	}{
		{
			name:     "OK",
			taskId:   "taskId",
			userId:   "userId",
			reqBody:  `{"name":"updated"}`,
			inputObj: domain.UpdateTaskInput{Name: "updated"},
			task: domain.Task{
				Id:        "taskId",
				Name:      "updated",
				UserId:    "userId",
				CreatedAt: time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
			},
			mockBehavior: func(s *mock_http.MockTaskServiceI, id, userId string, in domain.UpdateTaskInput, task domain.Task) {
				s.EXPECT().Update(context.Background(), id, userId, in).Return(task, nil)
			},
			respStatusCode: http.StatusOK,
			respBody:       `{"success":true,"data":{"id":"taskId","name":"updated","user_id":"userId","created_at":"2020-01-01T00:00:00Z","updated_at":"2020-01-01T00:00:00Z"}}`,
		},
		{
			name:           "Empty task.name",
			taskId:         "",
			userId:         "",
			reqBody:        `{}`,
			inputObj:       domain.UpdateTaskInput{Name: "updated"},
			task:           domain.Task{},
			mockBehavior:   func(s *mock_http.MockTaskServiceI, id, userId string, in domain.UpdateTaskInput, task domain.Task) {},
			respStatusCode: http.StatusBadRequest,
			respBody:       `{"success":false,"messages":["invalid 'Name' input"]}`,
		},
		{
			name:           "Empty taskId",
			taskId:         "",
			userId:         "userId",
			reqBody:        `{"name":"updated"}`,
			inputObj:       domain.UpdateTaskInput{Name: "updated"},
			task:           domain.Task{},
			mockBehavior:   func(s *mock_http.MockTaskServiceI, id, userId string, in domain.UpdateTaskInput, task domain.Task) {},
			respStatusCode: http.StatusBadRequest,
			respBody:       `{"success":false,"messages":["empty task id"]}`,
		},
		{
			name:           "Empty UserId",
			taskId:         "taskId",
			userId:         "",
			reqBody:        `{"name":"updated"}`,
			inputObj:       domain.UpdateTaskInput{Name: "updated"},
			task:           domain.Task{},
			mockBehavior:   func(s *mock_http.MockTaskServiceI, id, userId string, in domain.UpdateTaskInput, task domain.Task) {},
			respStatusCode: http.StatusUnauthorized,
			respBody:       `{"success":false,"messages":["not exists userId in context"]}`,
		},
		{
			name:     "Service error",
			taskId:   "taskId",
			userId:   "userId",
			reqBody:  `{"name":"updated"}`,
			inputObj: domain.UpdateTaskInput{Name: "updated"},
			task:     domain.Task{},
			mockBehavior: func(s *mock_http.MockTaskServiceI, id, userId string, in domain.UpdateTaskInput, task domain.Task) {
				s.EXPECT().Update(context.Background(), id, userId, in).Return(task, errors.New("service error"))
			},
			respStatusCode: http.StatusInternalServerError,
			respBody:       `{"success":false,"messages":["service error"]}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			taskService := mock_http.NewMockTaskServiceI(ctrl)
			testCase.mockBehavior(taskService, testCase.taskId, testCase.userId, testCase.inputObj, testCase.task)

			services := &Services{Task: taskService}
			handler := NewHandler(services)

			router := gin.New()
			router.PUT("/task/:id", func(ctx *gin.Context) {
				if testCase.userId != "" {
					ctx.Set(userCtx, testCase.userId)
				}
			}, handler.taskUpdate)

			w := httptest.NewRecorder()
			reqUrl := fmt.Sprintf("/task/%s", testCase.taskId)
			req := httptest.NewRequest("PUT", reqUrl, bytes.NewBufferString(testCase.reqBody))

			router.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.respStatusCode)
			assert.Equal(t, w.Body.String(), testCase.respBody)
		})
	}
}

func TestHandler_taskCreate(t *testing.T) {
	type mockBehavior func(s *mock_http.MockTaskServiceI, userId string, in domain.CreateTaskInput, task domain.Task)

	testCases := []struct {
		name           string
		reqBody        string
		inputObj       domain.CreateTaskInput
		userId         string
		task           domain.Task
		mockBehavior   mockBehavior
		respStatusCode int
		respBody       string
	}{
		{
			name:     "OK",
			userId:   "userId",
			reqBody:  `{"name":"test"}`,
			inputObj: domain.CreateTaskInput{Name: "test"},
			task: domain.Task{
				Id:        "taskId",
				Name:      "test",
				UserId:    "userId",
				CreatedAt: time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
			},
			mockBehavior: func(s *mock_http.MockTaskServiceI, userId string, in domain.CreateTaskInput, task domain.Task) {
				s.EXPECT().Create(context.Background(), userId, in).Return(task, nil)
			},
			respStatusCode: http.StatusOK,
			respBody:       `{"success":true,"data":{"id":"taskId","name":"test","user_id":"userId","created_at":"2020-01-01T00:00:00Z","updated_at":"2020-01-01T00:00:00Z"}}`,
		},
		{
			name:           "Empty task.name",
			userId:         "",
			reqBody:        `{}`,
			inputObj:       domain.CreateTaskInput{Name: "test"},
			task:           domain.Task{},
			mockBehavior:   func(s *mock_http.MockTaskServiceI, userId string, in domain.CreateTaskInput, task domain.Task) {},
			respStatusCode: http.StatusBadRequest,
			respBody:       `{"success":false,"messages":["invalid 'Name' input"]}`,
		},
		{
			name:           "Empty UserId",
			userId:         "",
			reqBody:        `{"name":"test"}`,
			inputObj:       domain.CreateTaskInput{Name: "test"},
			task:           domain.Task{},
			mockBehavior:   func(s *mock_http.MockTaskServiceI, userId string, in domain.CreateTaskInput, task domain.Task) {},
			respStatusCode: http.StatusUnauthorized,
			respBody:       `{"success":false,"messages":["not exists userId in context"]}`,
		},
		{
			name:     "Service error",
			userId:   "userId",
			reqBody:  `{"name":"test"}`,
			inputObj: domain.CreateTaskInput{Name: "test"},
			task:     domain.Task{},
			mockBehavior: func(s *mock_http.MockTaskServiceI, userId string, in domain.CreateTaskInput, task domain.Task) {
				s.EXPECT().Create(context.Background(), userId, in).Return(task, errors.New("service error"))
			},
			respStatusCode: http.StatusInternalServerError,
			respBody:       `{"success":false,"messages":["service error"]}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			taskService := mock_http.NewMockTaskServiceI(ctrl)
			testCase.mockBehavior(taskService, testCase.userId, testCase.inputObj, testCase.task)

			services := &Services{Task: taskService}
			handler := NewHandler(services)

			router := gin.New()
			router.POST("/task", func(ctx *gin.Context) {
				if testCase.userId != "" {
					ctx.Set(userCtx, testCase.userId)
				}
			}, handler.taskCreate)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/task", bytes.NewBufferString(testCase.reqBody))

			router.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.respStatusCode)
			assert.Equal(t, w.Body.String(), testCase.respBody)
		})
	}
}

func TestHandler_taskGetAll(t *testing.T) {
	type mockBehavior func(s *mock_http.MockTaskServiceI, userId string, tasks []domain.Task)

	testCases := []struct {
		name           string
		userId         string
		tasks          []domain.Task
		mockBehavior   mockBehavior
		respStatusCode int
		respBody       string
	}{
		{
			name:   "OK",
			userId: "userId",
			tasks: []domain.Task{
				{
					Id:        "taskId",
					Name:      "test",
					UserId:    "userId",
					CreatedAt: time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
				},
			},
			mockBehavior: func(s *mock_http.MockTaskServiceI, userId string, tasks []domain.Task) {
				s.EXPECT().GetAll(context.Background(), userId).Return(tasks, nil)
			},
			respStatusCode: http.StatusOK,
			respBody:       `{"success":true,"data":[{"id":"taskId","name":"test","user_id":"userId","created_at":"2020-01-01T00:00:00Z","updated_at":"2020-01-01T00:00:00Z"}]}`,
		},
		{
			name:           "Empty UserId",
			userId:         "",
			tasks:          nil,
			mockBehavior:   func(s *mock_http.MockTaskServiceI, userId string, tasks []domain.Task) {},
			respStatusCode: http.StatusUnauthorized,
			respBody:       `{"success":false,"messages":["not exists userId in context"]}`,
		},
		{
			name:   "Service error",
			userId: "userId",
			tasks:   nil,
			mockBehavior: func(s *mock_http.MockTaskServiceI, userId string, tasks []domain.Task) {
				s.EXPECT().GetAll(context.Background(), userId).Return(tasks, errors.New("service error"))
			},
			respStatusCode: http.StatusInternalServerError,
			respBody:       `{"success":false,"messages":["service error"]}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			taskService := mock_http.NewMockTaskServiceI(ctrl)
			testCase.mockBehavior(taskService, testCase.userId, testCase.tasks)

			services := &Services{Task: taskService}
			handler := NewHandler(services)

			router := gin.New()
			router.GET("/task", func(ctx *gin.Context) {
				if testCase.userId != "" {
					ctx.Set(userCtx, testCase.userId)
				}
			}, handler.taskGetAll)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/task", nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.respStatusCode)
			assert.Equal(t, w.Body.String(), testCase.respBody)
		})
	}
}
