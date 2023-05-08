package organizationMember

type ResponseDto struct {
	Members []Aggregate `json:"members"`
	Status  string      `json:"status"`
}
