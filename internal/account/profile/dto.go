package accountProfile

import "github.com/mrexmelle/connect-idp/internal/profile"

type ResponseDto struct {
	Profile profile.Entity `json:"profile"`
	Status  string         `json:"status"`
}
