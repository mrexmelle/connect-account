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

func (s *Service) PatchPassword(req PatchRequestDto) (ResponseDto, error) {
	err := s.CredentialRepository.UpdatePasswordByEmployeeIdAndPassword(
		req.NewPassword, req.EmployeeId, req.CurrentPassword)
	return ResponseDto{
		Status: mapper.ToStatus(err),
	}, err
}

func (s *Service) ResetPassword(employeeId string) error {
	return s.CredentialRepository.ResetPasswordByEmployeeId(employeeId)
}
