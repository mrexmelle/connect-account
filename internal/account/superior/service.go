package accountSuperior

import (
	accountOrganization "github.com/mrexmelle/connect-idp/internal/account/organization"
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/mapper"
	"github.com/mrexmelle/connect-idp/internal/superior"
)

type Service struct {
	Config                        *config.Config
	AccountOrganizationRepository *accountOrganization.Repository
	SuperiorRepository            *superior.Repository
}

func NewService(
	cfg *config.Config,
	aor *accountOrganization.Repository,
	sr *superior.Repository,
) *Service {
	return &Service{
		Config:                        cfg,
		AccountOrganizationRepository: aor,
		SuperiorRepository:            sr,
	}
}

func (s *Service) RetrieveByEhid(ehid string) ResponseDto {
	orgs, err := s.AccountOrganizationRepository.FindByEhid(ehid)
	if err != nil {
		return ResponseDto{
			Aggregate: superior.Aggregate{},
			Status:    err.Error(),
		}
	}

	result, err := s.SuperiorRepository.FindByOrganizationHierarchy(orgs[0].Hierarchy)
	return ResponseDto{
		Aggregate: result,
		Status:    mapper.ToStatus(err),
	}
}
