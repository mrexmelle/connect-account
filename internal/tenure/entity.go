package tenure

type TenureCreateRequest struct {
	Ehid           string
	StartDate      string
	EndDate        string
	EmploymentType string
	Ohid           string
}

type TenureRetrieveResponse struct {
	Ehid    string
	Tenures []Tenure
}

type Tenure struct {
	Id             int
	StartDate      string
	EndDate        string
	EmploymentType string
	Ohid           string
}
