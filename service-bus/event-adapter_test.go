package servicebus_test

import (
	"net/http"
	"testing"

	servicebus "github.com/jdomi1981-dev/go-utils/service-bus"
)

func TestPostEventAdapter(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		token   string
		req     *http.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := servicebus.PostEventAdapter(tt.token, tt.req)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("PostEventAdapter() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("PostEventAdapter() succeeded unexpectedly")
			}
		})
	}
}
