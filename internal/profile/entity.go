package profile

type ProfileCreateRequest struct {
	Ehid       string
	EmployeeId string
	Name       string
	Dob        string
}

type ProfileRetrieveResponse struct {
	Ehid       string
	EmployeeId string
	Name       string
	Dob        string
}
