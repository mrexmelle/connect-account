package tenure

type Entity struct {
	Id             int64  `json:"id"`
	Ehid           string `json:"ehid"`
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
	EmploymentType string `json:"employmentType"`
	OrganizationId string `json:"organizationId"`
	TitleGrade     string `json:"titleGrade"`
	TitleName      string `json:"titleName"`
}
