package hasher

import (
	"crypto/sha256"
	"fmt"
)

func ToEhid(employeeId string) string {
	hasher := sha256.New()
	hasher.Write([]byte(employeeId))

	return fmt.Sprintf("u%x", hasher.Sum(nil))
}

func ToOhid(organizationId string) string {
	hasher := sha256.New()
	hasher.Write([]byte(organizationId))

	return fmt.Sprintf("o%x", hasher.Sum(nil))
}
