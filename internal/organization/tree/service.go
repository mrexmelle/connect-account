package organizationTree

import (
	"fmt"
	"strings"

	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/organization"
)

type Service struct {
	Config                 *config.Config
	OrganizationRepository *organization.Repository
}

func NewService(
	cfg *config.Config,
	or *organization.Repository,
) *Service {
	return &Service{
		Config:                 cfg,
		OrganizationRepository: or,
	}
}

func (s *Service) RetrieveAncestralSiblingsById(id string) ResponseDto {
	orgResult, err := s.OrganizationRepository.FindById(id)
	if err != nil {
		return ResponseDto{
			Tree:   Aggregate{},
			Status: err.Error(),
		}
	}

	orgs, err := s.OrganizationRepository.FindAncestralSiblingsByHierarchy(orgResult.Hierarchy)
	if err != nil {
		return ResponseDto{
			Tree:   Aggregate{},
			Status: err.Error(),
		}
	}

	aggregate := Aggregate{
		Organization: organization.Entity{},
		Children:     []Aggregate{},
	}
	for i := 0; i < len(orgs); i++ {
		s.AssignEntityIntoTree(orgs[i].Hierarchy, orgs[i], &aggregate)
	}
	return ResponseDto{
		Tree:   aggregate,
		Status: "OK",
	}
}

func (s *Service) AssignEntityIntoTree(
	hierarchy string,
	entity organization.Entity,
	aggregate *Aggregate,
) {
	lineage := strings.Split(hierarchy, ".")
	if len(lineage) == 0 {
		aggregate.Organization = organization.Entity{}
		aggregate.Children = []Aggregate{}
		return
	} else if len(lineage) == 1 {
		aggregate.Organization = entity
		aggregate.Children = []Aggregate{}
		return
	}

	if aggregate.Organization.Id == lineage[0] {
		newHierarchy := lineage[1]
		if len(lineage) > 2 {
			for i := 2; i < len(lineage); i++ {
				newHierarchy += fmt.Sprintf(".%s", lineage[i])
			}
		}
		i := 0
		for i = 0; i < len(aggregate.Children); i++ {
			if aggregate.Children[i].Organization.Id == lineage[1] {
				s.AssignEntityIntoTree(newHierarchy, entity, &aggregate.Children[i])
				return
			}
		}
		if i == len(aggregate.Children) {
			aggregate.Children = append(
				aggregate.Children,
				Aggregate{
					Organization: organization.Entity{Id: lineage[1]},
				},
			)
			s.AssignEntityIntoTree(newHierarchy, entity, &aggregate.Children[len(aggregate.Children)-1])
		}
	}
}
