package currentOrganization

import "github.com/mrexmelle/connect-idp/internal/organization"

type ResponseDto struct {
	Organizations []organization.Entity `json:"organizations"`
	Status        string                `json:"status"`
}
