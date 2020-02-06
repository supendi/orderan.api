package identity

import "github.com/google/uuid"

//NewID generates new ID with prefix
func NewID(prefix string) string {
	return prefix + uuid.New().String()
}
