package superior

import (
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/currentOrganization"
	"github.com/mrexmelle/connect-idp/internal/mapper"
)

type Service struct {
	Config                        *config.Config
	CurrentOrganizationRepository *currentOrganization.Repository
	SuperiorRepository            *Repository
}

func NewService(
	cfg *config.Config,
	cor *currentOrganization.Repository,
	r *Repository,
) *Service {
	return &Service{
		Config:                        cfg,
		CurrentOrganizationRepository: cor,
		SuperiorRepository:            r,
	}
}

func (s *Service) RetrieveByEhid(ehid string) ResponseDto {
	orgs, err := s.CurrentOrganizationRepository.FindByEhid(ehid)
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
