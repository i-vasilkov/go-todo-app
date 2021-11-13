package hash

//go:generate mockgen -source=interface.go -destination=mocks/mock.go

type Hasher interface {
	Hash(string) (string, error)
}