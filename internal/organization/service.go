package organization

import (
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/mapper"
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

func (s *Service) Create(req Entity) SingleResponseDto {
	result, err := s.OrganizationRepository.Create(req)
	if err != nil {
		return SingleResponseDto{
			Organization: Entity{},
			Status:       err.Error(),
		}
	}
	return SingleResponseDto{
		Organization: result,
		Status:       "OK",
	}
}

func (s *Service) RetrieveById(id string) SingleResponseDto {
	result, err := s.OrganizationRepository.FindById(id)
	if err != nil {
		return SingleResponseDto{
			Organization: Entity{},
			Status:       err.Error(),
		}
	}
	return SingleResponseDto{
		Organization: result,
		Status:       "OK",
	}
}

func (s *Service) DeleteById(id string) SingleResponseDto {
	err := s.OrganizationRepository.DeleteById(id)
	return SingleResponseDto{
		Organization: Entity{},
		Status:       mapper.ToStatus(err),
	}
}
