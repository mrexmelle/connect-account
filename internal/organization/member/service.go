package organizationMember

import (
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/organization"
)

type Service struct {
	Config                       *config.Config
	OrganizationRepository       *organization.Repository
	OrganizationMemberRepository *Repository
}

func NewService(
	cfg *config.Config,
	or *organization.Repository,
	omr *Repository,
) *Service {
	return &Service{
		Config:                       cfg,
		OrganizationRepository:       or,
		OrganizationMemberRepository: omr,
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
	for i, agg := range aggResult {
		if agg.Ehid == leadEhid {
			aggResult[i].IsLead = true
		}
	}

	return ResponseDto{
		Members: aggResult,
		Status:  "OK",
	}
}
