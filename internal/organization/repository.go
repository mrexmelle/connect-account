package organization

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/mrexmelle/connect-idp/internal/config"
)

type Repository struct {
	Config    *config.Config
	TableName string
}

func NewRepository(cfg *config.Config) *Repository {
	return &Repository{
		Config:    cfg,
		TableName: "organizations",
	}
}

func (r *Repository) Create(req Entity) (Entity, error) {
	result := r.Config.Db.Exec(
		"INSERT INTO "+r.TableName+"(id, hierarchy, name, lead_ehid, "+
			"created_at, updated_at) "+
			"VALUES(?, ?, ?, ?, NOW(), NOW())",
		req.Id,
		req.Hierarchy,
		req.Name,
		req.LeadEhid,
	)
	if result.Error != nil {
		return Entity{}, result.Error
	}

	return Entity{
		Id:        req.Id,
		Hierarchy: req.Hierarchy,
		Name:      req.Name,
		LeadEhid:  req.LeadEhid,
	}, nil
}

func (r *Repository) FindById(id string) (Entity, error) {
	entity := Entity{
		Id: id,
	}

	err := r.Config.Db.
		Select("hierarchy, name, lead_ehid").
		Table(r.TableName).
		Where("id = ? AND deleted_at IS NULL", id).
		Row().
		Scan(
			&entity.Hierarchy,
			&entity.Name,
			&entity.LeadEhid)
	if err != nil {
		return Entity{}, err
	}
	return entity, nil
}

func (r *Repository) DeleteById(id string) error {
	now := time.Now()
	result := r.Config.Db.
		Table(r.TableName).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(
			map[string]interface{}{
				"lead_ehid":             "",
				"email_address":         "",
				"private_slack_channel": "",
				"public_slack_channel":  "",
				"deleted_at":            now,
				"updated_at":            now,
			},
		)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) FindSiblingsAndAncestralSiblingsByHierarchy(hierarchy string) ([]Entity, error) {
	lineage := strings.Split(hierarchy, ".")
	if len(lineage) == 0 {
		return []Entity{}, errors.New("no hierarchy found")
	} else if len(lineage) == 1 {
		o, err := r.FindById(lineage[0])
		if err != nil {
			return []Entity{}, err
		}
		return []Entity{o}, nil
	}

	for i := 1; i < len(lineage); i++ {
		lineage[i] = fmt.Sprintf("%s.%s", lineage[i-1], lineage[i])
	}

	whereClause := "hierarchy SIMILAR TO '[A-Z0-9]*' "
	for i := 0; i < len(lineage)-1; i++ {
		whereClause += fmt.Sprintf("OR hierarchy SIMILAR TO '%s.[A-Z0-9]*' ", lineage[i])
	}

	result, err := r.Config.Db.
		Select("id, hierarchy, name, lead_ehid").
		Table(r.TableName).
		Where(whereClause).
		Where("deleted_at IS NULL").
		Order("hierarchy ASC").
		Rows()
	if err != nil {
		return []Entity{}, err
	}
	defer result.Close()

	orgs := make([]Entity, 0)
	for result.Next() {
		org := Entity{}
		result.Scan(&org.Id, &org.Hierarchy, &org.Name, &org.LeadEhid)
		orgs = append(orgs, org)
	}
	return orgs, nil
}

func (r *Repository) FindChildrenByHierarchy(hierarchy string) ([]Entity, error) {
	whereClause := fmt.Sprintf("hierarchy = '%s' OR hierarchy SIMILAR TO '%s.[A-Z0-9]*'",
		hierarchy,
		hierarchy,
	)
	result, err := r.Config.Db.
		Select("id, hierarchy, name, lead_ehid").
		Table(r.TableName).
		Where(whereClause).
		Where("deleted_at IS NULL").
		Order("hierarchy ASC").
		Rows()
	if err != nil {
		return []Entity{}, err
	}
	defer result.Close()

	orgs := make([]Entity, 0)
	for result.Next() {
		org := Entity{}
		result.Scan(&org.Id, &org.Hierarchy, &org.Name, &org.LeadEhid)
		orgs = append(orgs, org)
	}
	return orgs, nil
}
