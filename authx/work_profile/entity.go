package wprof

type WorkProfilePostRequest struct {
	EmployeeId       string `json:"employeeId"`
	WorkEmailAddress string `json:"emailAddress`
	StartDate        string `json:"startDate`
	EmploymentType   string `json:"employmentType`
	Password         string `json:"password"`
}
