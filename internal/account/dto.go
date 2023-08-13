package account

type AccountPostRequest struct {
	EmployeeId   string `json:"employeeId"`
	Name         string `json:"name"`
	EmailAddress string `json:"emailAddress"`
	Dob          string `json:"dob"`
	Password     string `json:"password"`
}

type AccountResponse struct {
	Status string `json:"status"`
}
