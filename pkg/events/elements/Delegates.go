package elements

import "strings"

var (
	structMessage = "message"
)

// IsMessage ...
func IsMessage(field string) bool {
	return strings.ToLower(field) == structMessage
}
