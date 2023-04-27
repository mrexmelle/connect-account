package credential

type CredentialAuthRequest struct {
	EmployeeId string
	Password   string
}

type CredentialCreateRequest struct {
	EmployeeId string
	Password   string
}

type CredentialDeleteRequest struct {
	EmployeeId string
}
