package organization

type ResponseDto struct {
	Organization Entity `json:"organization"`
	Status       string `json:"status"`
}
