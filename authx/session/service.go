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

func Authenticate(req SessionPostRequest, dsn string) (bool, error) {
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

func GenerateJwt(employeeId string, jwtSecret string) (string, error) {
	authenticator := jwtauth.New("HS256", []byte(jwtSecret), nil)

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
