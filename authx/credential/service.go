package credential

import (
	"time"

	"gorm.io/gorm"
)

func Create(req CredentialCreateRequest, db *gorm.DB) error {
	res := db.Exec(
		"INSERT INTO credentials(employee_id, password_hash, "+
			"created_at, updated_at) "+
			"VALUES(?, CRYPT(?, GEN_SALT('bf', 8)), NOW(), NOW())",
		req.EmployeeId,
		req.Password,
	)
	return res.Error
}

func Authenticate(req CredentialAuthRequest, db *gorm.DB) (bool, error) {
	var idResult string
	err := db.
		Select("employee_id").
		Table("credentials").
		Where(
			"employee_id = ? AND password_hash = CRYPT(?, password_hash) AND deleted_at IS NULL",
			req.EmployeeId,
			req.Password,
		).
		Row().
		Scan(&idResult)
	return (idResult == req.EmployeeId), err
}

func Delete(req CredentialDeleteRequest, db *gorm.DB) error {
	now := time.Now()
	result := db.
		Table("credentials").
		Where("employee_id = ? AND deleted_at IS NOT NULL", req.EmployeeId).
		Updates(
			map[string]interface{}{
				"deleted_at": now,
				"updated_at": now,
			},
		)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
