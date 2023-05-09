package organizationTree

type ResponseDto struct {
	Tree   Aggregate `json:"tree"`
	Status string    `json:"status"`
}
