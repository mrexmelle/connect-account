package account

import (
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/credential"
	"github.com/mrexmelle/connect-idp/internal/mapper"
	"github.com/mrexmelle/connect-idp/internal/profile"
)

type Service struct {
	Config               *config.Config
	CredentialRepository *credential.Repository
	ProfileRepository    *profile.Repository
}

func NewService(
	cfg *config.Config,
	cr *credential.Repository,
	pr *profile.Repository,
) *Service {
	return &Service{
		Config:               cfg,
		CredentialRepository: cr,
		ProfileRepository:    pr,
	}
}

func (s *Service) Register(req AccountPostRequest) error {
	trx := s.Config.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()
	if trx.Error != nil {
		return trx.Error
	}

	err := s.CredentialRepository.CreateWithDb(
		trx,
		req.EmployeeId,
		req.Password,
	)
	if err != nil {
		trx.Rollback()
		return err
	}

	_, err = s.ProfileRepository.CreateWithDb(
		trx,
		profile.Entity{
			Ehid:         mapper.ToEhid(req.EmployeeId),
			EmployeeId:   req.EmployeeId,
			EmailAddress: req.EmailAddress,
			Name:         req.Name,
			Dob:          req.Dob,
		},
	)
	if err != nil {
		trx.Rollback()
		return err
	}

	return trx.Commit().Error
}
