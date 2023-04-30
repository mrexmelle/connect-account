package profile

type ProfileCreateRequest struct {
	Ehid string
	Name string
	Dob  string
}

type ProfileRetrieveResponse struct {
	Ehid string
	Name string
	Dob  string
}
