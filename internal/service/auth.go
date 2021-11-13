package service

import (
	"context"
	"github.com/i-vasilkov/go-todo-app/internal/domain"
	jwtauth "github.com/i-vasilkov/go-todo-app/pkg/auth/jwt"
	"github.com/i-vasilkov/go-todo-app/pkg/hash"
)

type AuthService struct {
	rep    UserRepositoryI
	hasher hash.Hasher
	jwt    jwtauth.TokenManagerI
}

func NewAuthService(rep UserRepositoryI, hasher hash.Hasher, jwt jwtauth.TokenManagerI) *AuthService {
	return &AuthService{rep: rep, hasher: hasher, jwt: jwt}
}

func (as *AuthService) SignUp(ctx context.Context, in domain.CreateUserInput) (string, error) {
	var err error
	var user domain.User

	in.Password, err = as.hasher.Hash(in.Password)
	if err != nil {
		return "", err
	}

	user, err = as.rep.Create(ctx, in)
	if err != nil {
		return "", err
	}

	return as.jwt.NewToken(user.Id)
}

func (as *AuthService) SignIn(ctx context.Context, in domain.LoginUserInput) (string, error) {
	var err error

	in.Password, err = as.hasher.Hash(in.Password)
	if err != nil {
		return "", err
	}

	user, err := as.rep.GetByCredentials(ctx, in)
	if err != nil {
		return "", err
	}

	return as.jwt.NewToken(user.Id)
}

func (as *AuthService) CheckToken(ctx context.Context, token string) (string, error) {
	return as.jwt.Parse(token)
}
