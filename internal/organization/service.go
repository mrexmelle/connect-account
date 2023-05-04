package organization

import (
	"github.com/mrexmelle/connect-idp/internal/config"
)

type Service struct {
	Config                 *config.Config
	OrganizationRepository *Repository
}

func NewService(
	cfg *config.Config,
	r *Repository,
) *Service {
	return &Service{
		Config:                 cfg,
		OrganizationRepository: r,
	}
}

func (s *Service) Create(req Entity) ResponseDto {
	result, err := s.OrganizationRepository.Create(req)
	if err != nil {
		return ResponseDto{
			Organization: Entity{},
			Status:       err.Error(),
		}
	}
	return ResponseDto{
		Organization: result,
		Status:       "OK",
	}
}

func (s *Service) RetrieveByOhid(ohid string) ResponseDto {
	result, err := s.OrganizationRepository.FindByOhid(ohid)
	if err != nil {
		return ResponseDto{
			Organization: Entity{},
			Status:       err.Error(),
		}
	}
	return ResponseDto{
		Organization: result,
		Status:       "OK",
	}
}
