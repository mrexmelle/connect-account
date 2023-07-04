package superior

import "github.com/mrexmelle/connect-idp/internal/profile"

type ResponseDto struct {
	Superiors []profile.Entity `json:"superiors"`
	Status    string           `json:"status"`
}
