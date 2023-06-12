package profile

import (
	"errors"
	"time"

	"github.com/mrexmelle/connect-idp/internal/config"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Repository struct {
	Config    *config.Config
	TableName string
}

func NewRepository(cfg *config.Config) *Repository {
	return &Repository{
		Config:    cfg,
		TableName: "profiles",
	}
}

func (r *Repository) CreateWithDb(db *gorm.DB, req Entity) (Entity, error) {
	ts, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return Entity{}, err
	}

	res := db.Exec(
		"INSERT INTO "+r.TableName+"(ehid, employee_id, name, email_address, dob, "+
			"created_at, updated_at) "+
			"VALUES(?, ?, ?, ?, ?, NOW(), NOW())",
		req.Ehid,
		req.EmployeeId,
		req.Name,
		req.EmailAddress,
		datatypes.Date(ts),
	)
	if res.Error != nil {
		return Entity{}, res.Error
	}

	return Entity{
		Ehid:         req.Ehid,
		EmployeeId:   req.EmployeeId,
		Name:         req.Name,
		EmailAddress: req.EmailAddress,
		Dob:          req.Dob,
	}, nil
}

func (r *Repository) FindByEhid(ehid string) (Entity, error) {
	response := Entity{
		Ehid: ehid,
	}
	var dob time.Time
	err := r.Config.Db.
		Select("employee_id, name, email_address, dob").
		Table(r.TableName).
		Where("ehid = ?", ehid).
		Row().
		Scan(&response.EmployeeId, &response.Name, &response.EmailAddress, &dob)
	if err != nil {
		return Entity{}, err
	}

	response.Dob = dob.Format("2006-01-02")
	return response, nil
}

func (r *Repository) UpdateEmailByEhid(email string, ehid string) error {
	result := r.Config.Db.
		Table(r.TableName).
		Where("ehid = ?", ehid).
		Updates(
			map[string]interface{}{
				"email_address": email,
				"updated_at":    time.Now(),
			},
		)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("request invalid")
	}

	return nil
}
