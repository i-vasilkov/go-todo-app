package jwt

//go:generate mockgen -source=interface.go -destination=mocks/mock.go

type TokenManagerI interface {
	NewToken(id string) (string, error)
	Parse(token string) (string, error)
}
