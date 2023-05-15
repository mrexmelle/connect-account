package superior

import "github.com/mrexmelle/connect-idp/internal/profile"

type Aggregate struct {
	Profile  profile.Entity `json:"profile"`
	Children []Aggregate    `json:"children"`
}
