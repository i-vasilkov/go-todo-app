package service

type AppServiceBuilder struct{
	deps *Dependencies
	reps *Repositories
}

func NewAppServiceBuilder(deps *Dependencies, reps *Repositories) *AppServiceBuilder {
	return &AppServiceBuilder{
		deps: deps,
		reps: reps,
	}
}

func (b *AppServiceBuilder) Build() *Services {
	return &Services{
		Auth: NewAuthService(b.reps.User, b.deps.Hasher, b.deps.JwtManager),
		Task: NewTaskService(b.reps.Task),
	}
}