package models_test

import (
	"github.com/xor111xor/blockchain/internal/models"
	"testing"
)

func TestNewBlockchain(t *testing.T) {
	testCases := []struct {
		name   string
		nonce  int
		length int
	}{
		{
			name:   "Test add one",
			nonce:  1,
			length: 2,
		},
		{
			name:   "Test two",
			nonce:  2,
			length: 3,
		},
	}

	bc := models.NewBlockchain()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hash := bc.LastBlock().Hash()
			bc.CreateBlock(tc.nonce, hash)
			if bc.Length() != tc.length {
				t.Errorf("Get length %d, expect %d\n", bc.Length(), tc.length)
			}
		})
	}
}
