package superior

type ResponseDto struct {
	Aggregate Aggregate `json:"profile"`
	Status    string    `json:"status"`
}
