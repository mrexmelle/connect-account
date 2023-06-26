package profile

type ResponseDto struct {
	Profile Entity `json:"profile"`
	Status  string `json:"status"`
}

type PatchResponseDto struct {
	Status string `json:"status"`
}
