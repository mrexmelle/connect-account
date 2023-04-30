package profile

import (
	"time"

	"github.com/mrexmelle/connect-iam/idp/config"
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
		"INSERT INTO "+r.TableName+"(ehid, name, dob, "+
			"created_at, updated_at) "+
			"VALUES(?, ?, ?, NOW(), NOW())",
		req.Ehid,
		req.Name,
		datatypes.Date(ts),
	)

	return res.Error
}

func (r *Repository) FindByEhid(ehid string) (ProfileRetrieveResponse, error) {
	var res ProfileRetrieveResponse
	var dob time.Time
	err := r.Config.Db.
		Select("name, dob").
		Table(r.TableName).
		Where("ehid = ?", ehid).
		Row().
		Scan(&res.Name, &dob)
	res.Dob = dob.Format("2006-01-02")

	if err != nil {
		return ProfileRetrieveResponse{}, err
	}

	res.Ehid = ehid
	return res, nil
}
