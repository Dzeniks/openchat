package secretKey

var secretKey []byte

func SetSecretKey(secretKeyString string) {
	secretKey = []byte(secretKeyString)
}

func GetSecretKey() []byte {
	return secretKey
}
