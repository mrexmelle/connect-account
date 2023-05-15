package accountSuperior

import "github.com/mrexmelle/connect-idp/internal/superior"

type ResponseDto struct {
	Aggregate superior.Aggregate `json:"profile"`
	Status    string             `json:"status"`
}
