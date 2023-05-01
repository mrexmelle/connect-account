package ehid

import (
	"crypto/sha256"
	"fmt"
)

func FromEmployeeId(employeeId string) string {
	hasher := sha256.New()
	hasher.Write([]byte(employeeId))

	return fmt.Sprintf("u%x", hasher.Sum(nil))
}
