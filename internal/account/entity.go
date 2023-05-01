package account

type AccountPostRequest struct {
	EmployeeId     string `json:"employeeId"`
	Name           string `json:"name"`
	Dob            string `json:"dob`
	Password       string `json:"password"`
	StartDate      string `json:"startDate`
	EmploymentType string `json:"employmentType`
}

type AccountPostResponse struct {
	Status string `json:"status"`
}

type AccountPatchRequest struct {
	Value string `json:"value"`
}

type AccountPasswordPatchRequest struct {
	CurrentValue string `json:"currentValue"`
	NewValue     string `json:"newValue"`
}

type AccountPatchResponse struct {
	Status string `json:"status"`
}

type AccountGetProfileResponse struct {
	Ehid string `json:"ehid"`
	Name string `json:"name"`
	Dob  string `json:"dob"`
}

type AccountGetTenureResponse struct {
	Ehid    string   `json:"ehid"`
	Tenures []Tenure `json:"tenures"`
}

type Tenure struct {
	Id             int    `json:"id"`
	EmployeeId     string `json:"employeeId"`
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
	EmploymentType string `json:"employmentType"`
}
