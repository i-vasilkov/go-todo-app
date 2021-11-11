package hash

type Hasher interface {
	Hash(string) (string, error)
}