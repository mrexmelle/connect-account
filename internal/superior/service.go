package superior

import (
	"github.com/mrexmelle/connect-idp/internal/accountOrganization"
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/mapper"
)

type Service struct {
	Config                        *config.Config
	AccountOrganizationRepository *accountOrganization.Repository
	SuperiorRepository            *Repository
}

func NewService(
	cfg *config.Config,
	aor *accountOrganization.Repository,
	r *Repository,
) *Service {
	return &Service{
		Config:                        cfg,
		AccountOrganizationRepository: aor,
		SuperiorRepository:            r,
	}
}

func (s *Service) RetrieveByEhid(ehid string) ResponseDto {
	orgs, err := s.AccountOrganizationRepository.FindByEhid(ehid)
	if err != nil {
		return ResponseDto{
			Aggregate: Aggregate{},
			Status:    err.Error(),
		}
	}

	result, err := s.SuperiorRepository.FindByOrganizationHierarchy(orgs[0].Hierarchy)
	return ResponseDto{
		Aggregate: result,
		Status:    mapper.ToStatus(err),
	}
}
