package accountTenure

import (
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/mapper"
	"github.com/mrexmelle/connect-idp/internal/tenure"
)

type Service struct {
	Config           *config.Config
	TenureRepository *tenure.Repository
}

func NewService(
	cfg *config.Config,
	r *tenure.Repository) *Service {
	return &Service{
		Config:           cfg,
		TenureRepository: r,
	}
}

func (s *Service) Create(request tenure.Entity) SingleResponseDto {
	result, err := s.TenureRepository.Create(request)
	return SingleResponseDto{
		Tenure: result,
		Status: mapper.ToStatus(err),
	}
}

func (s *Service) RetrieveByEhid(
	ehid string,
) MultipleResponseDto {
	result, err := s.TenureRepository.FindByEhid(ehid)
	return MultipleResponseDto{
		Tenures: result,
		Status:  mapper.ToStatus(err),
	}
}

func (s *Service) UpdateEndDateByEhidAndTenureId(
	req PatchRequestDto,
	ehid string,
	tenureId int,
) PatchResponseDto {
	err := s.TenureRepository.UpdateEndDateByIdAndEhid(
		req.Value,
		tenureId,
		ehid,
	)
	return PatchResponseDto{
		Status: mapper.ToStatus(err),
	}
}
