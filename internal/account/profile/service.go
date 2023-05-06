package accountProfile

import (
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/mapper"
	"github.com/mrexmelle/connect-idp/internal/profile"
)

type Service struct {
	Config            *config.Config
	ProfileRepository *profile.Repository
}

func NewService(
	cfg *config.Config,
	r *profile.Repository) *Service {
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
