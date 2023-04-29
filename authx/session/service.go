package session

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/mrexmelle/connect-iam/authx/config"
	"github.com/mrexmelle/connect-iam/authx/credential"
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
			"exp": now.Add(time.Hour * 3).Unix(),
			"iat": now.Unix(),
			"iss": "connect-iam",
			"nbf": now.Unix(),
			"sub": s.GenerateEhid(employeeId),
		},
	)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) GenerateEhid(employeeId string) string {
	hasher := sha256.New()
	hasher.Write([]byte(employeeId))

	return fmt.Sprintf("u%x", hasher.Sum(nil))
}
