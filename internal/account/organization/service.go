package accountOrganization

import (
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/mapper"
)

type Service struct {
	Config                        *config.Config
	AccountOrganizationRepository *Repository
}

func NewService(
	cfg *config.Config,
	aor *Repository,
) *Service {
	return &Service{
		Config:                        cfg,
		AccountOrganizationRepository: aor,
	}
}

func (s *Service) RetrieveByEhid(ehid string) ResponseDto {
	result, err := s.AccountOrganizationRepository.FindByEhid(ehid)
	return ResponseDto{
		Organizations: result,
		Status:        mapper.ToStatus(err),
	}
}
