package utils_test

import (
	"testing"
	"time"

	"github.com/jdomi1981-dev/go-utils/utils"
)

func TestHasWhiteSpaces(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		value string
		want  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.HasWhiteSpaces(tt.value)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("HasWhiteSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchValue(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		evaluatedValue string
		line           string
		separator      string
		want           bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.SearchValue(tt.evaluatedValue, tt.line, tt.separator)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("SearchValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCurrentTimeMillis(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		timeProvider func() time.Time
		want         int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.GetCurrentTimeMillis(tt.timeProvider)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GetCurrentTimeMillis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNowPostgres(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.NowPostgres()
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NowPostgres() = %v, want %v", got, tt.want)
			}
		})
	}
}
