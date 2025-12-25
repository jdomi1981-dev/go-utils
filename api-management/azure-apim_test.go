package apimanagement_test

import (
	"io"
	"net/http"
	"testing"

	apimanagement "github.com/jdomi1981-dev/go-utils/api-management"
)

func TestCreateNewHeaderAPIM(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		method     string
		url        string
		apiVersion string
		body       io.Reader
		apimKey    string
		want       *http.Request
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := apimanagement.CreateNewHeaderAPIM(tt.method, tt.url, tt.apiVersion, tt.body, tt.apimKey)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CreateNewHeaderAPIM() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CreateNewHeaderAPIM() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("CreateNewHeaderAPIM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTokenOAuth2(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req         *http.Request
		secretValue string
		want        string
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := apimanagement.CreateTokenOAuth2(tt.req, tt.secretValue)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CreateTokenOAuth2() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CreateTokenOAuth2() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("CreateTokenOAuth2() = %v, want %v", got, tt.want)
			}
		})
	}
}
