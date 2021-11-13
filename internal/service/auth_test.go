package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/i-vasilkov/go-todo-app/internal/domain"
	"github.com/i-vasilkov/go-todo-app/internal/repository/mongorep"
	mock_service "github.com/i-vasilkov/go-todo-app/internal/service/mocks"
	mock_jwt "github.com/i-vasilkov/go-todo-app/pkg/auth/jwt/mocks"
	"github.com/i-vasilkov/go-todo-app/pkg/hash"
	mock_hash "github.com/i-vasilkov/go-todo-app/pkg/hash/mocks"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestAuthService_SignIn(t *testing.T) {
	type tokenManagerMockBehaviour func(tm *mock_jwt.MockTokenManagerI, in domain.LoginUserInput)
	type hasherMockBehaviour func(h *mock_hash.MockHasher, in domain.LoginUserInput)
	type repositoryMockBehaviour func(r *mock_service.MockUserRepositoryI, in domain.LoginUserInput)

	testCases := []struct {
		name             string
		input            domain.LoginUserInput
		hasherMock       hasherMockBehaviour
		repositoryMock   repositoryMockBehaviour
		tokenManagerMock tokenManagerMockBehaviour
		token            string
		err              error
	}{
		{
			name: "OK",
			input: domain.LoginUserInput{
				Login:    "test",
				Password: "test",
			},
			hasherMock: func(h *mock_hash.MockHasher, in domain.LoginUserInput) {
				h.EXPECT().Hash(in.Password).Return(in.Password, nil)
			},
			repositoryMock: func(r *mock_service.MockUserRepositoryI, in domain.LoginUserInput) {
				r.EXPECT().GetByCredentials(context.Background(), in).Return(domain.User{Id: "userId"}, nil)
			},
			tokenManagerMock: func(tm *mock_jwt.MockTokenManagerI, in domain.LoginUserInput) {
				tm.EXPECT().NewToken("userId").Return("token", nil)
			},
			token: "token",
			err:   nil,
		},
		{
			name: "Hasher error",
			input: domain.LoginUserInput{
				Login:    "test",
				Password: "test",
			},
			hasherMock: func(h *mock_hash.MockHasher, in domain.LoginUserInput) {
				h.EXPECT().Hash(in.Password).Return("", errors.New("hasher error"))
			},
			repositoryMock:   func(r *mock_service.MockUserRepositoryI, in domain.LoginUserInput) {},
			tokenManagerMock: func(tm *mock_jwt.MockTokenManagerI, in domain.LoginUserInput) {},
			token:            "",
			err:              errors.New("hasher error"),
		},
		{
			name: "Repository error",
			input: domain.LoginUserInput{
				Login:    "test",
				Password: "test",
			},
			hasherMock: func(h *mock_hash.MockHasher, in domain.LoginUserInput) {
				h.EXPECT().Hash(in.Password).Return(in.Password, nil)
			},
			repositoryMock: func(r *mock_service.MockUserRepositoryI, in domain.LoginUserInput) {
				r.EXPECT().GetByCredentials(context.Background(), in).Return(domain.User{}, errors.New("repository error"))
			},
			tokenManagerMock: func(tm *mock_jwt.MockTokenManagerI, in domain.LoginUserInput) {},
			token:            "",
			err:              errors.New("repository error"),
		},
		{
			name: "TokenManager error",
			input: domain.LoginUserInput{
				Login:    "test",
				Password: "test",
			},
			hasherMock: func(h *mock_hash.MockHasher, in domain.LoginUserInput) {
				h.EXPECT().Hash(in.Password).Return(in.Password, nil)
			},
			repositoryMock: func(r *mock_service.MockUserRepositoryI, in domain.LoginUserInput) {
				r.EXPECT().GetByCredentials(context.Background(), in).Return(domain.User{Id: "userId"}, nil)
			},
			tokenManagerMock: func(tm *mock_jwt.MockTokenManagerI, in domain.LoginUserInput) {
				tm.EXPECT().NewToken("userId").Return("", errors.New("token manager error"))
			},
			token: "",
			err:   errors.New("token manager error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			hasher := mock_hash.NewMockHasher(ctrl)
			testCase.hasherMock(hasher, testCase.input)

			tokenManager := mock_jwt.NewMockTokenManagerI(ctrl)
			testCase.tokenManagerMock(tokenManager, testCase.input)

			userRepository := mock_service.NewMockUserRepositoryI(ctrl)
			testCase.repositoryMock(userRepository, testCase.input)

			auth := NewAuthService(userRepository, hasher, tokenManager)
			token, err := auth.SignIn(context.Background(), testCase.input)

			assert.Equal(t, token, testCase.token)
			assert.Equal(t, err, testCase.err)
		})
	}
}

func TestAuthService_SignUp(t *testing.T) {
	type tokenManagerMockBehaviour func(tm *mock_jwt.MockTokenManagerI, in domain.CreateUserInput)
	type hasherMockBehaviour func(h *mock_hash.MockHasher, in domain.CreateUserInput)
	type repositoryMockBehaviour func(r *mock_service.MockUserRepositoryI, in domain.CreateUserInput)

	testCases := []struct {
		name             string
		input            domain.CreateUserInput
		hasherMock       hasherMockBehaviour
		repositoryMock   repositoryMockBehaviour
		tokenManagerMock tokenManagerMockBehaviour
		token            string
		err              error
	}{
		{
			name: "OK",
			input: domain.CreateUserInput{
				Login:    "test",
				Password: "test",
			},
			hasherMock: func(h *mock_hash.MockHasher, in domain.CreateUserInput) {
				h.EXPECT().Hash(in.Password).Return(in.Password, nil)
			},
			repositoryMock: func(r *mock_service.MockUserRepositoryI, in domain.CreateUserInput) {
				r.EXPECT().Create(context.Background(), in).Return(domain.User{Id: "userId"}, nil)
			},
			tokenManagerMock: func(tm *mock_jwt.MockTokenManagerI, in domain.CreateUserInput) {
				tm.EXPECT().NewToken("userId").Return("token", nil)
			},
			token: "token",
			err:   nil,
		},
		{
			name: "Hasher error",
			input: domain.CreateUserInput{
				Login:    "test",
				Password: "test",
			},
			hasherMock: func(h *mock_hash.MockHasher, in domain.CreateUserInput) {
				h.EXPECT().Hash(in.Password).Return("", errors.New("hasher error"))
			},
			repositoryMock:   func(r *mock_service.MockUserRepositoryI, in domain.CreateUserInput) {},
			tokenManagerMock: func(tm *mock_jwt.MockTokenManagerI, in domain.CreateUserInput) {},
			token:            "",
			err:              errors.New("hasher error"),
		},
		{
			name: "Repository error",
			input: domain.CreateUserInput{
				Login:    "test",
				Password: "test",
			},
			hasherMock: func(h *mock_hash.MockHasher, in domain.CreateUserInput) {
				h.EXPECT().Hash(in.Password).Return(in.Password, nil)
			},
			repositoryMock: func(r *mock_service.MockUserRepositoryI, in domain.CreateUserInput) {
				r.EXPECT().Create(context.Background(), in).Return(domain.User{}, errors.New("repository error"))
			},
			tokenManagerMock: func(tm *mock_jwt.MockTokenManagerI, in domain.CreateUserInput) {},
			token:            "",
			err:              errors.New("repository error"),
		},
		{
			name: "TokenManager error",
			input: domain.CreateUserInput{
				Login:    "test",
				Password: "test",
			},
			hasherMock: func(h *mock_hash.MockHasher, in domain.CreateUserInput) {
				h.EXPECT().Hash(in.Password).Return(in.Password, nil)
			},
			repositoryMock: func(r *mock_service.MockUserRepositoryI, in domain.CreateUserInput) {
				r.EXPECT().Create(context.Background(), in).Return(domain.User{Id: "userId"}, nil)
			},
			tokenManagerMock: func(tm *mock_jwt.MockTokenManagerI, in domain.CreateUserInput) {
				tm.EXPECT().NewToken("userId").Return("", errors.New("token manager error"))
			},
			token: "",
			err:   errors.New("token manager error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			hasher := mock_hash.NewMockHasher(ctrl)
			testCase.hasherMock(hasher, testCase.input)

			tokenManager := mock_jwt.NewMockTokenManagerI(ctrl)
			testCase.tokenManagerMock(tokenManager, testCase.input)

			userRepository := mock_service.NewMockUserRepositoryI(ctrl)
			testCase.repositoryMock(userRepository, testCase.input)

			auth := NewAuthService(userRepository, hasher, tokenManager)
			token, err := auth.SignUp(context.Background(), testCase.input)

			assert.Equal(t, token, testCase.token)
			assert.Equal(t, err, testCase.err)
		})
	}
}

func TestAuthService_CheckToken(t *testing.T) {
	type tokenManagerMockBehaviour func(tm *mock_jwt.MockTokenManagerI, token string)
	testCases := []struct {
		name             string
		token            string
		tokenManagerMock tokenManagerMockBehaviour
		userId           string
		err              error
	}{
		{
			name:             "Valid token",
			token:            "token",
			tokenManagerMock: func(tm *mock_jwt.MockTokenManagerI, token string) {
				tm.EXPECT().Parse(token).Return("userId", nil)
			},
			userId:           "userId",
			err:              nil,
		},
		{
			name:             "Invalid token",
			token:            "token",
			tokenManagerMock: func(tm *mock_jwt.MockTokenManagerI, token string) {
				tm.EXPECT().Parse(token).Return("", errors.New("invalid token"))
			},
			userId:           "",
			err:              errors.New("invalid token"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			hasher := &hash.SHA1Hasher{}
			userRepository := &mongorep.UserRepository{}

			tokenManager := mock_jwt.NewMockTokenManagerI(ctrl)
			testCase.tokenManagerMock(tokenManager, testCase.token)

			auth := NewAuthService(userRepository, hasher, tokenManager)
			userId, err := auth.CheckToken(context.Background(), testCase.token)

			assert.Equal(t, userId, testCase.userId)
			assert.Equal(t, err, testCase.err)
		})
	}
}
