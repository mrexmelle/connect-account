package session

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/mrexmelle/connect-iam/authx/credential"
	"gorm.io/gorm"
)

func Authenticate(req SessionPostRequest, db *gorm.DB) (bool, error) {
	cred := credential.CredentialAuthRequest{
		req.EmployeeId,
		req.Password,
	}
	return credential.Authenticate(cred, db)
}

func GenerateJwt(employeeId string, tokenAuth *jwtauth.JWTAuth) (string, error) {
	now := time.Now()
	_, token, err := tokenAuth.Encode(
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
