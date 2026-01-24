package security

type IHash interface {
	HashPassword(in string) (string, error)
	CheckPasswordHash(in, inHash string) bool
}
