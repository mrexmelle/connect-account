package bprof

import (
	"errors"
	"time"

	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Create(bp BasicProfilePostRequest) (string, error) {
	dsn := "host=127.0.0.1 user=iam password=123 dbname=iam port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return "", err
	}

	var idResult string

	res := db.Exec(
		"INSERT INTO basic_profiles(employee_id_hash, name, dob, \
			created_at, updated_at) \
			VALUES(ENCODE(SHA256(?), 'hex'), ?, ?, NOW(), NOW())",
		bp.EmployeeId,
		bp.Name,
		datatypes.Date(time.Parse("YYYY-MM-DD", bp.Dob)),
	)

	if res.Error != nil {

	}

	return (idResult == id), err
}

func GenerateJwt(id string) (string, error) {
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
		Subject:   id,
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
