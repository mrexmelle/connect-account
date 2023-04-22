package session

import (
	"errors"
	"time"

	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Authenticate(req SessionPostRequest) (bool, error) {
	dsn := "host=127.0.0.1 user=iam password=123 dbname=iam port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return false, err
	}

	var idResult string

	err = db.
		Select("employee_id").
		Table("credentials").
		Where("employee_id = ? AND password_hash = CRYPT(?, password_hash)", req.EmployeeId, req.Password).
		Row().
		Scan(&idResult)

	return (idResult == req.EmployeeId), err
}

func GenerateJwt(employeeId string) (string, error) {
	secret := "1nt3rst3ll4r-*-a5tR0"

	signingKey := jose.SigningKey{
		Algorithm: jose.HS256,
		Key:       []byte(secret),
	}

	sig, err := jose.NewSigner(
		signingKey,
		(&jose.SignerOptions{}).WithType("JWT"),
	)

	if err != nil {
		return "", err
	}

	now := time.Now()

	claim := jwt.Claims{
		Subject:   employeeId,
		Issuer:    "connect-iam",
		NotBefore: jwt.NewNumericDate(now),
		Expiry:    jwt.NewNumericDate(now.Add(time.Minute * 3)),
		Audience:  jwt.Audience{"connect-iam"},
	}

	raw, err := jwt.Signed(sig).Claims(claim).CompactSerialize()

	if err != nil {
		return "", errors.New("Claim cannot be signed")
	}

	return raw, nil
}
