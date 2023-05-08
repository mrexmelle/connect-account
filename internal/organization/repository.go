package organization

import (
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
		Where("id = ?", id).
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
