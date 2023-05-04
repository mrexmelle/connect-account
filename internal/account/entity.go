package account

type AccountPostRequest struct {
	EmployeeId string `json:"employeeId"`
	Name       string `json:"name"`
	Dob        string `json:"dob"`
	Password   string `json:"password"`
}

type AccountPostResponse struct {
	Status string `json:"status"`
}

type AccountPatchRequest struct {
	Value string `json:"value"`
}

type AccountPatchResponse struct {
	Status string `json:"status"`
}

type AccountPasswordPatchRequest struct {
	CurrentValue string `json:"currentValue"`
	NewValue     string `json:"newValue"`
}

type AccountGetProfileResponse struct {
	Ehid       string `json:"ehid"`
	EmployeeId string `json:"employeeId"`
	Name       string `json:"name"`
	Dob        string `json:"dob"`
}

type AccountGetTenureResponse struct {
	Ehid    string   `json:"ehid"`
	Tenures []Tenure `json:"tenures"`
}

type Tenure struct {
	Id             int    `json:"id"`
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
	EmploymentType string `json:"employmentType"`
	Ohid           string `json:"ohid"`
}

type AccountPostTenureRequest struct {
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
	EmploymentType string `json:"employmentType"`
	Ohid           string `json:"ohid"`
}
