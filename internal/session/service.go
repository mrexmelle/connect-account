package session

import (
	"time"

	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/credential"
	"github.com/mrexmelle/connect-idp/internal/ehid"
)

type Service struct {
	Config               *config.Config
	CredentialRepository *credential.Repository
}

func NewService(cfg *config.Config, repo *credential.Repository) *Service {
	return &Service{
		Config:               cfg,
		CredentialRepository: repo,
	}
}

func (s *Service) Authenticate(req SessionPostRequest) (bool, error) {
	return s.CredentialRepository.ExistsByEmployeeIdAndPassword(
		req.EmployeeId,
		req.Password,
	)
}

func (s *Service) GenerateJwt(employeeId string) (string, error) {
	now := time.Now()
	_, token, err := s.Config.TokenAuth.Encode(
		map[string]interface{}{
			"aud": "connect-iam",
			"exp": now.
				Add(time.Minute * time.Duration(s.Config.JwtValidMinute)).
				Unix(),
			"iat": now.Unix(),
			"iss": "connect-iam",
			"nbf": now.Unix(),
			"sub": ehid.FromEmployeeId(employeeId),
		},
	)

	if err != nil {
		return "", err
	}

	return token, nil
}
