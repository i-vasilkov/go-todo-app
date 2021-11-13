package http

import (
	"bytes"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/i-vasilkov/go-todo-app/internal/domain"
	mock_http "github.com/i-vasilkov/go-todo-app/internal/handler/http/mocks"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_SignIn(t *testing.T) {
	type mockBehavior func(s *mock_http.MockAuthServiceI, in domain.LoginUserInput)

	testCases := []struct {
		name           string
		inputReqBody   string
		inputObj       domain.LoginUserInput
		mockBehavior   mockBehavior
		respStatusCode int
		respBody       string
	}{
		{
			name:         "OK",
			inputReqBody: `{"login":"test","password":"test"}`,
			inputObj: domain.LoginUserInput{
				Login:    "test",
				Password: "test",
			},
			mockBehavior: func(s *mock_http.MockAuthServiceI, in domain.LoginUserInput) {
				s.EXPECT().SignIn(context.Background(), in).Return("token", nil)
			},
			respStatusCode: http.StatusOK,
			respBody:       `{"success":true,"data":"token"}`,
		},
		{
			name:           "Empty login input",
			inputReqBody:   `{"password":"test"}`,
			inputObj:       domain.LoginUserInput{},
			mockBehavior:   func(s *mock_http.MockAuthServiceI, in domain.LoginUserInput) {},
			respStatusCode: http.StatusBadRequest,
			respBody:       `{"success":false,"messages":["invalid 'Login' input"]}`,
		},
		{
			name:           "Empty password input",
			inputReqBody:   `{"login":"test"}`,
			inputObj:       domain.LoginUserInput{},
			mockBehavior:   func(s *mock_http.MockAuthServiceI, in domain.LoginUserInput) {},
			respStatusCode: http.StatusBadRequest,
			respBody:       `{"success":false,"messages":["invalid 'Password' input"]}`,
		},
		{
			name:           "Empty input",
			inputReqBody:   `{}`,
			inputObj:       domain.LoginUserInput{},
			mockBehavior:   func(s *mock_http.MockAuthServiceI, in domain.LoginUserInput) {},
			respStatusCode: http.StatusBadRequest,
			respBody:       `{"success":false,"messages":["invalid 'Login' input","invalid 'Password' input"]}`,
		},
		{
			name:           "Empty body",
			inputReqBody:   ``,
			inputObj:       domain.LoginUserInput{},
			mockBehavior:   func(s *mock_http.MockAuthServiceI, in domain.LoginUserInput) {},
			respStatusCode: http.StatusBadRequest,
			respBody:       `{"success":false,"messages":["invalid input body"]}`,
		},
		{
			name:         "Service error",
			inputReqBody: `{"login":"test","password":"test"}`,
			inputObj: domain.LoginUserInput{
				Login:    "test",
				Password: "test",
			},
			mockBehavior: func(s *mock_http.MockAuthServiceI, in domain.LoginUserInput) {
				s.EXPECT().SignIn(context.Background(), in).Return("", errors.New("service error"))
			},
			respStatusCode: http.StatusInternalServerError,
			respBody:       `{"success":false,"messages":["service error"]}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_http.NewMockAuthServiceI(c)
			testCase.mockBehavior(auth, testCase.inputObj)

			s := &Services{Auth: auth}
			h := NewHandler(s)

			r := gin.New()
			r.POST("/sign-in", h.authSignIn)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-in", bytes.NewBufferString(testCase.inputReqBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.respStatusCode)
			assert.Equal(t, w.Body.String(), testCase.respBody)
		})
	}
}

func TestHandler_SignUp(t *testing.T) {
	type mockBehavior func(s *mock_http.MockAuthServiceI, in domain.CreateUserInput)

	testCases := []struct {
		name          string
		inputReqBody  string
		inputObj      domain.CreateUserInput
		mockBehavior   mockBehavior
		respStatusCode int
		respBody       string
	}{
		{
			name:         "OK",
			inputReqBody: `{"login":"test","password":"test"}`,
			inputObj: domain.CreateUserInput{
				Login:    "test",
				Password: "test",
			},
			mockBehavior: func(s *mock_http.MockAuthServiceI, in domain.CreateUserInput) {
				s.EXPECT().SignUp(context.Background(), in).Return("token", nil)
			},
			respStatusCode: http.StatusOK,
			respBody:       `{"success":true,"data":"token"}`,
		},
		{
			name:           "Empty login input",
			inputReqBody:   `{"password":"test"}`,
			inputObj:       domain.CreateUserInput{},
			mockBehavior:   func(s *mock_http.MockAuthServiceI, in domain.CreateUserInput) {},
			respStatusCode: http.StatusBadRequest,
			respBody:       `{"success":false,"messages":["invalid 'Login' input"]}`,
		},
		{
			name:           "Empty password input",
			inputReqBody:   `{"login":"test"}`,
			inputObj:       domain.CreateUserInput{},
			mockBehavior:   func(s *mock_http.MockAuthServiceI, in domain.CreateUserInput) {},
			respStatusCode: http.StatusBadRequest,
			respBody:       `{"success":false,"messages":["invalid 'Password' input"]}`,
		},
		{
			name:           "Empty input",
			inputReqBody:   `{}`,
			inputObj:       domain.CreateUserInput{},
			mockBehavior:   func(s *mock_http.MockAuthServiceI, in domain.CreateUserInput) {},
			respStatusCode: http.StatusBadRequest,
			respBody:       `{"success":false,"messages":["invalid 'Login' input","invalid 'Password' input"]}`,
		},
		{
			name:           "Empty body",
			inputReqBody:   ``,
			inputObj:       domain.CreateUserInput{},
			mockBehavior:   func(s *mock_http.MockAuthServiceI, in domain.CreateUserInput) {},
			respStatusCode: http.StatusBadRequest,
			respBody:       `{"success":false,"messages":["invalid input body"]}`,
		},
		{
			name:         "Service error",
			inputReqBody: `{"login":"test","password":"test"}`,
			inputObj: domain.CreateUserInput{
				Login:    "test",
				Password: "test",
			},
			mockBehavior: func(s *mock_http.MockAuthServiceI, in domain.CreateUserInput) {
				s.EXPECT().SignUp(context.Background(), in).Return("", errors.New("service error"))
			},
			respStatusCode: http.StatusInternalServerError,
			respBody:       `{"success":false,"messages":["service error"]}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_http.NewMockAuthServiceI(c)
			testCase.mockBehavior(auth, testCase.inputObj)

			services := Services{Auth: auth}
			handler := NewHandler(&services)

			router := gin.New()
			router.POST("/sign-up", handler.authSignUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(testCase.inputReqBody))

			router.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.respStatusCode)
			assert.Equal(t, w.Body.String(), testCase.respBody)
		})
	}
}
