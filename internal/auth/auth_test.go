package auth

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeAndValidateJWT(t *testing.T) {
	userId := uuid.New()

	correctSecret := "chirpy"
	wrongSecret := "lol"

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		userID         uuid.UUID
		createSecret   string
		validateSecret string
		expiresIn      time.Duration
		wantErr        bool
	}{
		{
			name:           "Valid token",
			userID:         userId,
			createSecret:   correctSecret,
			validateSecret: correctSecret,
			expiresIn:      5 * time.Minute,
			wantErr:        false,
		},
		{
			name:           "Expired token",
			userID:         userId,
			createSecret:   correctSecret,
			validateSecret: correctSecret,
			expiresIn:      -1 * time.Nanosecond,
			wantErr:        true,
		},
		{
			name:           "Wrong secret",
			userID:         userId,
			createSecret:   correctSecret,
			validateSecret: wrongSecret,
			expiresIn:      5 * time.Minute,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := MakeJWT(tt.userID, tt.createSecret, tt.expiresIn)
			if err != nil {
				t.Errorf("Couldn't create token: %s", err)
			}

			extractedId, err := ValidateJWT(token, tt.validateSecret)
			fmt.Printf("LOLOLOL%vLOLOLOLOLOL", extractedId)
			if tt.wantErr && err == nil {
				t.Errorf("Expected error but got none")
				return
			}
			if !tt.wantErr && err != nil {
				t.Errorf("Unexpected error: %s", err)
				return
			}

			if !tt.wantErr && err == nil && tt.userID != extractedId {
				t.Errorf("User ID doesn't match: %v != %v", tt.userID, extractedId)
				return
			}
		})
	}
}
