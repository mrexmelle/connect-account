package credential

import (
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/mapper"
)

type Service struct {
	Config               *config.Config
	CredentialRepository *Repository
}

func NewService(
	cfg *config.Config,
	cr *Repository,
) *Service {
	return &Service{
		Config:               cfg,
		CredentialRepository: cr,
	}
}

func (s *Service) PatchPassword(req PatchRequestDto) (PatchResponseDto, error) {
	err := s.CredentialRepository.UpdatePasswordByEmployeeIdAndPassword(
		req.NewPassword, req.EmployeeId, req.CurrentPassword)
	return PatchResponseDto{
		Status: mapper.ToStatus(err),
	}, err
}
