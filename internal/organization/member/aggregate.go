package organizationMember

type Aggregate struct {
	Ehid         string `json:"ehid"`
	EmployeeId   string `json:"employeeId"`
	Name         string `json:"name"`
	EmailAddress string `json:"emailAddress"`
	TitleName    string `json:"titleName"`
	IsLead       bool   `json:"isLead"`
}
