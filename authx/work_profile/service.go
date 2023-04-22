package wprof

import (
	"errors"
	"time"

	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Create(wp WorkProfilePostRequest) (string, error) {
	dsn := "host=127.0.0.1 user=iam password=123 dbname=iam port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return "", err
	}

	var idResult string

	res := db.Exec(
		"INSERT INTO work_profiles(employee_id_hash, employee_id, work_email_address, start_date, employment_type, \
			created_at, updated_at) \
			VALUES(ENCODE(SHA256(?), 'hex'), ?, ?, ?, ?, NOW(), NOW())",
		wp.EmployeeId,
		wp.EmployeeId,
		wp.WorkEmailAddress,
		datatypes.Date(time.Parse("YYYY-MM-DD", wp.StartDate)),
		wp.EmploymentType
	)

	if res.Error != nil {

	}

	return (idResult == id), err
}
