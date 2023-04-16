package session

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Authenticate(id string, password string) bool {
	dsn := "host=127.0.0.1 user=iam password=123 dbname=iam port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return false
	}

	var idResult string

	err = db.
		Select("id").
		Table("credentials").
		Where("id = ? AND password_hash = SHA256(?)", id, password).
		Row().
		Scan(&idResult)

	return (err == nil) && (idResult == id)
}

func GenerateJwt(id string) string {

}
