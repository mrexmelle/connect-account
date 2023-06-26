package currentOrganization

import (
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/mapper"
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

func (s *Service) RetrieveByEhid(ehid string) ResponseDto {
	result, err := s.CurrentOrganizationRepository.FindByEhid(ehid)
	return ResponseDto{
		Organizations: result,
		Status:        mapper.ToStatus(err),
	}
}
