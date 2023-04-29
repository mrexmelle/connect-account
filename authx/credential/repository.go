package credential

import (
	"time"

	"github.com/mrexmelle/connect-iam/authx/config"
	"gorm.io/gorm"
)

type Repository struct {
	Config    *config.Config
	TableName string
}

func NewRepository(cfg *config.Config) *Repository {
	return &Repository{
		Config:    cfg,
		TableName: "credentials",
	}
}

func (r *Repository) CreateWithDb(
	db *gorm.DB,
	req CredentialCreateRequest,
) error {
	res := db.Exec(
		"INSERT INTO ?(employee_id, password_hash, "+
			"created_at, updated_at) "+
			"VALUES(?, CRYPT(?, GEN_SALT('bf', 8)), NOW(), NOW())",
		r.TableName,
		req.EmployeeId,
		req.Password,
	)
	return res.Error
}

func (r *Repository) ExistsByEmployeeIdAndPassword(
	employeeId string,
	password string,
) (bool, error) {
	var idResult string
	err := r.Config.Db.
		Select("employee_id").
		Table(r.TableName).
		Where(
			"employee_id = ? AND password_hash = CRYPT(?, password_hash) "+
				"AND deleted_at IS NULL",
			employeeId,
			password,
		).
		Row().
		Scan(&idResult)
	return (idResult == employeeId), err
}

func (r *Repository) DeleteByEmployeeId(employeeId string) error {
	now := time.Now()
	result := r.Config.Db.
		Table(r.TableName).
		Where("employee_id = ? AND deleted_at IS NOT NULL", employeeId).
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
