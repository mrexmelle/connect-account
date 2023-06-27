package accountOrganization

import (
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/mapper"
	"github.com/mrexmelle/connect-idp/internal/organization"
)

type Service struct {
	Config                        *config.Config
	CurrentOrganizationRepository *Repository
}

func NewService(
	cfg *config.Config,
	cor *Repository,
) *Service {
	return &Service{
		Config:                        cfg,
		CurrentOrganizationRepository: cor,
	}
}

func (s *Service) RetrieveByEhid(ehid string) organization.MultipleResponseDto {
	result, err := s.CurrentOrganizationRepository.FindByEhid(ehid)
	return organization.MultipleResponseDto{
		Organizations: result,
		Status:        mapper.ToStatus(err),
	}
}
