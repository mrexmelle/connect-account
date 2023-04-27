package account

import "github.com/mrexmelle/connect-iam/authx/tenure"

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

type AccountPatchResponse struct {
	Status string `json:"status"`
}

type AccountGetProfileResponse struct {
	Ehid string `json:"ehid"`
	Name string `json:"name"`
	Dob  string `json:"dob"`
}

type AccountGetTenureResponse struct {
	Ehid    string          `json:"ehid"`
	Tenures []tenure.Tenure `json:"tenures"`
}
