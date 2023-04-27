package profile

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func Create(req ProfileCreateRequest, db *gorm.DB) error {
	ts, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return err
	}

	res := db.Exec(
		"INSERT INTO profiles(ehid, name, dob, "+
			"created_at, updated_at) "+
			"VALUES(?, ?, ?, NOW(), NOW())",
		req.Ehid,
		req.Name,
		datatypes.Date(ts),
	)

	return res.Error
}

func Retrieve(ehid string, db *gorm.DB) (ProfileRetrieveResponse, error) {
	var res ProfileRetrieveResponse
	var dob time.Time
	err := db.
		Select("name, dob").
		Table("profiles").
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
