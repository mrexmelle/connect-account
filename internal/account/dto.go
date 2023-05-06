package account

type AccountPostRequest struct {
	EmployeeId   string `json:"employeeId"`
	Name         string `json:"name"`
	EmailAddress string `json:"emailAddress"`
	Dob          string `json:"dob"`
	Password     string `json:"password"`
}

type AccountPostResponse struct {
	Status string `json:"status"`
}

type AccountPasswordPatchRequest struct {
	CurrentValue string `json:"currentValue"`
	NewValue     string `json:"newValue"`
}
