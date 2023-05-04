package organization

type Entity struct {
	Id       string `json:"id"`
	Ohid     string `json:"ohid"`
	ParentId string `json:"parentId"`
	Name     string `json:"name"`
	LeadEhid string `json:"leadEhid"`
}
