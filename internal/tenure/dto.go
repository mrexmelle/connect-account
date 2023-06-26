package tenure

type SingleResponseDto struct {
	Tenure Entity `json:"tenure"`
	Status string `json:"status"`
}

type MultipleResponseDto struct {
	Tenures []Entity `json:"tenures"`
	Status  string   `json:"status"`
}

type PatchRequestDto struct {
	Value string `json:"value"`
}

type PatchResponseDto struct {
	Status string `json:"status"`
}
