package accountOrganization

import (
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/organization"
)

type Repository struct {
	Config *config.Config
}

func NewRepository(cfg *config.Config) *Repository {
	return &Repository{
		Config: cfg,
	}
}

func (r *Repository) FindByEhid(ehid string) ([]organization.Entity, error) {
	result := r.Config.Db.
		Select("organizations.id, organizations.hierarchy, organizations.name, organizations.lead_ehid, organizations.email_address").
		Table("organizations").
		Joins("LEFT JOIN tenures ON tenures.organization_id = organizations.id").
		Where("tenures.ehid = ? AND tenures.start_date < NOW() AND (tenures.end_date IS NULL OR tenures.end_date > NOW())", ehid)
	var response = make([]organization.Entity, result.RowsAffected)

	rows, err := result.Rows()
	if err != nil {
		return response, err
	}
	defer rows.Close()

	for rows.Next() {
		o := organization.Entity{}
		rows.Scan(
			&o.Id,
			&o.Hierarchy,
			&o.Name,
			&o.LeadEhid,
			&o.EmailAddress,
		)
		response = append(response, o)
	}

	return response, result.Error
}
