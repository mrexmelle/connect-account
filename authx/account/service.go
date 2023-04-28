package account

import (
	"crypto/sha256"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/mrexmelle/connect-iam/authx/credential"
	"github.com/mrexmelle/connect-iam/authx/profile"
	"github.com/mrexmelle/connect-iam/authx/tenure"
	"gorm.io/gorm"
)

func Register(req AccountPostRequest, db *gorm.DB) error {
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

	err := credential.Create(cred, trx)
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

func UpdateEndDate(tenureId int, ehid string, req AccountPatchRequest, db *gorm.DB) error {
	data := tenure.TenureUpdateEndDateRequest{
		Id:      tenureId,
		Ehid:    ehid,
		EndDate: req.Value,
	}

	return tenure.UpdateEndDate(data, db)
}

func RetrieveProfile(ehid string, db *gorm.DB) (AccountGetProfileResponse, error) {
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

func RetrieveTenures(ehid string, db *gorm.DB) (AccountGetTenureResponse, error) {
	result, err := tenure.Retrieve(ehid, db)

	if err != nil {
		return AccountGetTenureResponse{}, err
	}

	data := AccountGetTenureResponse{Ehid: ehid}
	copier.Copy(&data.Tenures, &result.Tenures)

	return data, nil
}
