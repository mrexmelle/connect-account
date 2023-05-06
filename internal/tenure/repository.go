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

func (r *Repository) CreateWithDb(db *gorm.DB, req Entity) (Entity, error) {
	sd, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return Entity{}, err
	}

	var res *gorm.DB
	var id int64
	if req.EndDate == "" {
		res = db.Exec(
			"INSERT INTO "+r.TableName+"(ehid, start_date, employment_type, organization_id, "+
				"created_at, updated_at) "+
				"VALUES(?, ?, ?, ?, NOW(), NOW()) RETURNING id",
			req.Ehid,
			datatypes.Date(sd),
			req.EmploymentType,
			req.OrganizationId,
		).Scan(&id)
	} else {
		ed, err := time.Parse("2006-01-02", req.EndDate)
		if err != nil {
			return Entity{}, err
		}
		res = db.Exec(
			"INSERT INTO "+r.TableName+"(ehid, start_date, end_date, employment_type, organization_id, "+
				"created_at, updated_at) "+
				"VALUES(?, ?, ?, ?, ?, NOW(), NOW()) RETURNING id",
			req.Ehid,
			datatypes.Date(sd),
			datatypes.Date(ed),
			req.EmploymentType,
			req.OrganizationId,
		).Scan(&id)
	}

	if res.Error != nil {
		return Entity{}, res.Error
	}

	return Entity{
		Id:             id,
		Ehid:           req.Ehid,
		StartDate:      req.StartDate,
		EndDate:        req.EndDate,
		EmploymentType: req.EmploymentType,
		OrganizationId: req.OrganizationId,
	}, nil
}

func (r *Repository) Create(req Entity) (Entity, error) {
	return r.CreateWithDb(r.Config.Db, req)
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

func (r *Repository) FindByEhid(ehid string) ([]Entity, error) {
	result := r.Config.Db.
		Select("id, start_date, end_date, employment_type, organization_id").
		Table(r.TableName).
		Where("ehid = ?", ehid)
	var response = make([]Entity, result.RowsAffected)

	rows, err := result.Rows()
	if err != nil {
		return response, err
	}
	defer rows.Close()

	for rows.Next() {
		t := Entity{Ehid: ehid}
		var startDate time.Time
		var endDate sql.NullTime
		rows.Scan(
			&t.Id,
			&startDate,
			&endDate,
			&t.EmploymentType,
			&t.OrganizationId,
		)
		t.StartDate = startDate.Format("2006-01-02")
		if endDate.Valid {
			t.EndDate = endDate.Time.Format("2006-01-02")
		} else {
			t.EndDate = ""
		}
		response = append(response, t)
	}

	return response, result.Error
}

func (r *Repository) FindCurrentTenureByEhid(ehid string) ([]Entity, error) {
	result := r.Config.Db.
		Select("id, start_date, end_date, employment_type, organization_id").
		Table(r.TableName).
		Where("ehid = ? AND (end_date IS NULL OR end_date > NOW())", ehid)
	var response = make([]Entity, result.RowsAffected)

	rows, err := result.Rows()
	if err != nil {
		return response, err
	}
	defer rows.Close()

	for rows.Next() {
		t := Entity{Ehid: ehid}
		var startDate time.Time
		var endDate sql.NullTime
		rows.Scan(
			&t.Id,
			&startDate,
			&endDate,
			&t.EmploymentType,
			&t.OrganizationId,
		)
		t.StartDate = startDate.Format("2006-01-02")
		if endDate.Valid {
			t.EndDate = endDate.Time.Format("2006-01-02")
		} else {
			t.EndDate = ""
		}
		response = append(response, t)
	}

	return response, result.Error
}
