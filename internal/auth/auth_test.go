package auth

import (
	"net/http"
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

func TestGetBearerToken(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name: "Get Bearer Token Success",
			headers: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "Bearer Lol")
				return h
			}(),
			want:    "Lol",
			wantErr: false,
		},

		{
			name: "Get Bearer Token Success Trimmed",
			headers: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "Bearer Lol   ")
				return h
			}(),
			want:    "Lol",
			wantErr: false,
		},
		{
			name: "Get Bearer Token No Authorization header",
			headers: func() http.Header {
				h := http.Header{}
				return h
			}(),
			want:    "",
			wantErr: true,
		},
		{
			name: "Get Bearer Token Can't be Parsed 1",
			headers: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "wut is this")
				return h
			}(),
			want:    "",
			wantErr: true,
		},
		{
			name: "Get Bearer Token Can't be Parsed 2",
			headers: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "Bearer Lol Bearer Lol")
				return h
			}(),
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := GetBearerToken(tt.headers)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetBearerToken() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetBearerToken() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("GetBearerToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
