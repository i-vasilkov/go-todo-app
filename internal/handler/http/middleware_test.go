package http

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mock_http "github.com/i-vasilkov/go-todo-app/internal/handler/http/mocks"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_AuthMiddleware(t *testing.T) {
	type mockBehaviour func(s *mock_http.MockAuthServiceI, token string)

	testCases := []struct {
		name          string
		headerIsSet   bool
		header        string
		token         string
		mockBehaviour  mockBehaviour
		respStatusCode int
		respBody       string
	}{
		{
			name:        "OK",
			headerIsSet: true,
			header:      "Bearer token",
			token:       "token",
			mockBehaviour: func(s *mock_http.MockAuthServiceI, token string) {
				s.EXPECT().CheckToken(context.Background(), token).Return("userId", nil)
			},
			respStatusCode: http.StatusOK,
			respBody:       "userId",
		},
		{
			name:           "Header not found",
			headerIsSet:    false,
			header:         "",
			token:          "",
			mockBehaviour:  func(s *mock_http.MockAuthServiceI, token string) {},
			respStatusCode: http.StatusUnauthorized,
			respBody:       `{"success":false,"messages":["empty auth header"]}`,
		},
		{
			name:           "Header is empty",
			headerIsSet:    true,
			header:         "",
			token:          "",
			mockBehaviour:  func(s *mock_http.MockAuthServiceI, token string) {},
			respStatusCode: http.StatusUnauthorized,
			respBody:       `{"success":false,"messages":["empty auth header"]}`,
		},
		{
			name:           "Header without Bearer",
			headerIsSet:    true,
			header:         "token",
			token:          "",
			mockBehaviour:  func(s *mock_http.MockAuthServiceI, token string) {},
			respStatusCode: http.StatusUnauthorized,
			respBody:       `{"success":false,"messages":["not valid auth header"]}`,
		},
		{
			name:           "Header without token",
			headerIsSet:    true,
			header:         "Bearer",
			token:          "",
			mockBehaviour:  func(s *mock_http.MockAuthServiceI, token string) {},
			respStatusCode: http.StatusUnauthorized,
			respBody:       `{"success":false,"messages":["not valid auth header"]}`,
		},
		{
			name:           "Header with invalid BearerName",
			headerIsSet:    true,
			header:         "NotBearer token",
			token:          "",
			mockBehaviour:  func(s *mock_http.MockAuthServiceI, token string) {},
			respStatusCode: http.StatusUnauthorized,
			respBody:       `{"success":false,"messages":["not valid auth header"]}`,
		},
		{
			name:          "Parse token error",
			headerIsSet:   true,
			header:        "Bearer token",
			token:         "token",
			mockBehaviour: func(s *mock_http.MockAuthServiceI, token string) {
				s.EXPECT().CheckToken(context.Background(), token).Return("", errors.New("service error"))
			},
			respStatusCode: http.StatusUnauthorized,
			respBody:       `{"success":false,"messages":["service error"]}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_http.NewMockAuthServiceI(c)
			testCase.mockBehaviour(auth, testCase.token)

			services := &Services{Auth: auth}
			handler := NewHandler(services)

			router := gin.New()
			router.GET("/protected", handler.AuthMiddleware, func(c *gin.Context) {
				id, _ := c.Get(userCtx)
				c.String(http.StatusOK, id.(string))
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/protected", nil)
			if testCase.headerIsSet {
				req.Header.Set(authHeaderName, testCase.header)
			}

			router.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.respStatusCode)
			assert.Equal(t, w.Body.String(), testCase.respBody)
		})
	}
}
