package organizationTree

import "github.com/mrexmelle/connect-idp/internal/organization"

type Aggregate struct {
	Organization organization.Entity `json:"organization"`
	Children     []Aggregate         `json:"children"`
}
