package account

import (
	"crypto/sha256"
	"fmt"

	"github.com/mrexmelle/connect-iam/authx/credential"
	"github.com/mrexmelle/connect-iam/authx/profile"
	"github.com/mrexmelle/connect-iam/authx/tenure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Register(req AccountPostRequest) error {
	dsn := "host=127.0.0.1 user=iam password=123 dbname=iam port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	cred, bp, emp := Disperse(req)

	trx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()

	if trx.Error != nil {
		return trx.Error
	}

	err = credential.Create(cred, trx)
	if err != nil {
		trx.Rollback()
		return err
	}
	err = profile.Create(bp, trx)
	if err != nil {
		trx.Rollback()
		return err
	}
	err = tenure.Create(emp, trx)
	if err != nil {
		trx.Rollback()
		return err
	}

	return trx.Commit().Error
}

func Disperse(req AccountPostRequest) (
	credential.CredentialCreateRequest,
	profile.ProfileCreateRequest,
	tenure.TenureCreateRequest,
) {
	ehid := GenerateEhid(req.EmployeeId)

	cred := credential.CredentialCreateRequest{
		req.EmployeeId,
		req.Password,
	}

	bp := profile.ProfileCreateRequest{
		ehid,
		req.Name,
		req.Dob,
	}
	emp := tenure.TenureCreateRequest{
		ehid,
		req.EmployeeId,
		req.StartDate,
		req.EmploymentType,
	}

	return cred, bp, emp
}

func GenerateEhid(employeeId string) string {
	hasher := sha256.New()
	hasher.Write([]byte(employeeId))

	return fmt.Sprintf("u%x", hasher.Sum(nil))
}

func UpdateEndDate(tenureId int, ehid string, req AccountPatchRequest) error {
	dsn := "host=127.0.0.1 user=iam password=123 dbname=iam port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	data := tenure.TenureUpdateEndDateRequest{
		Id:      tenureId,
		Ehid:    ehid,
		EndDate: req.Value,
	}

	return tenure.UpdateEndDate(data, db)
}

func RetrieveProfile(ehid string) (AccountGetProfileResponse, error) {
	dsn := "host=127.0.0.1 user=iam password=123 dbname=iam port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return AccountGetProfileResponse{}, err
	}

	result, err := profile.Retrieve(ehid, db)

	if err != nil {
		return AccountGetProfileResponse{}, err
	}

	data := AccountGetProfileResponse{
		Ehid: ehid,
		Name: result.Name,
		Dob:  result.Dob,
	}

	return data, nil
}

func RetrieveTenures(ehid string) (AccountGetTenureResponse, error) {
	dsn := "host=127.0.0.1 user=iam password=123 dbname=iam port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return AccountGetTenureResponse{}, err
	}

	result, err := tenure.Retrieve(ehid, db)

	if err != nil {
		return AccountGetTenureResponse{}, err
	}

	data := AccountGetTenureResponse{
		Ehid:    ehid,
		Tenures: result.Tenures,
	}

	return data, nil
}
