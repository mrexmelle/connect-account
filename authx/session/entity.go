package session

type SessionPostRequest struct {
	EmployeeId string `json:"employeeId"`
	Password   string `json:"password"`
}

type SessionPostResponse struct {
	Token string `json:"token"`
}
