package organizationMember

import (
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/organization"
	"github.com/mrexmelle/connect-idp/internal/tenure"
)

type Service struct {
	Config                       *config.Config
	OrganizationRepository       *organization.Repository
	OrganizationMemberRepository *Repository
	TenureRepository             *tenure.Repository
}

func NewService(
	cfg *config.Config,
	or *organization.Repository,
	omr *Repository,
	tr *tenure.Repository,
) *Service {
	return &Service{
		Config:                       cfg,
		OrganizationRepository:       or,
		OrganizationMemberRepository: omr,
		TenureRepository:             tr,
	}
}

func (s *Service) RetrieveByOrganizationId(id string) ResponseDto {
	leadEhid := ""
	orgResult, err := s.OrganizationRepository.FindById(id)
	if err == nil {
		leadEhid = orgResult.LeadEhid
	}

	aggResult, err := s.OrganizationMemberRepository.
		RetrieveByOrganizationIdWithKnownLeadEhid(
			id,
			leadEhid,
		)
	if err != nil {
		return ResponseDto{
			Members: []Aggregate{},
			Status:  err.Error(),
		}
	}

	return ResponseDto{
		Members: aggResult,
		Status:  "OK",
	}
}
