package accountTenure

import "github.com/mrexmelle/connect-idp/internal/tenure"

type SingleResponseDto struct {
	Tenure tenure.Entity `json:"tenure"`
	Status string        `json:"status"`
}

type MultipleResponseDto struct {
	Tenures []tenure.Entity `json:"tenures"`
	Status  string          `json:"status"`
}

type PatchRequestDto struct {
	Value string `json:"value"`
}

type PatchResponseDto struct {
	Status string `json:"status"`
}
