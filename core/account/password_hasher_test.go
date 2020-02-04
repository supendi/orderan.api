package account

import "testing"

func TestHashAndVerify(t *testing.T) {
	hasher := NewBCryptHasher()
	plainPassword := "sayang"
	hashedPassword, err := hasher.Hash(plainPassword)
	if err != nil {
		t.Fatal(err)
	}
	verifyOK, err := hasher.Verify(plainPassword, hashedPassword)
	if err != nil {
		t.Fatal(err)
	}
	if !verifyOK {
		t.Fatal("Verify error, should return true")
	}
}
