package builder

import (
	"github.com/i-vasilkov/go-todo-app/internal/handler/http"
	"github.com/i-vasilkov/go-todo-app/internal/service"
)

type AppServiceBuilder struct{
	deps *service.Dependencies
	reps *service.Repositories
}

func NewAppServiceBuilder(deps *service.Dependencies, reps *service.Repositories) *AppServiceBuilder {
	return &AppServiceBuilder{
		deps: deps,
		reps: reps,
	}
}

func (b *AppServiceBuilder) Build() *http.Services {
	return &http.Services{
		Auth: service.NewAuthService(b.reps.User, b.deps.Hasher, b.deps.JwtManager),
		Task: service.NewTaskService(b.reps.Task),
	}
}
