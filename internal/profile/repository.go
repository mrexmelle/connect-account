package profile

import (
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

func (r *Repository) CreateWithDb(db *gorm.DB, req ProfileCreateRequest) error {
	ts, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return err
	}

	res := db.Exec(
		"INSERT INTO "+r.TableName+"(ehid, employee_id, name, dob, "+
			"created_at, updated_at) "+
			"VALUES(?, ?, ?, ?, NOW(), NOW())",
		req.Ehid,
		req.EmployeeId,
		req.Name,
		datatypes.Date(ts),
	)
	return res.Error
}

func (r *Repository) FindByEhid(ehid string) (ProfileRetrieveResponse, error) {
	response := ProfileRetrieveResponse{
		Ehid: ehid,
	}
	var dob time.Time
	err := r.Config.Db.
		Select("name, employee_id, dob").
		Table(r.TableName).
		Where("ehid = ?", ehid).
		Row().
		Scan(&response.Name, &response.EmployeeId, &dob)
	if err != nil {
		return ProfileRetrieveResponse{}, err
	}

	response.Dob = dob.Format("2006-01-02")
	return response, nil
}
