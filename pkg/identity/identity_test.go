package identity

import (
	"testing"
)

func TestNewID(t *testing.T) {
	newID := NewID("")
	if newID == "" {
		t.Fatal("ID shouldnt be an empy string")
	}
}
