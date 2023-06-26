package profile

import (
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/mapper"
)

type Service struct {
	Config            *config.Config
	ProfileRepository *Repository
}

func NewService(
	cfg *config.Config,
	r *Repository) *Service {
	return &Service{
		Config:            cfg,
		ProfileRepository: r,
	}
}

func (s *Service) RetrieveByEhid(ehid string) ResponseDto {
	result, err := s.ProfileRepository.FindByEhid(ehid)
	return ResponseDto{
		Profile: result,
		Status:  mapper.ToStatus(err),
	}
}
