package tenure

import (
	"database/sql"
	"errors"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func Create(req TenureCreateRequest, db *gorm.DB) error {
	ts, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return err
	}

	res := db.Exec(
		"INSERT INTO tenures(ehid, employee_id, start_date, employment_type, "+
			"created_at, updated_at) "+
			"VALUES(?, ?, ?, ?, NOW(), NOW())",
		req.Ehid,
		req.EmployeeId,
		datatypes.Date(ts),
		req.EmploymentType,
	)

	return res.Error
}

func UpdateEndDate(req TenureUpdateEndDateRequest, db *gorm.DB) error {
	ts, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return nil
	}
	result := db.
		Table("tenures").
		Where("id = ? AND ehid = ?", req.Id, req.Ehid).
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

func Retrieve(ehid string, db *gorm.DB) (TenureRetrieveResponse, error) {
	result := db.
		Select("id, employee_id, start_date, end_date, employment_type").
		Table("tenures").
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
			&t.EmployeeId,
			&startDate,
			&endDate,
			&t.EmploymentType,
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
