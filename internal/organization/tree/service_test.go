package organizationTree

import (
	"testing"

	"github.com/mrexmelle/connect-idp/internal/organization"
	"github.com/stretchr/testify/assert"
)

func Test_AssignEntityIntoTree_Single(t *testing.T) {
	agg := Aggregate{
		Organization: organization.Entity{},
		Children:     []Aggregate{},
	}

	entity := organization.Entity{
		Id:           "ABC",
		Hierarchy:    "ABC",
		Name:         "Team ABC",
		LeadEhid:     "abc123",
		EmailAddress: "abc123@test.com",
	}

	service := Service{}
	service.AssignEntityIntoTree(entity.Hierarchy, entity, &agg)
	assert.Equal(t, agg.Organization.Id, entity.Id)
	assert.Equal(t, agg.Organization.Hierarchy, entity.Hierarchy)
	assert.Equal(t, agg.Organization.Name, entity.Name)
	assert.Equal(t, agg.Organization.LeadEhid, entity.LeadEhid)
	assert.Equal(t, agg.Organization.EmailAddress, entity.EmailAddress)
	assert.Empty(t, agg.Children)
}

func Test_AssignEntityIntoTree_OneChild(t *testing.T) {
	agg := Aggregate{
		Organization: organization.Entity{},
		Children:     []Aggregate{},
	}

	entity := []organization.Entity{
		{
			Id:           "ABC",
			Hierarchy:    "ABC",
			Name:         "Team ABC",
			LeadEhid:     "abc123",
			EmailAddress: "abc123@test.com",
		},
		{
			Id:           "DEF",
			Hierarchy:    "ABC.DEF",
			Name:         "Team DEF",
			LeadEhid:     "def123",
			EmailAddress: "def123@test.com",
		},
	}

	service := Service{}
	service.AssignEntityIntoTree(entity[0].Hierarchy, entity[0], &agg)
	service.AssignEntityIntoTree(entity[1].Hierarchy, entity[1], &agg)

	assert.Equal(t, agg.Organization.Id, entity[0].Id)
	assert.Equal(t, agg.Organization.Hierarchy, entity[0].Hierarchy)
	assert.Equal(t, agg.Organization.Name, entity[0].Name)
	assert.Equal(t, agg.Organization.LeadEhid, entity[0].LeadEhid)
	assert.Equal(t, agg.Organization.EmailAddress, entity[0].EmailAddress)
	assert.Equal(t, 1, len(agg.Children))
	assert.Equal(t, agg.Children[0].Organization.Id, entity[1].Id)
	assert.Equal(t, agg.Children[0].Organization.Hierarchy, entity[1].Hierarchy)
	assert.Equal(t, agg.Children[0].Organization.Name, entity[1].Name)
	assert.Equal(t, agg.Children[0].Organization.LeadEhid, entity[1].LeadEhid)
	assert.Equal(t, agg.Children[0].Organization.EmailAddress, entity[1].EmailAddress)
}

func Test_AssignEntityIntoTree_TwoChildren(t *testing.T) {
	agg := Aggregate{
		Organization: organization.Entity{},
		Children:     []Aggregate{},
	}

	entity := []organization.Entity{
		{
			Id:           "ABC",
			Hierarchy:    "ABC",
			Name:         "Team ABC",
			LeadEhid:     "abc123",
			EmailAddress: "abc123@test.com",
		},
		{
			Id:           "DEF",
			Hierarchy:    "ABC.DEF",
			Name:         "Team DEF",
			LeadEhid:     "def123",
			EmailAddress: "def123@test.com",
		},
		{
			Id:           "GHI",
			Hierarchy:    "ABC.GHI",
			Name:         "Team GHI",
			LeadEhid:     "ghi123",
			EmailAddress: "ghi123@test.com",
		},
	}

	service := Service{}
	service.AssignEntityIntoTree(entity[0].Hierarchy, entity[0], &agg)
	service.AssignEntityIntoTree(entity[1].Hierarchy, entity[1], &agg)
	service.AssignEntityIntoTree(entity[2].Hierarchy, entity[2], &agg)

	assert.Equal(t, agg.Organization.Id, entity[0].Id)
	assert.Equal(t, agg.Organization.Hierarchy, entity[0].Hierarchy)
	assert.Equal(t, agg.Organization.Name, entity[0].Name)
	assert.Equal(t, agg.Organization.LeadEhid, entity[0].LeadEhid)
	assert.Equal(t, agg.Organization.EmailAddress, entity[0].EmailAddress)
	assert.Equal(t, 2, len(agg.Children))
	assert.Equal(t, agg.Children[0].Organization.Id, entity[1].Id)
	assert.Equal(t, agg.Children[0].Organization.Hierarchy, entity[1].Hierarchy)
	assert.Equal(t, agg.Children[0].Organization.Name, entity[1].Name)
	assert.Equal(t, agg.Children[0].Organization.LeadEhid, entity[1].LeadEhid)
	assert.Equal(t, agg.Children[0].Organization.EmailAddress, entity[1].EmailAddress)
	assert.Empty(t, agg.Children[0].Children)
	assert.Equal(t, agg.Children[1].Organization.Id, entity[2].Id)
	assert.Equal(t, agg.Children[1].Organization.Hierarchy, entity[2].Hierarchy)
	assert.Equal(t, agg.Children[1].Organization.Name, entity[2].Name)
	assert.Equal(t, agg.Children[1].Organization.LeadEhid, entity[2].LeadEhid)
	assert.Equal(t, agg.Children[1].Organization.EmailAddress, entity[2].EmailAddress)
	assert.Empty(t, agg.Children[1].Children)
}

func Test_AssignEntityIntoTree_TwoGenerations(t *testing.T) {
	agg := Aggregate{
		Organization: organization.Entity{},
		Children:     []Aggregate{},
	}

	entity := []organization.Entity{
		{
			Id:           "ABC",
			Hierarchy:    "ABC",
			Name:         "Team ABC",
			LeadEhid:     "abc123",
			EmailAddress: "abc123@test.com",
		},
		{
			Id:           "DEF",
			Hierarchy:    "ABC.DEF",
			Name:         "Team DEF",
			LeadEhid:     "def123",
			EmailAddress: "def123@test.com",
		},
		{
			Id:           "GHI",
			Hierarchy:    "ABC.DEF.GHI",
			Name:         "Team GHI",
			LeadEhid:     "ghi123",
			EmailAddress: "ghi123@test.com",
		},
	}

	service := Service{}
	service.AssignEntityIntoTree(entity[0].Hierarchy, entity[0], &agg)
	service.AssignEntityIntoTree(entity[1].Hierarchy, entity[1], &agg)
	service.AssignEntityIntoTree(entity[2].Hierarchy, entity[2], &agg)

	assert.Equal(t, agg.Organization.Id, entity[0].Id)
	assert.Equal(t, agg.Organization.Hierarchy, entity[0].Hierarchy)
	assert.Equal(t, agg.Organization.Name, entity[0].Name)
	assert.Equal(t, agg.Organization.LeadEhid, entity[0].LeadEhid)
	assert.Equal(t, agg.Organization.EmailAddress, entity[0].EmailAddress)
	assert.Equal(t, 1, len(agg.Children))
	assert.Equal(t, agg.Children[0].Organization.Id, entity[1].Id)
	assert.Equal(t, agg.Children[0].Organization.Hierarchy, entity[1].Hierarchy)
	assert.Equal(t, agg.Children[0].Organization.Name, entity[1].Name)
	assert.Equal(t, agg.Children[0].Organization.LeadEhid, entity[1].LeadEhid)
	assert.Equal(t, agg.Children[0].Organization.EmailAddress, entity[1].EmailAddress)
	assert.Equal(t, 1, len(agg.Children[0].Children))
	assert.Equal(t, agg.Children[0].Children[0].Organization.Id, entity[2].Id)
	assert.Equal(t, agg.Children[0].Children[0].Organization.Hierarchy, entity[2].Hierarchy)
	assert.Equal(t, agg.Children[0].Children[0].Organization.Name, entity[2].Name)
	assert.Equal(t, agg.Children[0].Children[0].Organization.LeadEhid, entity[2].LeadEhid)
	assert.Equal(t, agg.Children[0].Children[0].Organization.EmailAddress, entity[2].EmailAddress)
	assert.Empty(t, agg.Children[0].Children[0].Children)
}
