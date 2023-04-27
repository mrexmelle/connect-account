package tenure

type TenureCreateRequest struct {
	Ehid           string
	EmployeeId     string
	StartDate      string
	EmploymentType string
}

type TenureUpdateEndDateRequest struct {
	Id      int
	Ehid    string
	EndDate string
}

type TenureRetrieveResponse struct {
	Ehid    string
	Tenures []Tenure
}

type Tenure struct {
	Id             int
	EmployeeId     string
	StartDate      string
	EndDate        string
	EmploymentType string
}
