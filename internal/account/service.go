package account

import (
	"github.com/jinzhu/copier"
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/credential"
	"github.com/mrexmelle/connect-idp/internal/ehid"
	"github.com/mrexmelle/connect-idp/internal/profile"
	"github.com/mrexmelle/connect-idp/internal/tenure"
)

type Service struct {
	Config               *config.Config
	CredentialRepository *credential.Repository
	ProfileRepository    *profile.Repository
	TenureRepository     *tenure.Repository
}

func NewService(
	cfg *config.Config,
	cr *credential.Repository,
	pr *profile.Repository,
	tr *tenure.Repository) *Service {
	return &Service{
		Config:               cfg,
		CredentialRepository: cr,
		ProfileRepository:    pr,
		TenureRepository:     tr,
	}
}

func (s *Service) Register(req AccountPostRequest) error {
	cred, bp := Disperse(req)

	trx := s.Config.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()

	if trx.Error != nil {
		return trx.Error
	}

	err := s.CredentialRepository.CreateWithDb(trx, cred)
	if err != nil {
		trx.Rollback()
		return err
	}
	err = s.ProfileRepository.CreateWithDb(trx, bp)
	if err != nil {
		trx.Rollback()
		return err
	}

	return trx.Commit().Error
}

func Disperse(req AccountPostRequest) (
	credential.CredentialCreateRequest,
	profile.ProfileCreateRequest,
) {
	ehid := ehid.FromEmployeeId(req.EmployeeId)

	cred := credential.CredentialCreateRequest{
		EmployeeId: req.EmployeeId,
		Password:   req.Password,
	}

	bp := profile.ProfileCreateRequest{
		Ehid:       ehid,
		EmployeeId: req.EmployeeId,
		Name:       req.Name,
		Dob:        req.Dob,
	}

	return cred, bp
}

func (s *Service) UpdateEndDate(
	ehid string,
	tenureId int,
	req AccountPatchRequest,
) error {
	return s.TenureRepository.UpdateEndDateByIdAndEhid(
		req.Value,
		tenureId,
		ehid,
	)
}

func (s *Service) RetrieveProfile(
	ehid string,
) (AccountGetProfileResponse, error) {
	result, err := s.ProfileRepository.FindByEhid(ehid)

	if err != nil {
		return AccountGetProfileResponse{}, err
	}

	return AccountGetProfileResponse{
		Ehid:       ehid,
		EmployeeId: result.EmployeeId,
		Name:       result.Name,
		Dob:        result.Dob,
	}, nil
}

func (s *Service) RetrieveTenures(
	ehid string,
) (AccountGetTenureResponse, error) {
	result, err := s.TenureRepository.FindByEhid(ehid)

	if err != nil {
		return AccountGetTenureResponse{}, err
	}

	data := AccountGetTenureResponse{Ehid: ehid}
	copier.Copy(&data.Tenures, &result.Tenures)

	return data, nil
}

func (s *Service) PostTenure(
	ehid string,
	request AccountPostTenureRequest,
) error {
	data := tenure.TenureCreateRequest{
		Ehid:           ehid,
		StartDate:      request.StartDate,
		EndDate:        request.EndDate,
		EmploymentType: request.EmploymentType,
	}
	return s.TenureRepository.Create(data)
}
