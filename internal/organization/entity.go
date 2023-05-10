package organization

type Entity struct {
	Id           string `json:"id"`
	Hierarchy    string `json:"hierarchy"`
	Name         string `json:"name"`
	LeadEhid     string `json:"leadEhid"`
	EmailAddress string `json:"emailAddress"`
}
