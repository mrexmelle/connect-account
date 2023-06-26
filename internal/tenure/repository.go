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

func (r *Repository) Create(req Entity) (Entity, error) {
	sd, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return Entity{}, err
	}

	var res *gorm.DB
	if req.EndDate == "" {
		res = r.Config.Db.Raw(
			"INSERT INTO "+r.TableName+"(ehid, start_date, employment_type, organization_id, "+
				"title_grade, title_name, created_at, updated_at) "+
				"VALUES(?, ?, ?, ?, ?, ?, NOW(), NOW()) RETURNING id",
			req.Ehid,
			datatypes.Date(sd),
			req.EmploymentType,
			req.OrganizationId,
			req.TitleGrade,
			req.TitleName,
		).Scan(&req.Id)
	} else {
		ed, err := time.Parse("2006-01-02", req.EndDate)
		if err != nil {
			return Entity{}, err
		}
		res = r.Config.Db.Raw(
			"INSERT INTO "+r.TableName+"(ehid, start_date, end_date, employment_type, organization_id, "+
				"title_grade, title_name, created_at, updated_at) "+
				"VALUES(?, ?, ?, ?, ?, ?, ?, NOW(), NOW()) RETURNING id",
			req.Ehid,
			datatypes.Date(sd),
			datatypes.Date(ed),
			req.EmploymentType,
			req.OrganizationId,
			req.TitleGrade,
			req.TitleName,
		).Scan(&req.Id)
	}

	if res.Error != nil {
		return Entity{}, res.Error
	}

	return req, nil
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
		return errors.New("request invalid")
	}

	return nil
}

func (r *Repository) FindByEhid(ehid string) ([]Entity, error) {
	result := r.Config.Db.
		Select("id, start_date, end_date, employment_type, organization_id, title_grade, title_name").
		Table(r.TableName).
		Where("ehid = ?", ehid).
		Order("start_date DESC")
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
			&t.TitleGrade,
			&t.TitleName,
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
		Select("id, start_date, end_date, employment_type, organization_id, title_grade, title_name").
		Table(r.TableName).
		Where("ehid = ? AND start_date < NOW() AND (end_date IS NULL OR end_date > NOW())", ehid)
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
			&t.TitleGrade,
			&t.TitleName,
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
