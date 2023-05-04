package organization

import (
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/hasher"
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
	ohid := hasher.ToOhid(req.Id)
	result := r.Config.Db.Exec(
		"INSERT INTO "+r.TableName+"(id, ohid, parent_id, name, lead_ehid, "+
			"created_at, updated_at) "+
			"VALUES(?, ?, ?, ?, ?, NOW(), NOW())",
		req.Id,
		ohid,
		req.ParentId,
		req.Name,
		req.LeadEhid,
	)
	if result.Error != nil {
		return Entity{}, result.Error
	}

	return Entity{
		Id:       req.Id,
		Ohid:     ohid,
		ParentId: req.ParentId,
		Name:     req.Name,
		LeadEhid: req.LeadEhid,
	}, nil
}

func (r *Repository) FindByOhid(ohid string) (Entity, error) {
	entity := Entity{
		Ohid: ohid,
	}

	err := r.Config.Db.
		Select("id, parent_id, name, lead_ehid").
		Table(r.TableName).
		Where("ohid = ?", ohid).
		Row().
		Scan(
			&entity.Id,
			&entity.ParentId,
			&entity.Name,
			&entity.LeadEhid)
	if err != nil {
		return Entity{}, err
	}
	return entity, nil
}
