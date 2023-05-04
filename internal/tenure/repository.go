package tenure

import (
	"database/sql"
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
		TableName: "tenures",
	}
}

func (r *Repository) CreateWithDb(db *gorm.DB, req TenureCreateRequest) error {
	sd, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return err
	}

	ed, edErr := time.Parse("2006-01-02", req.EndDate)

	var res *gorm.DB
	if edErr == nil {
		res = db.Exec(
			"INSERT INTO "+r.TableName+"(ehid, start_date, end_date, employment_type, ohid, "+
				"created_at, updated_at) "+
				"VALUES(?, ?, ?, ?, ?, NOW(), NOW())",
			req.Ehid,
			datatypes.Date(sd),
			datatypes.Date(ed),
			req.EmploymentType,
			req.Ohid,
		)
	} else {
		res = db.Exec(
			"INSERT INTO "+r.TableName+"(ehid, start_date, employment_type, ohid, "+
				"created_at, updated_at) "+
				"VALUES(?, ?, ?, ?, NOW(), NOW())",
			req.Ehid,
			datatypes.Date(sd),
			req.EmploymentType,
			req.Ohid,
		)
	}

	return res.Error
}

func (r *Repository) Create(req TenureCreateRequest) error {
	sd, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return err
	}

	ed, edErr := time.Parse("2006-01-02", req.EndDate)

	var res *gorm.DB
	if edErr == nil {
		res = r.Config.Db.Exec(
			"INSERT INTO "+r.TableName+"(ehid, start_date, end_date, employment_type, ohid, "+
				"created_at, updated_at) "+
				"VALUES(?, ?, ?, ?, ?, NOW(), NOW())",
			req.Ehid,
			datatypes.Date(sd),
			datatypes.Date(ed),
			req.EmploymentType,
			req.Ohid,
		)
	} else {
		res = r.Config.Db.Exec(
			"INSERT INTO "+r.TableName+"(ehid, start_date, employment_type, ohid, "+
				"created_at, updated_at) "+
				"VALUES(?, ?, ?, ?, NOW(), NOW())",
			req.Ehid,
			datatypes.Date(sd),
			req.EmploymentType,
			req.Ohid,
		)
	}

	return res.Error
}

func (r *Repository) UpdateEndDateByIdAndEhid(
	endDate string,
	id int,
	ehid string,
) error {
	ts, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil
	}
	result := r.Config.Db.
		Table(r.TableName).
		Where("id = ? AND ehid = ?", id, ehid).
		Updates(
			map[string]interface{}{
				"end_date":   datatypes.Date(ts),
				"updated_at": time.Now(),
			},
		)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Request Invalid")
	}

	return nil
}

func (r *Repository) FindByEhid(ehid string) (TenureRetrieveResponse, error) {
	result := r.Config.Db.
		Select("id, start_date, end_date, employment_type, ohid").
		Table(r.TableName).
		Where("ehid = ?", ehid)
	var response = TenureRetrieveResponse{
		Ehid:    ehid,
		Tenures: make([]Tenure, result.RowsAffected),
	}

	rows, err := result.Rows()
	if err != nil {
		return response, err
	}
	defer rows.Close()

	for rows.Next() {
		var t Tenure
		var startDate time.Time
		var endDate sql.NullTime
		rows.Scan(
			&t.Id,
			&startDate,
			&endDate,
			&t.EmploymentType,
			&t.Ohid,
		)
		t.StartDate = startDate.Format("2006-01-02")
		if endDate.Valid {
			t.EndDate = endDate.Time.Format("2006-01-02")
		} else {
			t.EndDate = ""
		}
		response.Tenures = append(response.Tenures, t)
	}

	return response, result.Error
}

func (r *Repository) FindCurrentTenureByEhid(ehid string) (TenureRetrieveResponse, error) {
	result := r.Config.Db.
		Select("id, start_date, end_date, employment_type, ohid").
		Table(r.TableName).
		Where("ehid = ? AND (end_date IS NULL OR end_date > NOW())", ehid)
	var response = TenureRetrieveResponse{
		Ehid:    ehid,
		Tenures: make([]Tenure, result.RowsAffected),
	}

	rows, err := result.Rows()
	if err != nil {
		return response, err
	}
	defer rows.Close()

	for rows.Next() {
		var t Tenure
		var startDate time.Time
		var endDate sql.NullTime
		rows.Scan(
			&t.Id,
			&startDate,
			&endDate,
			&t.EmploymentType,
			&t.Ohid,
		)
		t.StartDate = startDate.Format("2006-01-02")
		if endDate.Valid {
			t.EndDate = endDate.Time.Format("2006-01-02")
		} else {
			t.EndDate = ""
		}
		response.Tenures = append(response.Tenures, t)
	}

	return response, result.Error
}
