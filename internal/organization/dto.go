package organization

type SingleResponseDto struct {
	Organization Entity `json:"organization"`
	Status       string `json:"status"`
}

type MultipleResponseDto struct {
	Organizations []Entity `json:"organizations"`
	Status        string   `json:"status"`
}
