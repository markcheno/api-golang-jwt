package shared

import (
	"strings"
)

// GetRootEndpoint return a root of endpoint
func GetRootEndpoint(endpoint string) string {
	s := strings.Split(endpoint, "/")
	return s[1]
}
