package superior

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/profile"
)

type Repository struct {
	Config *config.Config
}

func NewRepository(cfg *config.Config) *Repository {
	return &Repository{
		Config: cfg,
	}
}

func (r *Repository) FindByOrganizationHierarchy(hierarchy string) ([]profile.Entity, error) {
	lineage := strings.Split(hierarchy, ".")
	if len(lineage) == 0 {
		return []profile.Entity{}, errors.New("no hierarchy found")
	}

	whereClause := fmt.Sprintf("organizations.id = '%s' ", lineage[0])
	for i := 1; i < len(lineage); i++ {
		whereClause += fmt.Sprintf("OR organizations.id ='%s' ", lineage[i])
	}

	result, err := r.Config.Db.
		Table("profiles").
		Select("profiles.ehid, profiles.employee_id, profiles.name, profiles.email_address, profiles.dob").
		Joins("LEFT JOIN organizations ON profiles.ehid = organizations.lead_ehid").
		Where(whereClause).
		Order("organizations.hierarchy ASC").
		Rows()
	if err != nil {
		return []profile.Entity{}, err
	}
	defer result.Close()

	profiles := []profile.Entity{}
	for result.Next() {
		var p profile.Entity
		var dob time.Time
		result.Scan(&p.Ehid, &p.EmployeeId, &p.Name, &p.EmailAddress, &dob)
		p.Dob = dob.Format("2006-01-02")

		var i int
		for i = 0; i < len(profiles); i++ {
			if profiles[i].Ehid == p.Ehid {
				break
			}
		}
		if i == len(profiles) {
			profiles = append(profiles, p)
		}
	}

	return profiles, nil
}
