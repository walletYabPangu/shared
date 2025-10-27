package shared

type JWT struct {
	Key string
}

type IJWT interface {
	GenerateToken(userID uint)
}

func JwtNew(Key string) {

}

func (j *JWT) GenerateToken(userID uint) {

}
