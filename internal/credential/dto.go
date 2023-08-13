package credential

type PatchRequestDto struct {
	EmployeeId      string `json:"employeeId"`
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

type PatchResponseDto struct {
	Status string `json:"status"`
}
