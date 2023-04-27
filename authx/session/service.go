package session

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/mrexmelle/connect-iam/authx/credential"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Authenticate(req SessionPostRequest) (bool, error) {
	dsn := "host=127.0.0.1 user=iam password=123 dbname=iam port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return false, err
	}

	cred := credential.CredentialAuthRequest{
		req.EmployeeId,
		req.Password,
	}
	return credential.Authenticate(cred, db)
}

func GenerateJwt(employeeId string) (string, error) {
	authenticator := jwtauth.New("HS256", []byte("1nt3rst3ll4r-*-a5tR0"), nil)

	now := time.Now()
	_, token, err := authenticator.Encode(
		map[string]interface{}{
			"aud": "connect-iam",
			"exp": now.Add(time.Hour * 3).Unix(),
			"iat": now.Unix(),
			"iss": "connect-iam",
			"nbf": now.Unix(),
			"sub": GenerateEhid(employeeId),
		},
	)

	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateEhid(employeeId string) string {
	hasher := sha256.New()
	hasher.Write([]byte(employeeId))

	return fmt.Sprintf("u%x", hasher.Sum(nil))
}
